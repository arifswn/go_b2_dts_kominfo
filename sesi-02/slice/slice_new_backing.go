package main

import (
	"fmt"
)

func main() {
	// Slice (Creating a new backing array)
	// Ketika kita ingin mendapatkan element-element dari slice yang sudah ada,
	// namun kita juga ingin membuat backing array yang baru,
	// maka kita dapat menggunakan fungsi append untuk melakukannya
	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}

	// Variable newCars ingin mendapatkan element-element dimulai dari index ke-0
	// hingga index ke-1 dari variable cars.
	// Namun newCars mendapatkan element-elementnya
	// dengan menggunakan fungsi append
	// walaupun masih menggunakan slicing di dalam fungsi append nya.

	// Ketika index ke-0 dari cars dirubah,
	// maka newCars tidak ikut terubah
	//  dikarenakan mereka tidak memiliki backing array yang sama.
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "nissan"

	fmt.Println("cars : ", cars)
	fmt.Println("newCars : ", newCars)
}
