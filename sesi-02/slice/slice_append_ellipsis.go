package main

import "fmt"

func main() {
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	// jika ingin memasukan seluruh element dalam suatu array
	// ke dalam array lainnya
	// kita bisa menggunakan tanda ellipsis (...)
	fruits1 = append(fruits1, fruits2...)
	fmt.Printf("nama buah : %#v", fruits1)
}
