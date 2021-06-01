package main

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
)

type Poll struct {
	mutex    sync.Mutex
	closed   bool
	resource chan io.Closer            //是一个接口，可以保存所有实现了他的接口的struct，注意多态情况下，接口不必指定是指针类型，就这样没问题
	factory  func() (io.Closer, error) //这个方法可以创建具体的资源，也就是fn和Poll是绑定的，一个poll内部只能包含以一种资源，如果想创建多种资源，轻使用多个pool
}

var errPoolClose = errors.New("pool has close")

func NewPoll(fn func() (io.Closer, error), size int) (*Poll, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &Poll{
		//mutex:    sync.Mutex{},	不用显式赋值，零值即可
		closed:   false,
		resource: make(chan io.Closer, size),
		factory:  fn,
	}, nil
}

//资源池获取一个资源
func (p *Poll) Acquire() (io.Closer, error) {
	if p.closed {
		return nil, nil
	}

	//在并发场景下，close之后，依然获取了资源，没有问题，只是在释放资源的时候，会自动销毁
	select {
	case resource, ok := <-p.resource:
		fmt.Println("acquire one resource")
		if !ok {
			return nil, errPoolClose // 容易遗漏的地方
		}
		return resource, nil
	default:
		fmt.Println("没有资源，生成一个资源")
		return p.factory()
	}
}

// 释放资源, 即向资源放回到poll
func (p *Poll) Release(resource io.Closer) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	//释放资源的时一定要检测整个poll是否已经关闭，如果关闭，释放资源
	if p.closed {
		resource.Close()
		return nil
	}
	select {
	case p.resource <- resource: //注意，如果向已经close的通道写数据，会崩溃，所以前面做了互斥操作，close 和release使用的是同一把锁
		fmt.Println("release one resource")
	default:
		//chan满，直接close掉这个资源
		fmt.Println("release: close resource")
		resource.Close()
	}
	return nil
}

type DbConnection struct {
	ID int32
}

func (db *DbConnection) Close() error {
	fmt.Println("close connections ", db.ID)
	return nil
}

var idCounter int32

func createDb() (io.Closer, error) {
	return &DbConnection{
		atomic.AddInt32(&idCounter, 1),
	}, nil
}

//释放所有资源，关闭
func (p *Poll) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return nil
	}

	p.closed = true

	//关闭通道，注意，此时通道内可能还存在消息，是可以使用的
	close(p.resource)

	for resource := range p.resource {
		resource.Close()
	}
	return nil
}

const (
	maxGoroutines = 25
	poolResource  = 2
)

func main() {
	var wg sync.WaitGroup
	p, err := NewPoll(createDb, poolResource)
	if err != nil {
		fmt.Println("create pool failed")
	}

	for query := 0; query <= maxGoroutines; query++ {
		go func(q int) {
			defer wg.Done()
			performquery(query, p)
		}(query)
		wg.Add(1)
	}

	wg.Wait()
}

func performquery(query int, p* Poll)  {
	resource, err := p.Acquire()
	if err != nil {
		fmt.Println("acquire failed")
	}
	defer p.Release(resource)

	fmt.Println("this connections id is ", resource.(*DbConnection).ID)	//类型断言 x.(T)

}
