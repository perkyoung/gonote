package main

import (
	"os"
	"fmt"
	"time"
)

func main() {

	var user = os.Getenv("USER_")
	var result int
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {	//相当于panic 抛出异常，defer捕获异常，recover处理异常
				fmt.Println("recover success.")
				result = 2
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
			fmt.Println("after panic")
		}()
	}()
	defer func() {
		fmt.Println("defer main")
		result = 3
	}() // will this be called when panic?

	time.Sleep(1 * time.Second)
	fmt.Printf("get result %d\r\n", result)
}