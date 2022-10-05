package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	defer catchErr()

	// var password string
	var username string
	fmt.Print("Masukkan username: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password: ")
	password, _ := gopass.GetPasswdMasked()

	if valid, err := validateAccount(username, password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validateAccount(username string, password []byte) (string, error) {
	p := string(password[:])
	if username != "user" || p != "123" {
		return "", errors.New("Username/Password is wrong!")
	}

	return "Succesfully logged in!", nil
}
