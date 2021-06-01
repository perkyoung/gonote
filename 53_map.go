package main

import "fmt"

func main() {
	dict1 := make(map[string]int)
	fmt.Println(len(dict1))

	dict := map[string]int{"color": 10, "age" : 22}
	fmt.Println(len(dict))

	//map的key可以是string，int，也可以是结构体，只要能比较的，都可以，但是不能是切片，函数，或者包含切片的结构体，因为
	//切片是引用类型
	//dict3 := make(map[[]string]int) 切片不能用作key，会报错
	//切片可以作为value
	//创建一个空映射
	dict5 := map[string]string{}
	dict5["aaa"] = "bbb"
	//未初始化，相当于创建了一个nil映射, nil映射不能存储key value 会报错
	//panic: assignment to entry in nil map
	//var dict6 map[string]string
	//dict6["dfdfas"] = "dfasdfs"

	dict4 := make(map[string][]string)
	dict4["hahaha"] = []string{"name", "age", "filed"}
	fmt.Println(dict4)

	value, exist := dict4["akakaak"]
	fmt.Println(value, exist)	//如果不存在，返回的是零值，比如int 0， 字符串""， 切片[]
	if exist {	//是否存在
		fmt.Println(value)
	}

	for key, value := range dict5 {
		fmt.Println(key, value)	//返回的是键值对，不存在index，没这个概念
	}

	//删除操作
	delete(dict4, "xxx")


	//函数间传递参数不会传递完整的副本，函数内修改后，所有引用的地方都会察觉到
}
