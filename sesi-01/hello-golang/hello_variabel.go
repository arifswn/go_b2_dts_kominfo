package main

import "fmt"

func main() {

	var name string = "Arif"
	var age int = 23

	// variabel with data type
	fmt.Println("Nama saya adalah =>", name)
	fmt.Println("Umur saya adalah =>", age)

	// variabel without data type
	fmt.Printf("%T, %T", name, age)
	fmt.Println()

	//variabel without data type (shot declaration)
	my_name := "Arif Setiawan"
	my_age := 23

	fmt.Printf("%T, %T", my_name, my_age)
	fmt.Println()

	//multiple declaration
	var student1, student2, student3 string = "satu", "dua", "tiga"
	var first, second, third int

	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)
}
