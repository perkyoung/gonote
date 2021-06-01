package main

import (
	"fmt"
	"time"
)

func main() {
	//请求外部接口，经常会遇到超时的情况，如果在go协程中去调用，那超时应该如何处理呢？ 通道+select, 同时等待，看谁先到，也就是返回结果和超时哪个先到
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case msg1 := <-c1:	//等待c1的事件发生
		fmt.Println("result is ", msg1)
	case <-time.After(time.Second * 3):	//等待3s的超时
		fmt.Println("timeout ")
	default:	//这里这么写肯定是有问题的，不过只是为了说明，没事件就走default了，这个不等，这个适用于多路复用的情况啊
		fmt.Println("no single")
	}
}
