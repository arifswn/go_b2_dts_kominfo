package main

import "fmt"

func main() {
	var i = 0

	for {
		fmt.Println("angka", i)
		i++
		if i == 3 {
			break
		}
	}
}
