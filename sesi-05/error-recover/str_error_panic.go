package main

import (
	"errors"
	"fmt"
)

func main() {
	var password string
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())

		// Panic digunakan untuk menampilkan stack trace error
		// sekaligus menghentikan flow goroutine
		// (karena main() juga merupakan goroutine,
		// makabehaviour yang sama juga berlaku).

		// Setelah ada panic, proses akan terhenti,
		// apapun setelah tidakdi-eksekusi kecuali
		// proses yang sudah di-defer sebelumnya
		// (akan muncul sebelum panic error).

		// Panic menampilkan pesan error di console,
		// sama seperti fmt.Println().
		// Informasi error yang ditampilkan adalah
		// stack trace error, jadi sangat mendetail dan heboh.
		panic(err.Error())
	} else {
		fmt.Println("status password anda : ", valid)
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)
	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "valid password", nil
}
