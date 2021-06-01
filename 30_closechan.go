package main

import (
	"fmt"
	"time"
)

func main() {

	//用于发送任务
	jobs := make(chan int, 5)

	//子协程处理完成后，通过这个通道发送给主协程，主协程一致在阻塞等待这个事件的发生，主协程不能
	// 提前结束，因为子协程可能还没有完成任务
	done := make(chan bool)

	go func() {
		time.Sleep(time.Second * 10)
		for {
			msg, more := <-jobs
			if more { //如果有数据,则处理
				fmt.Println("recieve ", msg)
			} else {
				fmt.Println("主协程关闭，不会再有新任务 ")
				//告知主协程，已经处理完了
				done <- true
				//这里一定要加return，否则还会继续再循环，再发送一次true, 此时会阻塞在这里，因为主协程已经没有读chan的操作了。
				//一直阻塞，直到主协程结束，自己也会结束。或者如果done是有缓冲区的，那不会阻塞，会一直发，一直循环下去
				//直到填满缓冲区再阻塞
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("send one job ", i)
	}
	// close 之后，数据接收方依然可以接收数据，这个close可以理解为一个结束不再发送消息的标志，但并不妨碍另一端接收数据
	close(jobs)
	fmt.Println("已经close，等待结束, 等待数据接收方接收完数据")
	<-done	//这里还是需要等待的，因为close之后，并不代表数据接收端处理完了数据
	fmt.Println("关闭主协程")
	time.Sleep(time.Second * 10)
}
