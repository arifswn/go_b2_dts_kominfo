package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "makan"
	str2 := "mànca"

	// Ketika kita hendak mencari jumlah karakter nya,
	// dan bukan jumlah bytenya,
	// maka kita perlu merubah string tersebut
	// menjadi rune terlebih dahulu.

	// Tipe data rune merupakan tipe data alias dari int32.
	// Kita dapat menggunakan function RuneCountInString
	// dari package utf8 untuk merubah string menjadi rune
	// sekaligus mencari panjang karakternya.

	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

	str := "mànca"
	for i, s := range str {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}
}
