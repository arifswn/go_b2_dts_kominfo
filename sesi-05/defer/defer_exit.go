package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("invoke with defer")
	fmt.Println("before exiting")
	os.Exit(1)
}
