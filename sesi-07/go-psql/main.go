package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db_go"
)

var (
	db  *sql.DB
	err error
)

var option int

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// function Open tidak akan membangun koneksi
	// ke database, melainkan hanya berfungsi untuk
	// memvalidasi arugmen-argumen yang diberikan.
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// kita memanggil method Ping. Kita dapat menggunakan
	// method Ping karena setelah penggunaan function Open
	// berhasil tereksekusi, maka function Open akan
	// mengembalikan nilai berupa pointer dari struct
	// sql.DB yang dimana struct tersebut memiliki
	// method Ping.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	fmt.Println()

	Menu()
}

func Menu() {
	stdin := bufio.NewReader(os.Stdin)

	fmt.Println(".:: Menu ::.")
	fmt.Println()
	fmt.Println("1. Create Data")
	fmt.Println("2. Get Data")
	fmt.Println("3. Update Data")
	fmt.Println("4. Delete Data")
	fmt.Println("5. Exit")
	fmt.Println()
	fmt.Print("Choose an option from the menu : ")
	for {
		_, err := fmt.Fscan(stdin, &option)
		if err == nil {
			break
		}
		stdin.ReadString('\n')
		fmt.Print("Wrong option input. Please, try again : ")
	}

	switch option {
	case 1:
		CreateEmployee()
	case 2:
		GetEmployes()
	case 3:
		UpdateEmployee()
	case 4:
		DeleteEmployee()
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Data tidak ada !")

	}

}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
		insert into employees (full_name, email, age, division)
		values ($1, $2, $3, $4)
		Returning *
	`

	err = db.QueryRow(sqlStatement, "Airell Jordan", "arif@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New employee data : %+v\n", employee)
}

func GetEmployes() {
	var results = []Employee{}

	// method Query biasa digunakan untuk mendapatkan banyak
	//  data dari suatu tablepada database dikarenakan
	// method ini dapat mengembalikan satu atau lebih rows
	// dari suatutable pada database.
	sqlStatement := `select * from employees`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	// Dan kita juga harus menutup rows tersebut dengan
	// method Close yang kita panggil dengan keyword defer
	defer rows.Close()

	// perulangan/looping sebanyak data yang kita dapatkan
	//  dengan acuan rows.Next.
	// Method rows.Next akan menghasilkan nilai true
	// selama data nya masih ada, namun jika sudah tidak ada
	//  maka dia akan me-return false dan proses looping
	// akan terhenti

	for rows.Next() {
		var employee = Employee{}
		//  proses scaning untuk satu row,
		// maka kita akan memasukkan data dari row tersebut
		//  kedalam variable results
		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas: ", results)
}

func UpdateEmployee() {
	sqlStatement := `
		update employees set full_name=$2, email=$3, division=$4, age=$5
		where id=$1;
	`
	// kita menggunakanmethod Exec untuk mengupdate data
	// Sebetulnya kita juga bisa mengupdate sebuah data
	// dengan menggunakan method QueryRow, namun dianjurkan
	// untuk menggunakan method Exec untuk melakukan
	// proses insert, update, dan delete.

	// Dengan menggunakan method Exec,
	// kita tidak dapat mendapatkan data yang
	// baru saja diupdate maupun data yang baru kita buat,
	res, err := db.Exec(sqlStatement, 3, "Arif Setiawan", "arif@gmail.com", "Staff", 21)
	if err != nil {
		panic(err)
	}
	// kita dapat menggunakan method RowsAffected untuk
	// mengetahui berapa jumlah row atau data yang
	// baru diupdate.

	// Method RowsAffected didapatkan dari interface
	// sql.Result yang merupakan nilai kembalian dari
	// method Exec.

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated data amount: ", count)
}

func DeleteEmployee() {
	sqlStatement := `
		delete from employees where id = $1;
	`

	res, err := db.Exec(sqlStatement, 3)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount: ", count)
}
