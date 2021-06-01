package main

import (
	"fmt"
	"reflect"
)

type Myinterface interface {
	getCount() (ret int)
}

type OneStruct struct {
	count int
}
//类型与接口绑定
func (m *OneStruct) getCount() int {
	m.count++
	return m.count
}

type TwoStruct struct {
	count int
}
//绑定
func (m TwoStruct) getCount() int {
	m.count = m.count * 1
	return m.count
}

func test1() {
	var a = OneStruct{
		count: 10,
	}
	/*
		func (m OneStruct) getCount() int
		如果与值绑定，这里用值赋值没有问题
		var mt Myinterface = a
		fmt.Println(mt.getCount())

		这样赋值指针也没问题，因为指针肯定也可以解引用为值
		var mt Myinterface = &a
		fmt.Println(mt.getCount())
	*/

	//getcount与指针绑定，则只能传递指针，不能传值
	var mt Myinterface = &a
	fmt.Println(mt.getCount())
	fmt.Println(reflect.TypeOf(a))

}

func main() {
	test1()
}
