package main

import (
	"fmt"
	"runtime"
	"time"
)

// Perlu diingat disini bahwa ketika kita menjalankan sebuah Goroutine,
//  maka Goroutine tersebut akan membutuhkan waktu yang sedikit
// lebih lama untuk memulai dibandingkan dengan function biasa.

func main() {
	fmt.Println("main execution started")

	go firstProcess(8)
	secondProcess(8)
	fmt.Println("no. of go routines : ", runtime.NumGoroutine())

	// Maka dari itu untuk sekarang,
	//  kita akan membutuhkan suatu function yang akan
	// menahan function main untuk langsung
	// menyelesaikan eksekusinya.

	// menggunakan function Sleep yang berasal dari package time.
	//  Lalu kita konstanta bernama Second dari package time
	// yang merepresentasikan bilangan detik.
	time.Sleep(time.Second * 2) //ekseuksi timeout

	// Kemudian kita kalikan dengan angka 2 agar
	//  func Sleep mampu menahan function main
	// selama 2 detik sebelum function main menyelesaikan eksekusinya.
	fmt.Println("main execution ended")
}

func firstProcess(index int) {
	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("first process func ended")
}

func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process func ended")
}

// nBisa kita lihat bahwa sekarang function firstProcesstelah menampilkan hasilnya.
