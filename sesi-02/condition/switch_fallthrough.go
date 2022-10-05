package main

import "fmt"

func main() {

	// switch with fallthrough keyword
	// melanjutkan pengecekan kepada case selanjutnya
	// walaupun suatu cas telah terpenuhi kondisinya
	var score = 3
	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	case score < 5:
		fmt.Println("it is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

}
