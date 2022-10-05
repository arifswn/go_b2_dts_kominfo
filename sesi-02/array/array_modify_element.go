package main

import (
	"fmt"
)

func main() {
	var fruits = [4]string{"apel", "pisang", "mangga", "arif"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Printf("%#v\n", fruits)

}
