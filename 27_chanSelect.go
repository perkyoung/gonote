package main

import (
	"fmt"
	"time"
)

func main() {
	// go的通道选择器select 让你可以同时等待多个通道操作。协程+通道+select是go的一个强大特性
	// 步骤，创建多个通道，跑多个协程；for 循环等待多个通道的数据
	c1 := make(chan string)
	c2 := make(chan string, 100)

	//如果没有缓冲区，写通道只能等读通道准备就去才能写入。所以如果在这里这样， 肯定会死锁
	//因为你在等对方读，但是对方在等逻辑走到那里
	//	c1 <- "one"
	//，所以只能用协程去写,如下

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
		//close(c1)	，按照表象来看，close也算是一种消息，如果加了close，会触发 msg := <-c1 消息，但是是空
	}()
	go func() {
		time.Sleep(time.Second )
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		//注意，如果在select中加break,没用，不会跳出for循环， 只会跳出select
		case msgs1 := <-c1:	//这里有个很大的坑，如果chan被关闭，这里依然是可以读到事件的，所以外层的for控制很重要。最好用for range
			fmt.Println("receive1", msgs1)
		case msgs2 := <-c2:
			fmt.Println("receive2", msgs2)
		case <- time.After(time.Second * 3):
			return
		}
	}
}
