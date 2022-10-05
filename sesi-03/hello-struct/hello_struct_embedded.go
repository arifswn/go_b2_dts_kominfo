package main

import (
	"fmt"
)

// Struct juga dapat mengandung tipe data struct lainnya
//  dengan menjadikannya sebuah field

// Terdapat 2 tipe data struct dengan nama Person dan Employee.
type Person struct {
	name string
	age  int
}

// Lalu tipe data employee mengandung field person yang memiliki tipe data struct Person
type Employee struct {
	division string
	person   Person
}

func main() {
	// Terdapat 2 tipe data struct dengan nama Person dan Employee.
	// Lalu tipe data employee mengandung field person yang memiliki tipe data struct Person
	var employee1 = Employee{}
	employee1.person.name = "Airel"
	employee1.person.age = 23
	employee1.division = "Curriculum Developer"

	fmt.Printf("hasil nya : %+v\n", employee1)
}
