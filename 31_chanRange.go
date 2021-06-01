package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 2)

	//这个通道设置了缓冲，所以直接发送，数据会放在缓冲区中，等待读取,
	//可以尝试改变缓冲区大小来进行调试, 改成1, 第一个数据会发送到通道，第二个就deadlock
	queue <- "one"
	fmt.Println("send one")
	queue <- "two"
	fmt.Println("send two")
	close(queue)

	time.Sleep(time.Second * 3)
	for msg := range queue {
		//这个接收消息，是按照发送消息一条一条接收的，并不像tcp的字节流
		// 你看，我把缓冲区设置得很大了，而且还sleep了一会儿，还是依然无法收到整条数据，他们依然是分开的
		fmt.Println(msg)
	}

}
