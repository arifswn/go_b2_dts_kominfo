package main

import "fmt"

func main() {
	// Channel merupakan sebuah mekanisme untuk
	// Goroutine saling berkomunikasi dengan Goroutine lainnya.
	//  Maksud berkomunikasi disini adalah proses pengiriman
	// dan pertukaran data antara Goroutine
	//  satu dengan Goroutine lainnya.

	// Untuk membuat sebuah Goroutine,
	// kita memerlukan function make yang merupakan
	// sebuah built-in function dari bahasa Go.

	// Variable c merupakan sebuah variable yang memiliki
	// tipe data channel of string.

	// Maksud channel of string disini adalah
	// sebuah channel yang memiliki tipe data string.

	// Lalu keyword chan merupakan keyword untuk membuat channel.

	c := make(chan string)

	// Kita membutuhkan operator dari channel agar channel tersebut
	//  dapat dijadikan sebagai alat untuk berkomunikasi
	// antara Goroutine dengan yang lainnya

	go introduce("airell", c)
	go introduce("nanda", c)
	go introduce("mailo", c)

	//mengirim data kepada channel
	// c <- value

	// Tanda <- merupakan sebuah operator dari channel
	// untuk proses pengiriman data dari Goroutine
	// satu dengan yang lainnya.

	// result := <-c

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	// Dan yang terakhir maksud dari penulisan
	//  result := <-c adalah kita menerima data dari channel c
	// dan data tersebut kita simpan di dalam variable result.

	// Perlu diingat disini bahwa proses pengiriman
	// dan penerimaan data dari channeI bersifat synchronous.
	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("hai, my name is %s", student)
	c <- result

	// di tiap eksekusi memiliki menampilkan hasil yang berbeda-beda.
	//  Ini terjadi karena seperti yang pernah kita bahas
	// pada materi sebelumnya bahwa,

	// Goroutine bekerja secara asynchronous
	//  yang dimana Goroutine satu dengan
	// yang lainnya tidak akan saling menunggu

	// Dan Goroutine merupakan cara agar
	// kita dapat membuat concurrency pada bahasa Go,
	// dan dalam concurrency,

	// kita tidak akan tahu mana yang akan selesai tereksekusi
	// terlebih dahulu.

	// Ini lah yang menjadi penyebab mengapa di tiap eksekusi
	// menampilkan hasil yang berbeda-beda
}
