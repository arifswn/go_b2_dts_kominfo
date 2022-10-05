package main

import (
	"fmt"
)

// Anonymous struct adalah sebuah struct yang
// tidak dideklerasikan di awal sebagai sebuah
// tipe data struct baru,

// melainkan langsung dideklerasikan bersamaan dengan pembuatan variable.
// Terdapat sebuah struct bernama Person yang mempunyai 2 field.
type Person struct {
	name string
	age  int
}

func main() {

	// anonymous struct tanpa pengisian field
	//  Kemudianterdapat sebuah variable bernama employee1
	// yang mengandung tipe data struct yang merupakan
	// sebuah Anonymous struct tanpa pengisian nilai field.

	// Salah satu aturan yang perlu diingat dalam
	// pembuatan anonymous struct adalah,
	// deklarasi harus diikuti dengan inisialisasi.

	// Bisa dilihat pada variable employee1 setelah deklarasi struktur struct,
	// terdapat kurung kurawal untuk inisialisasi.
	// Meskipun nilai tidakdiisikan di awal,
	// kurung kurawal tetap harus ditulis.

	var employee1 = struct {
		person   Person
		division string
	}{}

	employee1.person = Person{name: "Airel", age: 23}
	employee1.division = "Curriculum Developer"

	fmt.Printf("hasil nya : %+v\n", employee1)

	// anonymous struct dengan pengisian field
	// sebuah variable bernama employee2 yang juga
	// mengandung tipe data struct berupa Anonymous struct
	// dengan pengisian nilai field.

	var employee2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Ananda", age: 23},
		division: "Finance",
	}

	fmt.Printf("employee 1 : %+v\n", employee1)
	fmt.Printf("employee 2 : %+v\n", employee2)

}
