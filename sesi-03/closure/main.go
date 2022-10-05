package main

import (
	"fmt"
)

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	// var employee Employee
	// employee.name = "Airell"
	// employee.age = 23
	// employee.division = "Curriculum Developer"

	// fmt.Println("Nama : ", employee.name)
	// fmt.Println("Age : ", employee.age)
	// fmt.Println("Division : ", employee.division)

	// step 1
	var employee1 = Employee{}
	employee1.name = "Airell"
	employee1.age = 23
	employee1.division = "Curriculum Developer"

	// step 2
	var employee2 = Employee{
		name:     "Airell",
		age:      23,
		division: "Curriculum Developer",
	}

	fmt.Printf("Employee1 : %+v\n", employee1)
	fmt.Printf("Employee2 : %+v\n", employee2)
}
