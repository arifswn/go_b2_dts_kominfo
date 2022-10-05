package main

import (
	"fmt"
)

func main() {
	go goroutine()
	fmt.Println("Hello")
}

func goroutine() {
	fmt.Println("Hello Goroutines")
}
