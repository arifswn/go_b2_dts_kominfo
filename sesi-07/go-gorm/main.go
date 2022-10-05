package main

import (
	"errors"
	"fmt"
	"go-gorm/database"
	"go-gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("johndoe@gmail.com")
	// getUserById(1)
	// updateUserById(1, "arif@gmail.com")
	// createProduct(2, "XXX", "Saos")
	// getUsersWithProducts()
	// deleteProductById(1)
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data : ", err)
		return
	}

	fmt.Println("New user data : ", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	// Method First dapat menerima 3 parameter, parameter pertamanya
	// adalah pointer terhadap data yang ingin dicari
	// dengan memberikan tipe data struct User
	//  sebagai argumen pertama dari method Firs

	//  parameter keduanya adalah condition dari query nya,
	//  dan yang terakhiir adalah data dari conditionny

	// Method First akan mengembalikan error berupa
	//  ErrRecordNotFound jika tidak ada data yang ditemuka
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user", err)
	}
	fmt.Printf("User data: %+v \n", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	//  method Model dapat langsung discan sehingga kita dapat langsung
	//  mengetahui hasilnya
	//  menggunakan method Where sehingga method Model
	// dapat dichaining dengan method Where
	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	// err := db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"email": email}).Error

	if err != nil {
		fmt.Println("Error updateing user data", err)
		return
	}

	fmt.Printf("Update user's email: %+v \n", user.Email)
	// fmt.Printf("Update user's email: %+v \n", email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data: ", err.Error())
		return
	}

	fmt.Println("New Product Data: ", Product)
}

// Eager Loading
func getUsersWithProducts() {

	// Sekarang data user dengan id 1 sudah memiliki data product,
	// maka kita dapat melakukan join statement.
	// Untuk melakukan join statement,
	// kita dapat menggunakan eager loading dari Gorm.

	db := database.GetDB()

	users := models.User{}
	// Caranya adalah dengan menggunakan method Preload dan kita perlu memberikan
	// nama table untuk parameter method Preload.
	// Walaupun nama tablenya adalah products
	// jika kita melihat pada database, tapi tetap huruf awal
	//  dari nama table untuk parameter Preload harus menggunakan uppercase.
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products: ", err.Error())
		return
	}

	fmt.Printf("User datas with products %+v\n", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product: ", err.Error())
		return
	}
	fmt.Printf("Product with id %d has been successfully deleted", id)
}
