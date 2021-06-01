package main

import (
	"context"
	"fmt"
	"time"
)

func someHandler() {
	// 创建继承Background的子节点Context
	ctx,_ := context.WithTimeout(context.Background(), 3 * time.Second)
	go doSth(ctx)

	//模拟程序运行 - Sleep 5秒
	time.Sleep(1 * time.Second)
	//cancel()	//3s后其实已经取消了，这里重复调用没有问题, 这里这么写只是为了测试，最好的方法是defer cancel()
	select {
	case <- ctx.Done():
		fmt.Println("master done")
	}

	//time.Sleep(10 * time.Second)
}

//每1秒work一下，同时会判断ctx是否被取消，如果是就退出
func doSth(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(5 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("dosth done")
			return
		default:
			fmt.Printf("work %d seconds: \n", i)
		}
		i++
	}
}

func main() {
	fmt.Println("start...")
	someHandler()
	fmt.Println("end.")
}
