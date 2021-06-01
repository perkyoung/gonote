package main

import (
	"fmt"
	"reflect"
)

// 通道方向Directions， 可以通过函数参数定义，通道的方向，指定某些通道只能用来接收数据，
//有些只能用来方法，注意，这里是限制了在函数中的使用，因为通道本来就是可读可写的，加箭头
//ping函数首先接受到msg
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//pong函数接受来自ping通道的数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings	//ping将数据传递给msg
	pongs<- msg		//msg将数据写入pong
}

func main() {
	pings := make(chan string, 1)
	fmt.Println(reflect.TypeOf(ping))
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)	//输出pong，这里并没有做方向上的限制
}

