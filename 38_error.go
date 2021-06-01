package main

import (
	"fmt"
)

type myError struct {

}

func (my myError) Error() string {
	fmt.Println("myError")
	return string("myError")
}

func main()  {
}
