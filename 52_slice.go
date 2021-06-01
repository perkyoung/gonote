package main

import (
	"fmt"
)

//切片本身是一个非常小的对象，主要有三个字段：指针，长度，容量；是对数组的抽象
//切片，动态数组，可以按需自动增长和缩小，用于管理数据集合
//底层也是连续的，很方便索引，迭代以及垃圾回收
//底层用append来增长，对切片再次切片来进行缩小
func main() {

	//创建字符串切片
	slice1 := make([]string, 5)	//默认长度和容量都是5,每个元素是string
	fmt.Println(len(slice1))	//5
	fmt.Println(cap(slice1))	//5

	slice2 := make([]string, 3, 5)	//指定了长度是3，容量是5, 注意不能超越长度越界访问
	slice2[2] = "aaaaaa"
	//slice2[4] = "aaaaaa" 虽然容量是5，但是不能越界访问
	slice2[1] = "bbbbb"

	slice2 = append(slice2, "ccccc")	//大脑中应该出现一个存储结构的图景，此时扩展了切片,长度增加了，容量依然没变
	fmt.Println(len(slice2))	//4
	fmt.Println(cap(slice2))	//5

	slice3 := slice2
	fmt.Println(len(slice3))	//4
	fmt.Println(cap(slice3))	//5
	//输出地址，是相同的，目前暂时是浅拷贝
	fmt.Printf("slice2 addr is %p\n", &slice2[0])
	fmt.Printf("slice3 addr is %p\n", &slice3[0])
	//继续深入，将slice3 继续追加超过容量，果然地址不同了，也就是从新给slice3分配了内存
	slice3 = append(slice3, "dddd", "eeee")
	fmt.Printf("slice2 addr is %p\n", &slice2[0])
	fmt.Printf("slice3 addr is %p\n", &slice3[0])
	//重新分配内存后slice3的len ，cap分别是多少呢？
	fmt.Println(len(slice3))	//6
	fmt.Println(cap(slice3))	//10 cap比len大了

	//除了make，还有一种创建切片的方法，就是字面量，直接赋值
	// 相当于创建了len 3, cap 3 的切片,是不是很眼熟，对，很像数组的初始化，但是一定要记住，有常量的是数组，没常量的才是切片
	stringslice := []string{"aaa", "bbb", "ccc"}
	fmt.Println(len(stringslice))	//3
	fmt.Println(cap(stringslice))	//3

	//那么用字面量的形式如何指定长度呢？使用索引声明切面
	stringslice2 := []string{99:""}
	fmt.Println(len(stringslice2))	//100
	fmt.Println(cap(stringslice2))	//100
	//回顾数组与切片的不同
	array := [3]int{1,3,4}	//这是数组
	slice := []int{1,3,4}	//这是切片，动态数组,len, cap, pointer
	fmt.Println(len(array))
	fmt.Println(len(slice))

	//创建nil切片，pointer = nil, len = 0, cap = 0
	var sliceNil []int	//在描述一个不存在的切片的时候很有用, 比如函数想返回一个切片，但是发生异常的时候，其实不太懂
	fmt.Println(sliceNil)
	//能否调用append，len，cap等，可以的
	sliceNil = append(sliceNil, 10, 4, 5, 6, 7, 10)
	fmt.Println(len(sliceNil))	// 6 对切片底层数据结构进行扩展，貌似比较保守一些
	fmt.Println(cap(sliceNil))	// 6

	//创建一个空切片
	slice0 := make([]int, 0)
	fmt.Printf("pointer is %p\n", &slice0)
	slice01 := []int{}
	fmt.Printf("pointer is %p\n", &slice01)


	//切片的性质都知道了，下面来看看如何在程序中使用切片
	//切片之所以叫切片，就是因为创建一个新的切片就是把底层数组切出一部分
	slice4 := []int{10, 20, 30, 40, 50}
	fmt.Println(len(slice4))	//5
	fmt.Println(cap(slice4))	//5
	slice5 := slice4[1:3]
	fmt.Println(len(slice5))	//2
	fmt.Println(cap(slice5))	//4
	//fmt.Println(slice5[2]) 这样写肯定越界了，不能超过len, 要想访问len后面的，必须将他合并到len里
	//对比一下第一个数组的地址, 果然，slice5的第一个元素就是slice4的第二个元素，他们共享了底层数组, 只是不同的切片看到了不同的部分而已
	fmt.Printf("pointer is %p\n", &slice4[0])	//0xc00008e060
	fmt.Printf("pointer is %p\n", &slice4[1])	//0xc00008e068
	fmt.Printf("pointer is %p\n", &slice5[0])	//0xc00008e068
	//既然是共享的，那么尝试修改一下
	slice4[1] = 100
	fmt.Println(slice4[1])	//100
	fmt.Println(slice5[0])	//也是100，都变了

	//有个很诡异的地方，现在slice4, slice5是共享底层数据的，修改一个，会同时修改另一个，如果用append，并且保证cap足够，会覆盖还是重新分配呢？
	slice5 = append(slice5, 88)	//会扩展继续扩展一个int
	fmt.Println(len(slice5))	//3
	fmt.Println(cap(slice5))	//4
	//下面两个元素对应的位置应该是一样的，就看有没有同时修改了。
	fmt.Println(slice5[2])	//88
	fmt.Println(slice4[3])	//88 果然都改了，觉得很坑，这样的话，工程代码很容易出问题的，这个肯定不是最佳实战


	slice6 := []int{10,20,30,40}	//len 4, cap 4
	newslice := append(slice6, 50)	//这个肯定会重新分配新的空间的, 因为容量不够了
	fmt.Println(cap(newslice))	//8 成倍增长， 一旦超过1000个，增长因子会变成1.25

	//刚才提到了，append的时候很容易不小心写出坑，所以现在引入创建切片时的第三个参数，来控制容量
	source := []string{"a", "b", "c", "d", "e"}
	target := source[2:3:4]	//len = 3 - 2 = 1; cap = 4 - 2 = 2	控制了len，又控制了cap
	fmt.Println(len(target))
	fmt.Println(cap(target))
	//所以这个也算是一个最佳实战，因为经常存在底层共享了数组，但是自己却不自知的情况
	//控制len和cap相等，一旦append就会重新分配底层数组，是个不错的方法
	target2 := source[2:3:3]
	fmt.Println(len(target2))
	fmt.Println(cap(target2))

	//追加切片
	//append(s1, s2....)

	//迭代切片
	for index, value := range source {	//注意这里每个元素的副本, 而不是对该元素的引用
		fmt.Println(index, value)
		//fmt.Printf("X%", &value)	//输出的值是相同的，如果想获取元素的地址，可以使用 &source[index]
	}

	//多维切片
	//回顾一下，切片是抽象的数据结构，定义一个切片需要3个值： pointer, len, cap; pointer 指向了连续的数据块。
	//那多维的切片呢？本质就是包含多个切片的切片；所以多维切片的pointer指向的是连续的切片对象（pointer, len, cap), 每个切片对象
	//的pointer指向的是连续的内存块；
	multislice := [][]int32{{10}, {20, 30}, {40,50,60}}
	fmt.Printf("pointer is %p\n", &multislice[0][0])
	//fmt.Printf("pointer is %X\n", &multislice[0][1]) 这个数组越界，因为没有在len范围内
	fmt.Printf("pointer is %p\n", &multislice[1][0])
	fmt.Println("len ", len(multislice[0]))	//1
	fmt.Println("cap ", cap(multislice[0]))	//1
	fmt.Println("len ", len(multislice[1]))	//2
	fmt.Println("cap ", cap(multislice[1]))	//2
	fmt.Println("len ", len(multislice[2]))	//3
	fmt.Println("cap ", cap(multislice[2]))	//3
	//重点看下在内存中不同slice的数据块是否连续
	fmt.Printf("pointer is %p\n", &multislice[0][0])	//C0000141A8
	fmt.Printf("pointer is %p\n", &multislice[1][0])	//C0000141B0
	fmt.Printf("pointer is %p\n", &multislice[2][0])	//C0000141C0
	fmt.Printf("重新分配前第一个切片的地址 pointer is %p\n", &multislice[0])		//0xc000080050
	//还真是连续的，那么如果append第一个slice再看效果呢
	multislice[0] = append(multislice[0], 70, 80)
	fmt.Printf("重新分配后第一个切片的地址 pointer is %p\n", &multislice[0])		//0xc000080050
	//很显然，重新分配的是底层的数据，而这个切片的三个字段(pointer, len, cap)所在的位置并没有变化

	fmt.Printf("pointer is %p\n", &multislice[0][0])	//C0000141F8	地址变了，重新分配了底层数组
	fmt.Printf("pointer is %p\n", &multislice[1][0])	//C0000141B0	没变
	fmt.Printf("pointer is %p\n", &multislice[2][0])	//C0000141C0	没变

	//输出每个切片首字节的地址，以后依次相差24个字节，也就是pointer(8字节）, len(8字节）, cap(8字节）
	fmt.Printf("pointer is %p\n", &multislice[0])		//0xc000080050
	fmt.Printf("pointer is %p\n", &multislice[1])		//0xc000080068
	fmt.Printf("pointer is %p\n", &multislice[2])		//0xc000080080

	//总结，关于多维slice，每个slice在内存中是连续的，
	//每个slice对应的数据在内存中初始化的时候也是连续的，这么做是出于性能的考虑，如果对某个slice做append操作，会重新开辟空间

	//函数之间传递切片，
	//这个很高效，因为切片本身（pointer, len, cap) 只占用24个字节（64位机器）,复制的时候，只复制了这三个值而已，并没有动底层的数据块
	//所以非常高效

}
