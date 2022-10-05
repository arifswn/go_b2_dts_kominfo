package main

import (
	"fmt"
	"strings"
)

func main() {
	profile("airell", "pasta", "ayam geprek", "ikan roa", "sate padang")
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("hello there!! i'm ", name)
	fmt.Println("i really love to eat ", mergeFavFoods)
}
