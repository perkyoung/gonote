package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	gResult struct {
		//标签，提供每个字段的元信息的一种机制，json文档与结构文档提供一种映射
		//如果不存在标签，会自动对应（忽略大小写）,对应不上，就是零值
		GsearchResultClass string `json:"gsearchResultClass"`
		UnescapedURL	string `json:"unescapedUrl"`
		URL string 	`json:"url"`
		VisibleURL string `json:"visibleUrl"`
		CacheURL string `json:"cacheUrl"`
		Title string `json:"title"`
		TitleNoFormatting string `json:"titleNoFormatting"`
		Content	string `json:"content"`
	}
	gResponse struct {
		ResponseData struct{
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)
func main() {
	uri := "http://ajax.googleapis/com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	resp, err := http.Get(uri)

	if err != nil {
		log.Println("error", err)
		return
	}
	defer resp.Body.Close()
	var gr gResponse
	//如果要处理来自网络或者文件的json输入流，那么一定会用到这个方法，如果是字符串json，看下一节的内容
	//NewDecoder 返回一个解码器
	//Decode接受一个interface{}类型的值做参数，注：任何类型都实现了一个空 接口，意味着Decode可以接受任何类型的值
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(gr)
}
