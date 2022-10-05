package main

import (
	"fmt"
)

func main() {
	var randomValues interface{}
	_ = randomValues

	randomValues = "Jl sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"airel", "nanda"}

	fmt.Printf("output random value : %+v\n", randomValues)

}
