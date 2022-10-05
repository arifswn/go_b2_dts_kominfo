package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "airel",
		"age":     "23",
		"address": "jl sudirman",
	}

	fmt.Println("before deleting:", person)
	delete(person, "age")
	fmt.Println("after deleting:", person)

}
