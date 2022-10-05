package main

import (
	"fmt"
)

func main() {

	// Cara untuk membuat closure yang dapat disimpan
	// sebagai variable

	// Jika kita perhatikan, terdapat sebuah variable
	// bernama evenNumbers yang mengandung sebuah closure
	// berfungsi untuk mengumpulkan angka-angka genap
	// pada sebuah slice of int.
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	// Jika suatu variable mengandung sebuah closure,
	//  maka variable tersebut memiliki sifat
	// seperti closure yang dikandungnya.

	// Cara untuk memanggil closure tersebut,
	// kita perlu memanggil variable yang mengandungnya.

	// Jika kita perhatikan pada gambar pertama,
	// cara kita memanggil closure nya adalah dengan
	// memanggil variable evenNumbers.

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}
	fmt.Println(evenNumbers(numbers...))
}
