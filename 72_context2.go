package main

import (
	"context"
	"fmt"
	"time"
)

func helloHandler3(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with ", duration)
	}
}

func main() {
	//context.Background() 返回一个空context。这个只能用于高等级中的，比如main或者顶级请求处理中
	//在多数情况下，如果当前函数没有上下文作为入参，我们都会使用 context.Background 作为起始的上下文向下传递
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)	//定义超时事件，1s
	//父子协程都会检测ctx.Done();

	defer cancel()

	go helloHandler3(ctx, 500*time.Millisecond)	//这个超时时间是小于1s，所以提前结束
	select {
	case <-ctx.Done():
		fmt.Println("hello handler3", ctx.Err())
	}
}
