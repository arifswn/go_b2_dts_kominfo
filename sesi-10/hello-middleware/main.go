package main

import (
	"fmt"
	"log"
	"net/http"
)

// Untuk membuat sebuah middleware, maka kita harus suatu function yang
// memuaskan interface http.Handler yang dimana interface ini
// memiliki satu buah method dengan berupa
// ServeHTTP(http.ResponseWriter, *http.Request).

// Function middleware1 merupakan sebuah middleware.
// Setiap middleware pada bahasa Go akan menerima satu parameter
// dengan tipe data http.Handler dan harus mereturn tipe data http.Handler.

// function middleware1 mereturn sebuah anonymous function dengan nama http.HandlerFunc
// yang dimana function ini menerima satu argument berupa tipe data function
// yang mempunyai argument yang sama seperti method ServeHTTP

func main() {

	// kita mendeklarasikan sebuah variable bernama mux yang akan menjadi multiplexer nya.
	// Kemudian kita membuat variable bernama endpoint yang menampung satu-satunya
	//  endpoint pada aplikasi kita saat ini yaitu function greet. 
	
	// function greet sebagai argument dari function http.HandlerFunc 
	// agar function greet dapat menjadi sebuah function dengan tipe data http.Handler.
	
	// membuat routingnya dengan menggunakan method Handle dari pointer 
	// struct ServeMux.Method Handle menerima 2 parameter berupa string 
	// dari route path nya, dan handler nya yang bertipe http.Handler.
	
	// menggunakan 2 middleware sebelum kita dapat mencapai 
	// function greet yang merupakan endpointnya.

	mux := http.NewServeMux()
	endpoint := http.HandlerFunc(greet)
	mux.Handle("/", middleware1(middleware2(endpoint)))

	fmt.Println("listening to port 8080")
	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World !!!"))
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware pertama")
		// diingat bahwa function http.Handlerf unc secara default mengimplementasikan
		// method ServeHTTP dan setiap function yang mempunyai skema parameter yang sama
		// dengan method ServeHTTP  dapat di ubah tipe datanya menjadi interface http.Handler
		// dengan cara menjadikannya parameter dari function http.HandlerFunc.

		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware kedua")
		next.ServeHTTP(w, r)
	})
}
