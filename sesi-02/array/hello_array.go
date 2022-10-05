package main

import (
	"fmt"
)

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"arif", "andi", "adam"}
	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)

}
