package main

import "fmt"

//为普通类型绑定方法；先type定义
type NN int

func (n NN) test() {
	fmt.Printf("test.n: %p, %v\n", &n, n)
}

func test() {
	var n NN = 100
	p := &n

	//注意，当method value被赋值给变量或者作为参数传递时，会
	//立即计算并赋值该方法执行所需要的review对象，并绑定，稍后执行时，会隐式传入receiver参数
	n++
	f1 := n.test	//receiver时N类型，所以复制n，101

	n++
	f2 := p.test	//复制 *p, 102

	n++
	fmt.Printf("main.n: %p, %v\n", p, n)

	f1()
	f2()
}

func call(m func()) {
	m()
}

func test2() {
	var n NN = 100
	p := &n
	fmt.Println("main.n: %p, %v\n", p, n)

	n++
	call(n.test)

	n++
	call(p.test)
}

func main() {
	test()
}



