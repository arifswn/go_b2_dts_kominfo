package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

// kita memakai function HandleFunc yang berasal dari package http.
// Function HandleFunc digunakan untuk keperluan routing,
// yang dimana function tersebut menerima 2 parameter

// Parameter pertama digunakan untuk
// mendefinisikan path outingnya,
// sedangkan parameter kedua menerima sebuah function
// dengan 2 parameter yaitu http.ResponseWriter
// dan pointer http.Request

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}

// http.ResponseWriter adalah sebuah interface dengan
// berbagai method yang digunakan untuk
//  mengirim response balik
// kepada client yang mengirimkan request.

// http.Request adalah sebuah struct yang digunakan
// untuk mendapat berbagai macam data request
// seperti form value,headers, url parameter dan lain-lain

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprintf(w, msg)
}
