package main

import "fmt"

func main() {
	greet("Arif Setiawan", 23)
}

func greet(name string, age int8) {
	fmt.Printf("hello there ! my name is %s and i'm %d years old", name, age)
}
