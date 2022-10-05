package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = []string{"airel", "jordan"}
	var printMsg = greet("heii", name)
	fmt.Println(printMsg)
}

func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)
	return result
}
