package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行该函数。
// 注意，WaitGroup 必须通过指针传递给函数。
func worker2(id int, wg *sync.WaitGroup, job <-chan int, result chan<- int) {
	fmt.Printf("Worker %d starting\n", id)

	// 睡眠一秒钟，以此来模拟耗时的任务。
	time.Sleep(time.Second)
	//fmt.Printf("Worker %d done\n", id)

	for msg := range job{
		fmt.Println("receive " , msg)
		result<- msg * 2
	}
	// 通知 WaitGroup ，当前协程的工作已经完成。
	wg.Done()
}

func main() {

	jobs := make(chan int, 100)
	result := make(chan int, 100)

	// 这个 WaitGroup 被用于等待该函数开启的所有协程。
	var wg sync.WaitGroup

	// 开启几个协程，并为其递增 WaitGroup 的计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1)	//当然也可以在for循环之前直接wg.Add(5)
		go worker2(i, &wg, jobs, result)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)


	//不能这么写，又是无限等待，会死锁
	//for ret := range result {
	//	fmt.Println("result is " , ret)
	//}
	// 阻塞，直到 WaitGroup 计数器恢复为 0，即所有协程的工作都已经完成。
	//这里虽然写的简陋一点，但是基本描绘了go协程的最佳实战，就是在main结束之前，清理掉所有已经开启的协程，编写状态清晰的协程
	wg.Wait()

	//为了防止range阻塞导致deadlock, 这里close一下，算是给通道追加了一个标记，并不代表这个通道不可用
	close(result)
	for ret := range result {
		fmt.Println("result is ", ret)
	}
	// 下面这种写法也可以实现，不过有点丑
	//for {
	//	select {
	//	case msg:= <- result:
	//		fmt.Println("result is ", msg)
	//	default:
	//		return
	//	}
	//}
}