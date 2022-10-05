package main

import "fmt"

func main() {
	// Ketika kita ingin membuat sebuah map,
	//  maka kita juga haruslangsung menginisialisasikannya.
	var person map[string]string //deklarasi
	person = map[string]string{} //inisialisasi

	// Penulisan map[string]string berarti tipe data key
	// dari map harus string dan
	// value dari map juga harus string.

	// Lalu jika kita ingin memberikan nilai
	// atau value kepada map nya,
	// maka harus diinisialisasi terlebih dahulu.

	// kita memberikan value kepada variable person
	// dengan key yang berbeda-beda atau unik.
	// Kita juga dapat mendapatkan value dari map
	// dengan cara mengakses key darimap
	// maka hasilnya akan seperti padagambar kedua di sebelah kanan.

	person["name"] = "airell"
	person["age"] = "23"
	person["address"] = "jl sudirman"

	fmt.Println("name", person["name"])
	fmt.Println("age", person["age"])
	fmt.Println("address", person["address"])
}
