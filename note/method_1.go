package main

import (
	"fmt"
)

type N int

func (n N) value() {
	n++
	fmt.Println("value is ", n)
}

func (n *N) pointer()  {
	*n++
	fmt.Println("value is ", *n)
}

func (n N) test() {
	value := N(10)
	value.pointer()
}

type user struct {}

type manager struct {
	user
}

func (user) toString() string {
	return "user"
}

func (m manager) toString() string {
	return m.user.toString() + " : manage"
}

func (m manager) test() {
	fmt.Println(m.toString())
	fmt.Println(m.user.toString())

}

type S struct {

}
type T struct {
	S
}

func (S) sVal() {
	fmt.Println("sVal")
}

func (*S) sPtr() {
	fmt.Println("sPtr")
}

func (T) tVal() {
	fmt.Println("tVal")
}

func (*T) tPtr() {
	fmt.Println("tPtr")
}

func main() {
	var t T
	t.sPtr()
	t.sVal()
	t.tVal()
	b := t.tPtr
	b()
}
