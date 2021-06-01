package main

import (
	"context"
	"fmt"
	"time"
)

//对 73 进行完善
func gen2(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
				time.Sleep(time.Second)	//注意这里最好sleep一下，如果生产速度大于消费速度，那么就会阻塞在这里，无法跳出select, deadlock
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()	//避免其他地方忘记，所以这里defer，重复调用不影响
	for a := range gen2(ctx) {	//注意这里range，其实就是调用了一次gen2(), 返回了chan后，后续都是range这个chan
		println(a)
		if a == 5 {
			cancel()
			break
		}
	}
	fmt.Println("all goroutine kill")
	time.Sleep(time.Second)
}
