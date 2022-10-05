package main

//  Karena function greet dan Greet merupakan
// function-funtion dari package helpers,

// maka dari itu kita perlu mengimport nya terlebih
//  dahulu dari package helpers.
//  Maksud dari penulisan hello-exported/helpers”

// kita mencoba untuk meng-import package helpers
// yang berada dalam package hello-expoted
// yang merupakan nama dari go modules
// pada project kita kali ini
import "hello-exported/helpers"

// kita coba untuk memanggil function
// greet dan Greet dari package helpers pada file main.go
func main() {

	// cara kita mengakses function-function adalah
	// menuliskan nama package terlebih dahulu kemudian
	// diikuti dengan tanda titik lalu
	//  nama function nya ( helpers.Greet() ).

	// helpers.Greet()
	// helpers.greet()

	// output :
	// ./main.go:25:10: greet not exported by package helpers
	// dimana terjadi error pada saat compile time.
	// Error ini terjadi karena function greet
	// bukan merupakan suatu function
	// yang ter-eksport dari package helpers
	//  karena penulisan function nya dimulai
	// dari huruf kecil, maka otomatis
	// function greet tidak ter-eksport.Exported & Unexported- Sesi 3

	// Jika kita lihat kembali pada file method.go
	// yang berada dalam folder helpers,
	// terdapat sebuah method bernama Invokegreet
	// yang digunakan untuk memanggil function greet
	// yang tidak ter-eksport dari package helpers.

	// #########################

	// Maka dari itu, jika kita ingin memanggil function
	// greet dari file main.go, maka kita perlu memanggil
	// method Invokegreet.

	helpers.Greet()

	// Lalu variable tersebut akan memiliki tipe data
	//  dari struct Person yang berasal dari package helpers.
	//  kita memanggil method Invokegreet dan tidak terjadi error apapun.
	var person = helpers.Person{}
	person.Invokegreet()

}
