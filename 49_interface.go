package main

import "fmt"

type Matcher interface {
	search() (ret *int)
}

type defaultMatcher struct {
	count int
}

//这个类型和search进行了绑定
func (m *defaultMatcher) search() (ret *int) {
	a := new(int)
	*a = 10
	//需要维护这个状态，所以最好使用指针传递这个defaultMatcher
	//反之，如果不需要维护内部数据状态，也不需要使用指针传递参数
	m.count++
	return a
}

func (m *defaultMatcher) getCount() (count int)  {
	return m.count
}

type newMatcher struct {
}

//如果这里定义为指针类型，那么就会出现错误，想想为什么
//直接用对象或者指针调用是没问题的，但是一旦用到了多态，就会有问题
func (m newMatcher) search() (ret *int) {
	a := new(int)
	*a = 11
	return a
}

func test(matcher Matcher) (ret *int) {
	return matcher.search()
}

func main() {
	a := new(defaultMatcher)

	var b newMatcher
	c := new(newMatcher)
	fmt.Println("ret is ", *(test(a)))	//因为defaultMatcher用指针与接口方法绑定的，所以这里给Matcher类型赋值只能传指针
	fmt.Println("ret is ", *(test(b)))	//因为newMatcher 用值与接口方法绑定的，所以这里可以是值，也可以是指针
	fmt.Println("ret is ", *(test(c)))

	//var d defaultMatcher
	//fmt.Println("ret is ", *(test(d)))	这个不行, 绑定到的是defaultMatcher指针类型，

	//为什么？因为go编译器并不是总能获得值的地址，所以干脆，这么用就编译失败吧
}