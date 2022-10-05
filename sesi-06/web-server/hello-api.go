package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// http.ResponseWriter adalah sebuah interface dengan
// berbagai method yang digunakan untuk
//  mengirim response balik
// kepada client yang mengirimkan request.
type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

// http.Request adalah sebuah struct yang digunakan untuk
// mendapat berbagai macam data request
// seperti form value,headers, url parameter dan lain-lain

// membuat data-data employee itu terlebih dahulu
// dengan menyimpannya didalam sebuah variable
// dengan tipe data slice yang berisikan
// tipe data structEmployee
var employees = []Employee{
	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
	{ID: 2, Name: "Nanda", Age: 25, Division: "Finance"},
	{ID: 3, Name: "Mailo", Age: 20, Division: "IT"},
}

var PORT = ":8080"

func main() {
	// index/home
	http.HandleFunc("/", getHome)
	// get employees
	http.HandleFunc("/employees", getEmployees)
	// get employees with html parse
	http.HandleFunc("/employees_html", getHtml)
	// create employee
	// meregistrasikan functioncreateEmployee kedalam routingan kita
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	msg := "Welcome to RESTFull API"
	fmt.Fprintf(w, msg)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	// menggunakan method Header dari
	// interfacehttp.ResponseWriter
	// yang kemudian di chaining dengan method Set.

	// Karena saat ini kita ingin mengirim response data
	// dalam bentuk JSON, maka kita dapat mengatur
	// Content-Type nya menjadiapplication/json
	// dalam method Set.

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	// http.StatusBadRequest merupakan sebuah konstanta
	// dari package http.StatusBadRequest
	// yang merepresentasikan status 400.

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// menerakan method POST yang dimana kita gunakan untuk
	// menambah data employee baru
	if r.Method == "POST" {

		// // untuk mendapatkan nilai inputform dari client dengan
		// menggunakan method FormValue
		// yang berasal dari struct http.Request
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		// mengkonversi input age dari tipedata string menjadi
		//  tipe data int karena seluruh nilai input
		// yang kita dapat dari client akan
		// memiliki tipe data string,

		// sedangkanfield age pada struct Employee
		// memiliki tipe data int
		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age : "+err.Error(), http.StatusBadRequest)
			return
		}

		// membuat variable newEmployee yan gakan
		// menampung nilai-nilai input dari client
		// yang kemudian akan kita append
		//  atau masukkan kedalam slice employees
		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method : ", http.StatusBadRequest)
}

func getHtml(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// menggunakan function template.ParseFiles yang berasal
		// dari package html/template yang digunakan untuk
		// mem-parsing file html kita

		// jika sudah di parsing oleh function template.ParseFiles
		// function tersebut akanm engembaikan
		// suatu tipe data struct template
		tpl, err := template.ParseFiles("hello.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusInternalServerError)
}
