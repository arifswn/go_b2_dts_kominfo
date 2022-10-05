package main

import "fmt"

func main() {
	var poeple = []map[string]string{
		{"name": "airel", "age": "23"},
		{"name": "nanda", "age": "20"},
		{"name": "mailo", "age": "15"},
	}

	for i, person := range poeple {
		fmt.Printf("index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
