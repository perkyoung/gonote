package main

import (
	"time"
)
import "fmt"

func main() {

	//全部限速
	request := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)	//防止for range deadlock

	limit := time.Tick(time.Millisecond * 200)	//其实就是返回了 NewTrick.C ; for .. range NewTrick.C 也没问题
	for req := range request {
		<-limit
		fmt.Println("request ", req)
	}

	//前面的不限速，只对后面的限速
	//利用通道缓冲
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()	//设置通道内值的时间，连续的时间
	}

	fmt.Println("here1")

	burstyLimiter = make(chan time.Time, 100) //带缓冲的，这个值最少是3
	// 想将通道填充需要临时改变3次的值，做好准备。 先写入通道，这样读chan数据的时候，就不会阻塞
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//刚开始有一个误解，以为写入了连续的时间，后面<-burstyLimiter 就会不停顿地连续，读取；如果写入的时间是有时间差的，
	//<-burstyLimiter就会等会读取下一个值，其实不是这样的，只要准备好了值，就会不停地读到新的值，不会发生阻塞
	//所以，需要用go协程去真的异步地去写入，主携程延后延后才能读到
	//主协程去处理请求，子协程用来设置限速，
	go func() {
		for newTick := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- newTick
		}
	}()

	newRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		newRequest <- i
	}
	close(newRequest)

	for req := range newRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}