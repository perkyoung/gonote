package main

import (
	"encoding/json"
	"fmt"
)

//如果要处理的json文档是以string的形式出现
// json.Unmarshal

//注意里面两个的名字一样，其中一个是类型名，另外一个是字段的name，所以相同是没关系的
type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct{
		Home	string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home" : "333.333.",
		"cell" : "dfadfasd"
	}
}`


func main() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(c)


	//有时候，无法为json的格式生命一个结构类型，而是需要更佳灵活的方式来处理json文档
	//可以将json文档解码到map变量中。
	//可以使用任意类型的值作为给顶键的值。虽然节省了定义结构的时间，但是访问的时候确实比较麻烦一些
	var d map[string]interface{}
	err = json.Unmarshal([]byte(JSON), &d)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(d)

	//因为每个值的类型都是interface{}，所以必须将值转化为合适的类型，才能处理这个值.
	//所以使用的时候可以做一下取舍
	fmt.Println(d["contact"].(map[string]interface{})["cell"])
}
