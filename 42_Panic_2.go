package main

import "fmt"

/*
defer的执行顺序:
	返回值 = xxx
	调用defer函数(这里可能会有修改返回值的操作)
	return 返回值
 */

func f() (result int) {	//定义了返回值就是result
	defer func() {
		result++
	}()
	return 0	//0 赋值给返回值result，然后执行defer，修改了result（+1），返回result，所以应该是1
}

func f_2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t	//定义了返回值为r；t= 5;t赋值给r为5；执行defer，将t改为10；返回r,此时r肯定还是5；因为没有改变
}

func f_3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1	//定义返回值为r，将1赋值给r；执行defer，将r传入defer，注意，这里是形参，并没有改变上层r的值，换个名字更直观
	//这样写可读性很差，要么换个名字；如果真的想修改，那就传地址
}

func c() {
	panic("throw c")
}
func b() {
	c()
}

func a() (i int){
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("catch error")
		}
	}()
	b()	//内部会panic，程序即将结束，结束前要调用defer, 输出catch error，下面的代码不会被执行了
	i = i + 100
	return
}

//注意recover只有在当前的goroutine内进行捕获，外层函数无法捕获，而且会造成程序崩溃
func main() {
	fmt.Println(a())
}