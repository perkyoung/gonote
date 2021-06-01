package main

import "fmt"

//go语言的类型系统

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

//将普通方法绑定到user
//func 与函数名之间的参数被称为接收者，将函数与接收者的类型绑定起来
//使用值接收者声明方法，调用时会使用这个值的副本来执行。
func (u user) notify() {
	fmt.Println(u.name, u.email)
}

//将方法绑定到user, 注意这里要改变成员的值，所以要绑定到指针
//使用指针接收者声明方法，调用时依然会使用这个指针的副本来执行，指针的副本同样指向原始的值，没问题
func (u *user) changeEmail(newname string)  {
	u.email = newname
}

type admin struct {
	person user
	level  int
}

func main() {
	var bill user //成员被初始化为零值
	bill.name = "perkyoung"

	newbill := user{
		name:  "lisa",
		email: "163.com",
	}
	fmt.Println(newbill)

	bill2 := user{"aaa", "bbb", 10, true}
	fmt.Println(bill2)

	//嵌套赋值
	admin1 := admin{
		person: user{
			"aaa", "bbb", 11, false,
		},
		level: 2,
	}
	fmt.Println(admin1)
	//直接取结构体赋值
	admin1 = admin{
		newbill,22,
	}
	fmt.Println(admin1)

	type newtype int64
	var a newtype
	a = newtype(100)
	//a = int64(200)	虽然底层他俩是同一个类型，但是不能这么做，go认为他俩是完全不同的类型, 编译器没有隐式转换
	fmt.Println(a)


	//方法
	bill3 := user{
		"name", "email",10,false,
	}

	//这个调用看起来像使用某个包，其实不是；
	//本质上，在调用notify时，使用bill的值作为接收者进行调用，方法notify会接收到bill的值的副本，操作的是这个副本
	bill3.notify()

	//问题，能否使用指针作为接收者，调用notify方法呢？可以的
	bill4 := &user{
		"myname", "myemail", 11, true,
	}
	bill4.notify()
	//为什么？调用notify传递是指针吗？肯定不是，因为方法声明需要传递值。其实是go做了从指针到值的转换，解引用，如下
	(*bill4).notify()	//也就是可以从指针到值的转换, 然后再做一个副本，notify操作的是这个副本

	bill4.changeEmail("my new email")	//指针作为接收者，操作的是指针的副本，但是指针的副本指向的依然是那个对象，没有任何问题
	bill4.notify()

	//这样也可以，编译器做了取地址的操作，然后方法操作的依然是地址的副本，指向的依然是那个对象。
	bill3.changeEmail("last email")
	bill3.notify()	// (&bill3).notify()	所以就算传递的是值，也可以最终修改email，因为接收者要接受的是指针。go做了去地址操作

	//总结起来，操作的都是副本，
	//如果声明接收者是值；1. 传值，副本，操作副本； 2. 传指针，go编译器解引用，副本，操作副本
	//如果声明接收者是指针；1. 传值，取地址，地址副本，操作这个副本，实际操作的还是原始对象。 2.传指针，地址副本，操作这个副本，实际操作的还是原对象

}
