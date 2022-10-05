package main

import (
	"fmt"
	"strings"
)

func main() {

	var firstPerson string = "john"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)
	fmt.Println("secondPerson (value) : ", secondPerson)
	fmt.Println("secondPerson (memori address) : ", &secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30))
	fmt.Println()

	*secondPerson = "doe"

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)

	fmt.Println("secondPerson (value) : ", secondPerson)
	fmt.Println("secondPerson (memori address) : ", &secondPerson)

}
