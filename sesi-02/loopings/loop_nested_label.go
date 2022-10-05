package main

import "fmt"

func main() {
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
