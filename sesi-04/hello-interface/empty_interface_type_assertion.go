package main

import (
	"fmt"
)

func main() {
	var v interface{}
	v = 2
	v = v.(int) * 3 //dibolehkan

	//rekomendasi
	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	fmt.Printf("output v interface : %d\n", v)
}
