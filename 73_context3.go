package main

import (
	"fmt"
	"time"
)

func gen() chan int {
	ch := make(chan int)
	var n int
	go func() {
		for {
			n++
			ch <- n
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	for a := range gen() {
		fmt.Println(a)
		if a == 5 {
			break	//潜在的问题，break后，子协程还会继续生成数据，称之为goroutine泄露
		}
	}
}