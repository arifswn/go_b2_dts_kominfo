package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "airel",
		"age":     "23",
		"address": "jl sudirman",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
