package main

import "fmt"

func main() {
	// membuat sebuah fungsi slice dengan fungi make
	// argumen pertama diberikan pada tipe dari slice nya
	// argumen kedua adalah panjang dari slice nya
	var fruits = make([]string, 3)
	_ = fruits

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"
	// output 3 string buah
	fmt.Printf("nama buah : %#v", fruits)
}
