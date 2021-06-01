package main

import (
	"fmt"
)

/*
数组：定长的序列
 */
func main() {
	var a [6]int
	a[3] = 10

	b := [6]int{2,3,4,5,6,7}
	for index, value := range b {
		fmt.Println(index, value)
	}

	//不知道多少个, 让go自动计算声明数组的长度
	c := [...]int{2,3,4,5,6,7}
	for index, value := range c {
		fmt.Println(index, value)
	}

	//指定下标进行初始化
	d := [6]int{0: 2,1 : 3, 4: 7}
	for i := range d {	//这里输出的是index
		fmt.Println(i)
	}
	//len := 10
	//var d [len]int	这样是不行的len必须是常量

	//可以是指针
	e := [...]*int{0 : new(int), 20: new(int)}
	fmt.Println(len(e))	//输出21

	//在go中数组是值，所以可以赋值，但是要注意两个数组必须长度相同才可以赋值
	var f [21]*int	//编译的时候e的len已经固定，所以他俩是相同的
	f = e
	fmt.Println(len(f))

	tmp := new(int)
	*tmp = 100
	f[2] = tmp

	fmt.Println(*(f[2]))

	*f[0] = 300
	fmt.Println(*(e[0]))	//f和e的0号元素的指针指向同一块内存区域，修改了*f[0]的值，那么e也会更改了

	//再看一个字符串指针的例子
	array1 := [3]*string{new(string), new(string), new(string)}
	*array1[0] = "hello"
	*array1[2] = "world"	//注意，如果初始化array1时，第三个元素不new空间，这个地方同样也不会为world开辟空间，编译会出错

	var array2 [3]*string
	array2 = array1
	fmt.Println(*array2[0])	//会输出hello
	fmt.Println(*array2[2])	//会输出world

	//多维数组
	var array3 [4][2]int	//数组有4个元素，每个元素是一个包含两个int的数组
	array3[0][0] = 10
	array4 := [4][2]int{{1,3}}
	fmt.Println(array4)
	array4 = [4][2]int{1: {1,3}}	//指定下标初始化
	fmt.Println(array4)

	var array5 [2]int = array4[1]	//子数组对其他数组赋值
	fmt.Println(array5)

	//函数传递，为了减少开销，一定要传指针，否则每次都会拷贝整个数组
}
