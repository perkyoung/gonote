package main

import (
	"fmt"
	"net/http"
	"time"
)


func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(&request)

	//默认情况下，对于每个请求，会单独开辟一个goroutine
	go func(){
		for range time.Tick(time.Second) {
			fmt.Println("current goroutine is in processing")
		}
	}()
	//主协程没有结束，这个协程也不会结束, 即使这个handler已经处理完了，但是上面这个goroutine也会继续输出

	time.Sleep(time.Second * 3)
	writer.Write([]byte("Hi"))
}

func helloHandler2(write http.ResponseWriter, request *http.Request)  {
	fmt.Println(&request)

	go func() {
		for range time.Tick(time.Second) {
			select {
			//会检测处理请求的协程是否处理完了，如果处理晚了，那么这里检测到这个事件，就return
			case <- request.Context().Done():
				fmt.Println("request is outgoing")
				return
			default:
				fmt.Println("Current request is in processing")
			}
		}
	}()

	time.Sleep(time.Second * 5)
	write.Write([]byte("helloHandler2"))
}

//go语言，每个请求都是通过一个单独的goroutine来做的，但是这个goroutine可能调用了其他http，rpc服务，这个也是用goroutine来做的
//这么多的goroutine，如何同步数据？goroutine应运而生，还有其他功能
//如果一个api请求的goroutine的整个调用链路是多层的话，那么context就是从上层一层一层得传递到了下层，如果没有context，生成如果发生错误，下层
//不会受到影响，会一直运行下去，其实此时可以终止了，可以利用context机制来及时停止下层无用的工作，避免计算资源的浪费
//而且context还可以携带以请求为作用域的键值对信息
func main() {

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
