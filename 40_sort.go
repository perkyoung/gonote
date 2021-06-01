package main

import (
	"fmt"
	"sort"
)


type ByLength []string
//实现sort所有接口
func (s ByLength) Len() int {	//快速排序的长度n
	return len(s)
}

func (s ByLength) Swap(i, j int) {	//如何交换
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {	//对比大小
	return len(s[i]) < len(s[j])
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7,4,5}
	sort.Ints(ints)
	fmt.Println("ints: ", ints)

	//检查是否已经排序好了
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted: ", s)

	fruits := []string{"peach", "banana", "kiwi"}
	fmt.Println("the len is ", len(fruits))
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
