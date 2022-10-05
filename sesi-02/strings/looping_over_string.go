package main

import "fmt"

func main() {
	city := "jakarta"

	// Pada looping tersebut, kita melakukannya dengan pendekatan byte-per-byte,
	//  atau dapat diartikan kita meloopingstring “Jakarta”
	// berdasarkan byte nya.
	// Kode city[i] tidak akan menghasilkan karakternya,
	//  melainkan byte nya.

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}

	var city1 []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println(string(city1))

	str1 := "makan"
	str2 := "mànca"

	// String “mânca” sebetulnya memiliki 5 karakter,
	//  namun angka 6 pada gambar diatas bukan bermaksud untuk
	// menunjukkan panjan karakter dari string “mânca”,
	// namun bermaksud untuk menunjukkan jumlah byte nya.

	// Jika kita perhatikan pada string “mânca”,
	// terdapat sebuah karakter yang tidak didukung pada table ASCII Code,
	//  yaitupada karakter “â” yang merupakan sebuah accented-character.
	// Karakter “â” memiliki 2 byte,
	// inilah yang menyebabkan keluarnya angka 6 pada terminal
	// ketika kita mencoba untukmenampilkan jumlah byte pada string “mânca”.

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

}
