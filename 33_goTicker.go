package main

import (
	"fmt"
	"time"
)

//固定的时间间隔去执行某任务, 打点器， 每隔一段时间发生一个时间
func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)
	go func() {
		//像是对定时器的进一步封装，这一步可以一直循环的去获取ticker的事件，而NewTimer只能获取一次就到时了
		i := 0
		for t := range ticker.C {
			fmt.Println("tick at ", t)
			i++
			if i > 4 {
				ticker.Stop()
				done <- true
			}
		}
	}()

	//由于每次获取ticker.C都是一次计数，所以每次相当于一个无限循环，想什么时候停止的时候，就只能stop了
	//time.Sleep(time.Millisecond * 1600)
	//ticker.Stop()
	<- done
}
