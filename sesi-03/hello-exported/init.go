package main

import "fmt"

// Karena file init.go merupakan file yang berada
// dalam root direktori project kita,
// maka file init.go akan dianggap sebagai bagian dari package main
func init() {
	fmt.Println("Hallo !!! ini dari function init")
}

// how to run
// go run *.go
