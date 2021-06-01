package main

//嵌入类型

import (
	"fmt"
	"goExample/testPackage"
)

type Person struct {
	name  string
	email string
	age   uint8
}

func (p Person) showName() string {
	return p.name
}

func (p Person) showEmail() string {
	return p.email
}

type Admin struct {
	Person	//嵌入类型，他的成员会提升到admin这个层级, Person是外部类型Admin的内部类型
	level uint8
}

//外部类型可以直接访问内部类型的成员
func (admin *Admin) changeName()  {
	admin.name = "new name"
}

//如果定义了接口，Person实现了结构，并且作为Admin的内部类型，Person类型是可以多态的，那么Admin类型可以多态吗？
type notifer interface {
	notify()
}

//Person 作为了Admin的内部类型，自动实现的接口自动提升到外部类型
//内部类型实现了这个接口，相当于外部类型也同样实现了这个接口
//很方便
func (p *Person) notify() {
	fmt.Println("hello, my name is ", p.name)
}

//，但是。。。。如果外部类型并不希望这样，想自己实现一套呢？
//这就导致内部方法没有被提升
func (admin *Admin) notify()  {
	fmt.Println("hello, I'm admin")

}
func duotaifunc(no notifer) {
	no.notify()
}

func main() {
	//person := Person{
	//	name:  "perkyoung",
	//	email: "xxx@gmail.com",
	//	age:   33,
	//}
	adminUser := Admin{
		Person: Person{	//也可以使用已经声明过的对象赋值, Person: person
			name:  "perkyoung",
			email: "xxx@gmail.com",
			age:   33,
		},
		level: 10,
	}

	//这两种方法都可以
	fmt.Println(adminUser.Person.showEmail())
	fmt.Println(adminUser.showEmail())
	adminUser.changeName()
	fmt.Println(adminUser.showName())

	//厉害，强大，因为内部类型标识符提升到了外部类型，所以相当于admin实现了那个接口
	duotaifunc(&adminUser)
	duotaifunc(&adminUser.Person)


	//公开或者未公开的标识符，可见性
	outcount := testPackage.Outcount(11)	//大写字母开头，是对外公开的
	fmt.Println(outcount)

	//pricount := testPackage.mycounter(11)	//小写字母开头，是未对外公开的, 会编译错误

	fmt.Println(testPackage.New())	//工厂函数命名为New是一个go的一个习惯, 这种方法可以用到了未公开的标识符

	testuser := testPackage.TestUser{
		Name: "perkyong",
		//email是不可见的
	}
	fmt.Println(testuser)

	duser := testPackage.DUser{
		Level: 100,	//只能访问公开的，otherUser是未公开的，无法访问
		//无法访问Name，Email
	}
	//但是可以通过这种方法访问
	duser.Name = "aaaa"
	duser.Email = "163.com"
	fmt.Println(duser)

	//或者通过方法也可以访问，
	duser.Setname("zhenzhen")
	fmt.Println(duser)
}
