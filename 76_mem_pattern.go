package main

import (
	"fmt"
)

var cc = make(chan int)
var aa string

func ff() {
	aa = "hello, world" //1
	<-cc            //2
}

func main() {
	go ff()   //3
	cc <- 0      //4
	fmt.Print(aa) //5
}
