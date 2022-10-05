package main

import (
	"fmt"
)

func main() {
	rs := []interface{}{1, "airell", true, 2, "ananda", true}

	rm := map[string]interface{}{
		"name":   "airel",
		"status": true,
		"age":    23,
	}

	_, _ = rs, rs
	fmt.Printf("slice : %+v\n", rs)
	fmt.Printf("map : %+v\n", rm)
}
