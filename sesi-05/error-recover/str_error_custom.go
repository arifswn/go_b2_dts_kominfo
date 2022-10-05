package main

import (
	"errors"
	"fmt"
)

func main() {
	var password string
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("status password anda : ", valid)
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)
	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "valid password", nil
}
