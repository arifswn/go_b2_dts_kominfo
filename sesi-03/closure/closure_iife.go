package main

import (
	"fmt"
)

func main() {

	// IIFE atau singkatan dari immediately-invoked
	// function expression merupakan sebuah closure
	// yang dapat langsung tereksekusi ketika pertama kali dideklarasikan.
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	// Variable isPalindrome merupakan sebuah penampung
	// dari closure IIFE yang digunakan untuk mengetahui
	// apakah parameter yang diberikan merupakan sebuah kalimat
	// palindrome atau bukan.

	// Jika kita ingin membuat suatu closure menjadi IIFE,
	// maka kita perlu langsung memanggil closure tersebut
	// secara langsung pada saat dideklarasikan.

	// Perlu diingat bahwa kita tidak perlu lagi
	//  memanggil closure IIFE dengan tanda kurung ()
	// karena closure IIFE tereksekusi pada saat dideklarasikan.

	fmt.Println("data sesuai : ", isPalindrome)
}
