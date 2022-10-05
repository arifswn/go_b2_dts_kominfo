package database

import (
	"fmt"
	// "go-gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbName   = "db_go"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		// Agar tidak terjadi error, maka jumlah karakter Name dari Product
		// harus terdiri dari 4 karaketer atau lebih.
		log.Fatal("error connecting to database : ", err)
	}

	fmt.Println("Successfully connected to database")
	fmt.Println()
	// db.Debug().AutoMigrate(models.User{}, models.Product{})
}

// Untuk melakukan CRUD data melalui Gorm, maka kita perlu mendapatkan
// referensi database yang telah ditampung pada suatu variable
func GetDB() *gorm.DB {
	return db
}
