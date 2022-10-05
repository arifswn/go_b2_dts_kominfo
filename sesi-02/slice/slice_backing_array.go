package main

import "fmt"

func main() {
	var fruits1 = []string{"apple", "banana", "durian", "mango", "startfruit"}

	// Ketika kita mencoba untuk mendapatkan  beberapa element
	// dari sebuah slice yang sudah ada
	// dengan cara melakukans licing,

	// maka Go tidak akan membuat suatu backing array baru
	// melainkan slice tersebut akan berbagi backing array
	// yang sama dengan slice yang sudah ada
	var fruits2 = fruits1[2:4]
	fruits2[0] = "rambutan"

	// Variable fruits2 melakukan slicing terhadap fruits1
	// untuk mendapatkan element dari index ke-2 sampai ke-3
	// dari fruits1 atau yang berarti buah durian dan banana.

	// Setelah itu fruits2 mencoba untuk mengganti
	// buah yang berada pada index ke-0 (buah durian) menjadi buahrambutan
	fmt.Println("fruits1 => ", fruits1)
	fmt.Println("fruits2 => ", fruits2)

}
