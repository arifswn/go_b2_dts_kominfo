package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits1 = []string{"apple", "mango", "durian", "banana"}

	// Fungsi cap dapat kita gunakan untuk mengetahui kapasitas
	// dari sebuah array maupun slice
	// Ketika pertama kali kita membuat suatu slice,
	// panjang dan kapasitasnya dipastikan sama,
	//  namun dapat berubah seiring dengan slicing yang kita lakukan

	// Variablefruits1 telah disiapkan di awal dengan panjang 4
	//  dan kapasitas 4.
	// Lalu terdapat 2 variable baru bernama fruits2
	// dan fruits3 yangmelakukan slicing terhadap fruits1
	//  untuk mendapatkan 3element dari fruits1.
	fmt.Println("fruits1 cap:", cap(fruits1))
	fmt.Println("fruits1 len:", len(fruits1))

	fmt.Println(strings.Repeat("#", 20))

	// Lalu kenapa fruits2 memiliki panjang dan kapasitas
	// yang berbeda sedangkan fruits3 panjang
	var fruits2 = fruits1[0:3]

	fmt.Println("fruits2 cap:", cap(fruits2))
	fmt.Println("fruits2 len:", len(fruits2))

	fmt.Println(strings.Repeat("#", 20))

	var fruits3 = fruits1[1:]

	fmt.Println("fruits3 cap:", cap(fruits3))
	fmt.Println("fruits3 len:", len(fruits3))
}
