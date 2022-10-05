package main

import (
	"fmt"
)

func main() {
	// Pointer merupakan sebuah variable spesial yang digunakan
	//  untuk menyimpan alamat memori variable lainnya.

	// Sebagaicontoh sebuah variabel bertipe integer memiliki nilai 4,
	// Pointer merupakan sebuah variable spesial yang digunakan
	// untuk menyimpan alamat memori variable lainnya.

	// Sebagai contoh sebuah variabel bertipe integer
	// memiliki nilai 4, maka yang dimaksud pointer adalah
	// alamat memori dimana nilai 4disimpan,
	// bukan nilai 4 itu sendiri.

	// Variabel-variabel yang memiliki reference atau
	// alamat memori yang sama, saling berhubungan satu sama lain
	// dan nilainyapasti sama.

	// Ketika ada perubahan nilai,
	// maka akan memberikan efek kepada variabel lain
	// (yang alamat memorinya sama) yaitu nilainya juga ikut berubah.

	// Cara membuat suatu variable menjadi sebuah pointer cukup mudah.
	// Caranya adalah dengan menggunakan tanda asterisk
	// *sebelum penulisan tipe datanya.

	// Contohnya seperti 2 variable p maka yang dimaksud pointer
	// adalah alamat memori dimana nilai 4 disimpan,
	// bukan nilai 4 itu sendiri.

	// Variabel-variabel yang memiliki reference atau alamat memori
	//  yang sama, saling berhubungan satu sama lain
	// dan nilainya pasti sama.

	// Ketika ada perubahan nilai,
	// maka akan memberikan efek kepada variabel lain
	// (yang alamat memorinya sama) yaitunilainya juga ikut berubah.

	var firstNumber *int
	var secondNumber *int
	_, _ = firstNumber, secondNumber

	fmt.Println("number 1 : ", firstNumber)
	fmt.Println("number 2 : ", secondNumber)

	
}
