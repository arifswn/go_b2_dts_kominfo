package main

import (
	"fmt"
)

func main() {

	// ntuk menampilkan nilai asli (bukan alamat memori)
	//  yang dikandung oleh secondNumber dengan cara
	// memberikan tanda asterisk *.

	// Perlu diingat disini bahwa karena
	// variable secondNumber merupakan sebuah variable pointer
	// yang mengandung alamat memori dari firstNumber,
	// maka ketika kita ingin mendapatkan nilai asli
	// yang dikandung oleh sebuah variable pointer,
	// kita perlu menggunakan tanda asterisk *
	var firstNumber int = 4
	var secondNumber *int = &firstNumber
	_, _ = firstNumber, secondNumber

	fmt.Println("first number (value) : ", firstNumber)
	fmt.Println("first number (memori address) : ", &firstNumber)

	fmt.Println("second number (value) : ", secondNumber)
	fmt.Println("second number (memori address) : ", &secondNumber)

}
