package main

import "fmt"

// agar kita dapat memberikan nilai kepada field yang
// terdapat pada sebuah struct, kita perlu terlebih dahulu
// menyimpan tipe data dari struct kepada sebuah variable.
type Employee struct {
	name     string
	age      int
	division string
}

func main() {

	// Kita membuat sebuah struct dengan nama Employee yang memiliki 3 field.
	// Lalu kita membuat sebuah variable bernama employee yang
	// memiliki tipe data berupa struct Employee.
	var employee Employee
	// Kemudian kita memberikan nilai- nilai kepada field-field
	// pada struct Employee dengan cara mengakses field-field
	// melewati variable yang menyimpan tipe data struct Employee.
	employee.name = "Airel"
	employee.age = 23
	employee.division = "Curriculum Developer"

	var employee2 = Employee{name: "Ananda", age: 33, division: "Finance"}

	fmt.Printf("Employee 1 : %+v\n", employee)
	fmt.Printf("Employee 2 : %+v\n", employee2)
}
