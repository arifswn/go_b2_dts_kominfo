package main

import "fmt"

func main() {

	c := make(chan string)

	students := []string{"airell", "mailo", "indah"}

	// Sekarang kita akan mengubah codingan kita dari halaman sebelumnya.
	//  Kita akan menggunakan range loop untuk melooping
	// data-data yang kita simpan pada sebuah
	// slice of string  sekaligus memanggil 3 Goroutine

	for _, v := range students {
		// sekarang function introduce kita ganti dengan sebuah function
		// anonymous yang kita letakkan di dalam range loop.

		// Lalu untuk proses penerimaan datanya,
		// kita membuat function bernama print
		// yang akan menerima data melalui channel c.

		// kita tidak ingin menampung sebuah kiriman data dari channel,
		//  maka kita bisa langsung menuliskan operator nya secara langsung <- c.

		// Jika kita jalankan maka hasilnya akan sama seperti sebelumnya,
		// yaitu di tiap eksekusi hasilnya akan berbeda-beda.

		go func(student string) {
			fmt.Println("student", student)
			result := fmt.Sprintf("hai, my name is %s", student)
			c <- result
		}(v)
	}

	for i := 0; i < 3; i++ {
		// fmt.Printf("student : %+v\n", <-c)
		print(c)
	}

	close(c)
}

func print(c chan string) {
	fmt.Println("hasil : ", <-c)
}
