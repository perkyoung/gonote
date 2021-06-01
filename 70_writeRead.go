package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Fprintf(&b, " world")
	b.WriteTo(os.Stdout)

	r, err := http.Get("baidu.com")
	if err != nil {
		fmt.Println("baidu.com error")
		return
	}

	file, err := os.Create("./writeReade.txt")
	if err != nil {
		fmt.Println("open txt file failed")
		return
	}

	//返回多个地址
	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
