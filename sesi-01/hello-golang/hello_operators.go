package main

import "fmt"

func main() {

	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourCondition bool = 11 <= 11

	fmt.Println("first condition : ", firstCondition)
	fmt.Println("second condition : ", secondCondition)
	fmt.Println("third condition : ", thirdCondition)
	fmt.Println("four condition : ", fourCondition)

	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t (%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t (%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}
