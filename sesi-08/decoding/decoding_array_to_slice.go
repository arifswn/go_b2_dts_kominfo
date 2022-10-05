package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `jsong:"age"`
}

func main() {

	var jsonString = `
		[
			{
				"full_name" : "Airel Jordan",
				"email" : "airel@mail.com",
				"age" : 23
			},
			{
				"full_name" : "Ananda",
				"email" : "ananda@mail.com",
				"age" : 23	
			}
		]
	`
	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index : %d :  %+v\n", i+1, v)
	}
}
