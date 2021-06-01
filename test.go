package main

func partition(a []int, low, high int) int {
	base := a[high]
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] < base {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partition2(a []int, low int, high int) int {
	base := a[high]
	i := low - 1
	for j:= low; j < high; j++ {
		if a[j] < base {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i+1
}

func quitSort2(a []int, low, high int) {
	if low >= high {
		return
	}
	p := partition2(a, low, high)
	quitSort2(a, low, p-1)
	quitSort2(a, p+1, high)

}

func test() {
	//再看一个字符串指针的例子
	array1 := [3]*string{new(string), new(string), new(string)}
	*array1[0] = "hello"
	*array1[2] = "world"	//注意，如果初始化array1时，第三个元素不new空间，这个地方同样也不会为world开辟空间，编译会出错
}
func main() {

	test()
	return
	arr := []int{1,4,9,7,2,5}
	len := len(arr)
	quitSort2(arr, 0, len - 1)
	for _, v := range arr {
		println(v)
	}
	return
	/*
	rand.Seed(time.Now().UnixNano())
	testData1, testData2, testData3, testData4 := make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000), make([]int, 0, 100000000)
	times := 100000000
	for i := 0; i < times; i++ {
		val := rand.Intn(20000000)
		testData1 = append(testData1, val)
		testData2 = append(testData2, val)
		testData3 = append(testData3, val)
		testData4 = append(testData4, val)
	}
	start := time.Now()
	quickSort(testData1, 0, len(testData1)-1)
	fmt.Println("single goroutine: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData1) {
		fmt.Println("wrong quick_sort implementation")
	}
	done := make(chan struct{})
	start = time.Now()
	go quickSort_go(testData2, 0, len(testData2)-1, done, 5)
	<-done
	fmt.Println("multiple goroutine: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData2) {
		fmt.Println("wrong quickSort_go implementation")
	}
	start = time.Now()
	sort.Ints(testData3)
	fmt.Println("std lib: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData3) {
		fmt.Println("wrong std lib implementation")
	}
	start = time.Now()
	timsort.Ints(testData4, func(a, b int) bool { return a <= b })
	fmt.Println("timsort: ", time.Now().Sub(start))
	if !sort.IntsAreSorted(testData4) {
		fmt.Println("wrong timsort implementation")
	}
	*/

}


func quickSort_go(a []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}
	depth--
	p := partition(a, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSort_go(a, lo, p-1, childDone, depth)
		go quickSort_go(a, p+1, hi, childDone, depth)
		<-childDone
		<-childDone
	} else {
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}
	done <- struct{}{}
}
