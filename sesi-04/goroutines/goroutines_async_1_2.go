package main

import (
	"fmt"
	"runtime"
)

// Sekarang kita akan mempelajari sifat dari Goroutine
// yang bekerja secara asynchronous.
func main() {
	fmt.Println("main execution started")

	go firstProcess(8)
	secondProcess(8)

	// kita menggunakan function NumGoroutine dari
	// package runtime untuk mengetahui jumlah Goroutine
	// yang sedang berjalan.
	fmt.Println("no. of go routines : ", runtime.NumGoroutine())
	fmt.Println("main execution ended")
}

// Terdapat 2 function bernama firstProcess dan secondProcess.
// Kedua function tersebut digunakan untuk menampilkan angka dari 1
// hingga bilangan yang ditentukan dari parameter
// yang diterima dengan melakukan looping.

// function firstProcess dijadikan sebagai
// sebuah Goroutine karena dipanggil dengan
//  menggunakan keyword go.

func firstProcess(index int) {
	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	// function firstProcess tidak menampilkan hasilnya.
	//  Ini terjadi karena setiap Goroutine bekerja secara asynchronous
	// dan satu Goroutine tidak akan saling tunggu menunggu dengan Goroutine lainnya.
	fmt.Println("first process func ended")
}

// Kemudian jika kita perhatikan kembali pada hasilnya,
//  terdapat 2 jumlah Goroutine yang sedang berjalan padahal
//  kita hanya menjalankan satu function yang
// dijadikan sebagai sebuah Goroutine.

// Ini terjadi karena faktanya,
// function main juga merupakan sebuah Goroutine sehingga
// function main tidak akan menunggu Goroutine
//  lainnya selesai tereksekusi.

// Inilah yang menjadi penyebab function firstProcess
//  tidak menampilkan hasilnya walaupun
// sebetulnya function tersebut telah tereksekusi.

func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process func ended")
}
