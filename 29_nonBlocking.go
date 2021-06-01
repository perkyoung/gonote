package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	/*
		select 如果没有事件，就等一个事件，
		如果有多个事件，就任意选择一个事件
		如果没有事件，想非阻塞，那就加default
	*/
	select {
	case msg := <-messages: //默认阻塞，所以没有事件
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 一个非阻塞发送的实现方法和上面一样。
	msg := "hi"
	select {
	case messages <- msg://注意这里，不要以为有msg就可以读到数据，因为默认情况都是阻塞的，只有在接收方准备好了，才可以发送给管道， 这次select并没有检测到任何事件，除非加上缓冲
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//类似多路复用
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
