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

	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Airel", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 25}, division: "Finance"},
		{person: Person{name: "Mailo", age: 21}, division: "Marketing"},
	}

	// Kemudian kita juga dapat me-looping slice of struct
	// tersebut dengan range loop.
	for _, v := range employee {
		fmt.Printf("hasilnya : %+v\n", v)
	}
}
