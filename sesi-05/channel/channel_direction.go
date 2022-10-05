package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"airell", "mailo", "indah"}

	for _, v := range students {
		go introduce(v, c)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)
}

// karena parameter channel pada function print
// hanya digunakan untuk menerima data,
// maka penulisan parameternya kita ubah menjadi c <- chan string
func print(c <-chan string) {
	fmt.Println("", <-c)
}

// karena parameter channel pada function introducekarena parameter channel pada  function introduce hanya digunakan untuk mengirim data, maka penulisan parameternya kita ubah menjadi c chan <- string
//  hanya digunakan untuk mengirim data,
// maka penulisan parameternya
// kita ubah menjadi c chan <- string
func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("hei, my name is %s", student)
	c <- result
}
