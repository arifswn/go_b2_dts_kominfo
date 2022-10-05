package main

import (
	"fmt"
)

func main() {

	var a int = 10
	fmt.Println("before : ", a)

	changeValue(&a)

	fmt.Println("after : ", a)
}

func changeValue(number *int) {
	*number = 20
}
