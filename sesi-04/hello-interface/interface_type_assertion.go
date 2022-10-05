package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	fmt.Printf("type of c1 : %T\n", c1)
	fmt.Printf("type of r1 : %T\n", r1)
	fmt.Println()

	// page 4
	fmt.Println("circle area", c1.area())
	fmt.Println("circle perimeter", c1.perimeter())
	fmt.Println()

	fmt.Println("rectangle area", r1.area())
	fmt.Println("rectangle perimeter", r1.perimeter())
	fmt.Println()

	print("rectangle", c1)
	fmt.Println()
	print("perimeter", r1)
	fmt.Println()

	// Interface (Type assertion)
	// Ketika struct circle menambahkan satu method selain
	// dari method-method yang didefinisikan oleh interface shape,
	// maka variable c1 tidak dapat menggunakan method baru tersebut.
	c1.(circle).volume()

	// Kita juga dapat memeriksa apakah sebuah type assertion
	// yang kita gunakan berhasil atau tidak.
	// Caranya dengan memberikan 2 variable opsional yang
	// menjadi penampung dari hasil yang akan
	// dikembalikan oleh type assertionnya.
	value, ok := c1.(circle)

	// Variable value akan menerima tipe data asli dan
	// nilai yang sudah diberikan kepada field dari
	// struct circle jika type assertion berhasil,
	// dan variable ok akan menerima nilai true jika
	// type assertion berhasil dan nilai false jika type assertion gagal.
	if ok == true {
		fmt.Printf("circle value : %+v\n", value)
		fmt.Printf("circle volume : %f\n", value.volume())
	}
}
