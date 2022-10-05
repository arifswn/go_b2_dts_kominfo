package main

import (
	"fmt"
	"sync"
)

func main() {
	// variable fruits merupakan sebuah variable yang
	// menampung 4 data buah-buahan.

	fruits := []string{"apple", "mango", "durian", "rambutan"}

	// terdapat sebuah variable bernama wg dengan
	// tipe data sync.Waitgroup yang merupakan
	// sebuah struct daripackage sync.

	var wg sync.WaitGroup

	// Kemudian data buah-buahan pada variable fruits
	// hendak kita looping karena kita ingin melakukan
	// print terhadap data buah-buahannya.
	for index, fruit := range fruits {
		// Method Add digunakan untuk menambah counter dari waitgroup.Counter
		// dari waitgroup ini berguna untuk memberitahu waitgroup
		// tentang jumlah go routine yang harus ditunggu.
		//  Karena kita melooping sebanyak 4 kali,
		//  berarti terdapat 4 go routine yang harus
		// ditunggu sebelum function main menghentikan eksekusinya
		wg.Add(1)
		// Didalam looping tersebut,
		//  kita memanggil function printFruit
		// yang menerima 3 parameter,
		// dan function printFruit kita jadikan sebagai goroutine.
		go printFruit(index, fruit, &wg)
	}

	wg.Wait()
}

// dan function printFruit kita jadikan sebagai goroutine.
func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}
