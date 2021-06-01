package main

import "log"

//多个goroutine同时写是没有问题的，是安全的
func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Llongfile | log.Lmicroseconds)	//仔细看常量的定义都是位移操作，所以取或操作就非常方便了
}
func main() {
	log.Println("hello, world")

	log.SetPrefix("DEBUG: ")
	log.Println("go language")


	defer func() {
		if err := recover(); err != nil {
			log.Println("捕获异常，成功处理")
		} else {
			log.Fatalln("无法修复的日常，还是退出吧")	//println后会调用exit
		}

	}()
	log.Panicln("状态非法，此处不应该出现这个状态，ID 100")	//调用完日志后，会调用panic
}
