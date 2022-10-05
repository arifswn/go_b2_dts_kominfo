package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var err error

	// jika error
	number, err = strconv.Atoi("123GH")
	if err == nil {
		fmt.Println("output : ", number)
	} else {
		fmt.Println("error nya : ", err.Error())
	}

	// jika benar
	number, err = strconv.Atoi("123")

	if err == nil {
		fmt.Println("output : ", number)
	} else {
		fmt.Println("error nya : ", err.Error())
	}

}
