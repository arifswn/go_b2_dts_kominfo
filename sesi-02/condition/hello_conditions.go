package main

import "fmt"

func main() {
	var currentyear = 2021

	// conditions (temporary variabel)
	// membuat suatu variabel yang dimana variabel tersebut hanya bisa di akses
	// dan digunakan pada scope block dari suatu kondisional
	if age := currentyear - 1998; age < 17 {
		fmt.Printf("kamu belum boleh membuat kartu sim => %d", age)
	} else {
		fmt.Println("kamu sudah boleh membuat kartu sim")
	}

}
