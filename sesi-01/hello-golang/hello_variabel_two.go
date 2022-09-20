package main

import "fmt"

func main() {

	//multiple declaration
	var name, age, address = "Arif", 23, "Jl Dr Selamet Baundung"
	first, second, third := "1", 2, "3"

	fmt.Println(name, age, address)
	fmt.Println(first, second, third)

	var firstVariable string

	//bahasa Golang memiliki satu aturan yang cukup strict/ketat
	//yang dimana tidak ada variable yang boleh menganggur ketika sudah kita buat
	//dengan variable underscore maka kita dapat menghindari error
	//yang akan terjadi jika kita mempunyai suatu variable yang menganggur atau tidak dipakai
	_, _, _, _ = firstVariable, name, age, address

	my_first, my_second := 1, "2"
	fmt.Printf("Tipe data variabel first adalah %T \n", my_first)
	fmt.Printf("Tipe data variabel second adalah %T \n", my_second)

	//menggunakan verb
	fmt.Printf("Hallo nama ku %s, umurku adalah %d, dan aku tinggal di %s", name, age, address)
}
