package main

import "fmt"

func main() {
	var fruits1 = []string{"apple", "banana", "mango", "durian", "pinneaple"}

	// mengkombinasikan fungsi append dengan slicing.
	// Jika kitaperhatikan,
	// variable fruits1 ingin me-replace index ketiga
	// hingg aseterusnya dengan hanya buah “rambutan” saja
	fruits1 = append(fruits1[:3], "rambutan")
	fmt.Printf("fruits 1 => %#v\n", fruits1)

}
