package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// Slice juga dapat dikombinasikan dengan tipe data struct,
// cara penulisannya mirip seperti slice of map
//  yang telah kita bahas pada materi sebelumnya.

func main() {
	// Terdapat sebuah variable bernama people
	// yang memiliki tipe data slice of struct
	//  dan tipe data struct yang dimasukkan
	// kedalam slice nya adalah struct bernama Person.

	var poeple = []Person{
		{name: "Airel", age: 23},
		{name: "Ananda", age: 22},
		{name: "Mailo", age: 44},
	}

	// Kemudian kita juga dapat me-looping slice of struct
	// tersebut dengan range loop.
	for _, v := range poeple {
		fmt.Printf("hasilnya : %+v\n", v)
	}
}
