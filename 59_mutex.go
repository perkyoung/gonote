package main

import (
	"fmt"
	"runtime"
	"sync"
)

//互斥锁创建临界区
var (
	counter2 int
	wg2      sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	wg2.Add(2)

	go incCounter2(1)
	go incCounter2(2)

	wg2.Wait()
	fmt.Println("final counter ", counter2)
}

func incCounter2(id int64) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter2
			runtime.Gosched()	//退出当前线程，会再分配一个线程继续运行
			value++
			counter2 = value
		}
		mutex.Unlock()
	}
}
