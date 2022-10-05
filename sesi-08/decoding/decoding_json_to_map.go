package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var jsonString = `
		{
			"full_name" : "Airel Jordan",
			"email" : "airel@mail.com",
			"age" : 23
		}
	`
	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Kita juga bisa men-decode data JSON kepada
	// sebuah tipe data map.
	// Kita tidak perlu membuat tag seperti yang kita
	// lakukan pada sebuah struct.

	fmt.Println("full_name:", result["full_name"])
	fmt.Println("email:", result["email"])
	fmt.Println("age:", result["age"])
}
