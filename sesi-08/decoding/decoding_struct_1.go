package main

import (
	"encoding/json"
	"fmt"
)

// pada struct Employee yang kita telah kita buat.
//  Terdapat sebuah format tulisan
// seperti `json:”full_name”`, `json:”email”`, dan `json:”age”`

// Tulisan-tulisan tersebut disebut sebagai tag.
// Tag kita gunakan ketika kita ingin mendecode
//  data-data sepertiJSON, form data, hingga xml
// kemudian kita simpan datadecode tersebut
// kepada field-field struct nya
type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `jsong:"age"`
}

func main() {
	// Pada sesi kali ini, kita akan
	// mempelajari cara mendecodedata JSON kepada sebuah struct

	// kita membuah data JSON sederhanamenggunakan tanda backtick `
	var jsonString = `
		{
			"full_name" : "Airel Jordan",
			"email" : "airel@mail.com",
			"age" : 23
		}
	`
	var result Employee

	// kita menggunakan function json.Unmarshal untuk
	// mendecode data JSON kepada struct Employee.

	// Argumen pertama dari function json.Unmarshal
	// menerimas ebuah nilai dengan tipe data slice of byte.

	// Pada argument pertama itulah kita meletakkan
	// data JSONnya tetapi harus kita ubah terlebih dahulu
	// menjadi tipedata slice of byte.

	// Kemudian pada argumen kedua, kita meletakkan pointer
	// dari variable result agar setelah data
	// JSON berhasil didecode, datanya akan disimpan kepada variable result.

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)
}
