package main

import "fmt"

func main() {

	//deklarasi variabel dengan tipe data uint8
	var a uint8 = 10
	var b byte //byte adalah alias dari tipe dari uint8

	b = a //no error, karena byte memiliki tipe data yang sama dengan uint8
	_ = b

	fmt.Printf("hello aliase : %+v\n", b)

	//mendekelarasi tipe data alias bernama second
	//type nama_alias = nama_tipe_data
	type second = uint

	// Aliase
	// Kita juga dapat mendeklarasikan tipe data alias
	// dengan nama yang kita buat sendiri,
	// contohnya seperti padagambar diatas:
	// 1. Terdapat sebuah tipe data bernama second  yang merupakan tipe data alias dari uint.
	// 2. Variable hour memilki tipe data second yang merupakan sebuah tipe data alias.
	// 3. Ketika variable hour ingin di ketahui tipe data nya
	// dengan menggunakan flag %T,
	// maka hasil padaterminal akan tetapi menunjukan bahwa
	// tipe data dari variable hour ada uint.

	// Hal ini terjadi karena tipedata asli atau
	// underlying type dari i tipe data second adalah tipe data uint.
	var hour second = 3600
	fmt.Printf("hour type : %T\n", hour) // hour type : uint
}
