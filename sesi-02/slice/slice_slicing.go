package main

import "fmt"

func main() {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pinneaple"}
	fmt.Printf("fruits 1 => %#v\n", fruits1)

	// -Variable fruits2 ingin mendapatkan element dari fruits1
	// dari index kesatu hingga index ke tiga,
	// maka dari itu cara penulisannya adalahfruits[1:4].
	var fruits2 = fruits1[1:4]
	fmt.Printf("fruits 2 => %#v\n", fruits2)

	// Variable fruits3 ingin mendapatkan element dari fruits1
	// dari index kenol hingga index terakhirnya,
	//  maka dari itu cara penulisannya adalahfruits[0:]
	// yang dimana keterangan stop boleh dihilangkan
	// jika inginmendapatkan hingga index terakhir.
	var fruits3 = fruits1[0:]
	fmt.Printf("fruits 3 => %#v\n", fruits3)

	// Variable fruits4 ingin mendapatkan element dari fruits1
	// dari index kenol hingga index kedua,
	// maka dari itu cara penulisannya adalahfruits[:3]
	// yang dimana keterangan start nya boleh dihilangkan
	// ketikaingin mendapatkan element dari index ke nol.
	var fruits4 = fruits1[:3]
	fmt.Printf("fruits 4 => %#v\n", fruits4)

	// Variable fruits5 ingin mendapatkan element dari fruits1
	// dari index 0hingga index terakhir,
	// maka cara penulisannya adalah fruits[:]
	// yang dimana jika ingin mendapatkan element seluruhnya,
	// berarti tidak perlu memberi keterangan start dan stop nya,
	// cukup memberikantanda titik dua saja [:].
	var fruits5 = fruits1[:]
	fmt.Printf("fruits 5 => %#v\n", fruits5)

}
