package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Number    int
	Full_name string
	Address   string
	Job       string
	MyReason  string
}

var option int

var students = []Student{
	{Number: 1, Full_name: "Dian", Address: "Lampung", Job: "Staff", MyReason: "Learning"},
	{Number: 2, Full_name: "Dina", Address: "Jakarta", Job: "IT", MyReason: "Learning"},
	{Number: 3, Full_name: "Dini", Address: "Palembang", Job: "Finance", MyReason: "Learning"},
	{Number: 4, Full_name: "Diah", Address: "Bandung", Job: "Admin", MyReason: "Learning"},
	{Number: 5, Full_name: "Reni", Address: "Yogyakarta", Job: "Helpdesk", MyReason: "Learning"},
}

func main() {
	if os.Args != nil && len(os.Args) > 1 {
		command := os.Args[1]
		if command != "" {
			y, e := strconv.Atoi(command)
			if e == nil {
				checkAbsen(y, students)
			} else {
				fmt.Println("Inputan harus angka !!")
			}
		} else {
			command = ""
			fmt.Println("No command given")
		}
	} else {
		Menu()
	}

}

func Menu() {
	stdin := bufio.NewReader(os.Stdin)

	fmt.Println(".:: My Absen ::.")
	fmt.Print("Masukan No Absen Student : ")
	for {
		_, err := fmt.Fscan(stdin, &option)
		if err == nil {
			break
		}
		stdin.ReadString('\n')
		fmt.Print("Wrong option input. Please, try again : ")
	}

	checkAbsen(option, students)
}

func checkAbsen(c int, list []Student) {
	for _, v := range list {
		if v.Number == c {
			fmt.Printf("My Absen : %v\n", v)
		}
	}

	if c > len(list) {
		fmt.Println("Tidak Ada !!")
	}
}
