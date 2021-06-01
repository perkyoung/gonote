package main

import "fmt"

func main() {
	x := 10
	x, s := 100, "young" //x 已经被定义过，依然可以这么写，因为起码有一个s是未定义过的, 「退化赋值」
	println(x, s)

	p := (*int)(&x) //类型转换，变量一定要加括号
	var px *int = &x
	println(px)
	println(p)

	type ( //定义一个类型组
		user struct {
			name string
			age  int
		}
		event func(string) bool
	)

	u := user{"adfasd", 11}
	var f event = func(s string) bool {
		return s != ""
	}
	fmt.Println(u)
	f("aaa")

	//未命名类型
	var c struct {
		name string
		age  int
	}
	c.name = "dfad"
	c.age = 10

	//var a, b struct{}	//零长度对象
	ddd := struct {
		x int
	}{}
	ddd.x = 10

	//匿名函数，chan
	nfunc := make(chan func(int,int) bool, 2)
	nfunc <- func(x, y int) (bool) {
		if x < y {
			return true
		}
		return false
	}
	if (<-nfunc)(10, 100) {
		println("yes")
	} else {
		println("no")
	}
}
