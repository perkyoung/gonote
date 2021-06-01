package main

import (
	"fmt"
	"time"
)

//利用go的协程去实现线程池, 启动多个go 协程即可
func main() {

	jobs := make(chan int, 100)		//用户接收任务
	result := make(chan int, 100)	//用于收集最后的结果

	for i:= 0; i < 3; i++ {
		go worker(i, jobs, result)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	//这种模式像极了master进行接受新的连接，给jobs通道
	//worker进程获取已连接套接字，进行处理
	//发送完任务之后，close发送任务通道
	close(jobs)

	//开始收集结果
	stop := false
	for !stop {
		select {
		//注意，如果用到了无限循环去等待，没有数据发送，只有数据接收，那就会陷入无限阻塞，go会提示死锁，所以， 要加上下面的「等待超时」
		case msg := <- result:
			fmt.Println("the result is ", msg)
		case <- time.After(time.Second * 5):	//这里只是一个举例，现实中可能这个事件并不确定是多久才合适
				stop = true	//这里写break只会跳出select, 所以需要用一个变量做标记
		}
	}

	//或者知道result的结果数量，循环读取结果
}

func worker(id int, job <-chan int, result chan<- int) {
	//记住记住，轮询chan，一定要用for range ，不要用for循环，后者即使在通道close后，依然可以读到数据（零值), 再close之前是阻塞的，这个是不是坑
	//for range 时，close(chan) 会跳出循环
	for msg := range job {
		ret := msg * 2
		fmt.Println("work " , id, ", receive msg" , msg, "resunt is " , ret)
		time.Sleep(time.Second * 1)
		result <- ret
	}
}