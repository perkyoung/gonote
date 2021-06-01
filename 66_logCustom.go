package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

//自定义日志
var (
	Trace *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile("./gologfile", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open file failed")
	}
	logflag := log.Ldate | log.Lmicroseconds | log.Llongfile

	Trace = log.New(ioutil.Discard, "TRACE: ", logflag)
	Debug = log.New(os.Stdout, "DEBUG: ", logflag)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", logflag)	//可以指定多个输出
}

func main() {
	Error.Println("status is failed")
	Trace.Println("this is trace log")
}