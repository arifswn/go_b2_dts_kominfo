package main

import "fmt"

func main() {

	//default type data
	// var first = 89
	// var second = -12

	//inisialisai type data
	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("tipe data first : %T\n", first)
	fmt.Printf("bilangan second : %T\n", second)

	//type data decimal number
	var decimalNumber float32 = 3.63

	fmt.Printf("decimal number : %f\n", decimalNumber)
	fmt.Printf("decimal number : %.3f\n", decimalNumber)

	//type data boolean
	var condition bool = true
	fmt.Printf("is it permitted ? %t\n", condition)

	//type data string
	var message string = "Pagi"
	fmt.Printf("hello %s \n", message)

	//type data string
	//deklarasi string juga bisa dengan tanda grave accent/backticks (``)
	//keistimewaan string yang dideklarasikan menggunakan backtics adalah membuat semua karakter didalamnya tidak diescape

	var my_message = `selamat datang di hactiv8.
	salam kenal :).
	mari belajar "scalable web service with go."`

	fmt.Printf(my_message)

}
