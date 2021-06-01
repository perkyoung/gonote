package main

import (
	"encoding/json"
	"fmt"
)

//编码encode
//map转化为json就很方便了
func main()  {
	c := make(map[string]interface{})

	c["name"] = "perkyoung"
	c["job"] = map[string]interface{}{
		"home" : "dfad",
		"cell" : "dfadsddd",
	}

	//indent带缩进的
	//函数
	data, err := json.MarshalIndent(c, " ", "	")	//会使用反射来确定如何将map类型转换成json字符串
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(string(data))
}
