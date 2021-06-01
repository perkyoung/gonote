package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//go的协程思想来源于通信顺序进程(Communicating Sequential Processes，CSP) 的范型(paradigm)
//通过传递数据来传递消息，而不是通过锁来实现同步
//调度器是非常复杂的软件，可以管理调度所有goroutine，为其分配事件
var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	//fmt.Println("cpu is ", runtime.NumCPU())
	wg.Add(2)

	//go incCounter(1)	//存在竞争条件，会有问题
	//go incCounter(2)

	go atomicIncCounter(1)
	go atomicIncCounter(2)

	//atomic.StoreInt64(&counter, 1)	//这两行代码放在这里不合适，只是为了笔记，安全得获取设置值
	//atomic.LoadInt64(&counter)
	wg.Wait()
	fmt.Println("final counter: ", counter)
}

//go build -race 竞争检测器标志来编译程序
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter	//读取-计算-写入， 期间出让了cpu，计算的结果可能会有问题
		runtime.Gosched()	//出让cpu，先调度到别的goroutine
		value++
		counter = value
	}
}

func atomicIncCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)	//利用原子函数
		runtime.Gosched()
	}
}
