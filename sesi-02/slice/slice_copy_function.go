package main

import "fmt"

func main() {
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// digunakan untuk copy seluruh element pada sebuah slice
	// ke dalam slice lainnya
	// disaat kita meng-copy sebuah slice, maka
	// seluruh element pada slice lainnya akan ter-replace
	// oleh element yang di copy
	nn := copy(fruits1, fruits2)

	fmt.Println("fruits 1 => ", fruits1)
	fmt.Println("fruits 2 => ", fruits1)
	fmt.Println("copied elements => ", nn)
}
