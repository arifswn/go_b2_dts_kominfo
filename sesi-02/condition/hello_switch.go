package main

import "fmt"

func main() {

	// hello switch
	var score = 8
	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

}
