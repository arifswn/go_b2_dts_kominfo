package main

import (
	"fmt"
	"strings"
)

// agar kita dapat memberikan nilai kepada field yang
// terdapat pada sebuah struct, kita perlu terlebih dahulu
// menyimpan tipe data dari struct kepada sebuah variable.
type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	// Kita juga dapat menggunakan pointer pada sebuah struct
	var employee1 = Employee{name: "Airell", age: 23, division: "Curriculum Developer"}
	//  kita membuat sebuah variable bernama employee2
	// dengan tipe data pointer to a struct yang menyimpan
	// alamat memori dari variable employee1.

	var employee2 *Employee = &employee1

	fmt.Println("Employee 1 name: ", employee1.name)
	fmt.Println("Employee 2 name: ", employee2.name)

	fmt.Println(strings.Repeat("#", 20))
	// employee2 merubah nilai dari field name nya
	// menjadi “Ananda” yang semula nilai nya adalah “Airell”.
	employee2.name = "Ananda"
	// Ketika variable employee2 merubah nilai dari fieldnya,
	//  maka variable employee1 juga akan ikut terubah nilainya.
	fmt.Println("Employee 1 name: ", employee1.name)
	fmt.Println("Employee 2 name: ", employee2.name)
}
