package main

import "fmt"

func main() {

	// Constants & Operators - Sesi 1 ConstantConstant (const) atau Konstanta merupakan jenis variable pada bahasa Go yang nilainya tidak dapat diubah
	const full_name string = "Arif Setiawan"
	fmt.Printf("hello %s \n", full_name)

	var value = 2 + 2*3
	fmt.Println("Hasil 2 + 2*3 => ", value)

	var my_value = (2 + 2) * 3
	fmt.Println("Hasil (2 + 2) * 3 => ", my_value)

}
