package controllers

// mengimport package net/http untuk penentuan status code
// dan package fmt untuk memformat data string
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

//  function dengan nama CreateCar yang menerima
// satu parameter dengan tipe data *gin.Context

// Function CreateCar akan menjadi endpoint
// untuk proses create data.

// Function ini perlu menerima sebuah parameter
// dengan tipe data *gin.Context yang merupakan
// sebuah tipe data yang telah disediakan
// oleh framework Gin.

// *gin.Context mempunyai berbagai macam method yang dapat
// kita gunakan untuk mendapat request body dari client,
// mengirim response dan lain-lain.

func CreateCar(ctx *gin.Context) {
	var newCar Car

	// houldBindJSON sebuah method dari tipe data *gin.Context
	// yang digunakan untuk mem-binding data JSON
	// yang kirimkan oleh client sebagai request body
	// kepada server.

	// Method ShouldBindJSON menerima sebuah parameter
	// yang dimana kita perlu meletakkan pointer dari variable
	// yang akan menampung hasil dari data binding tersebu

	// Method ShouldBindJSON akan mereturn sebuah error
	// jika memang terjadi sebuah error,
	// maka dari itu kita perlu memvalidasi
	// terlebih dahulu jika terjadi sebuah error.
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		// AbortWithError untuk melempar error
		// jika memang ada error yang terjadi.
		// Method AbortWithError menerima 2 parameter.
		// Parameter pertama adalah status error nya,
		// kemudian parameter kedua adalah data error nya.
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// kita mencoba untuk mengenarate id untuk tiap data mobil baru
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	// jika tidak terjadi error ketika proses data binding
	// maka data request body yang telah ditampung
	// oleh variable newCar, akan kita tambahkan
	// pada variable global CarDatas.
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

// membuat proses untuk mengupdate data mobil
func UpdateCar(ctx *gin.Context) {
	// method ctx.Param merupakan sebuah method
	// yang digunakan untuk mendapat request parameter
	// yang dikirimkan oleh client.

	// Method ctx.Param menerima satu parameter
	// dimana kita perlu meletekkan nama parameter
	// yang kita buat pada router nya nanti.

	carID := ctx.Param("carID")

	condition := false
	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// sebelum kita betul-betul mengupdate data mobilnya,
	// kita melooping data-data mobil yang ditampung
	// oleh variable CarDatas
	// una untuk mencari terlebih dahulu data mobil
	// yang ingin di update oleh client nantinya.
	for i, car := range CarDatas {
		// // mencari data mobilnya berdasarkan
		// id data mobil yang dikirmkan client
		// melalui request parameter.

		if carID == car.CarID {
			// ka data mobilnya ada, maka kita langsung
			// mengupdate data mobilnya dan
			// menghentikan proses loopingnya.
			condition = true
			CarDatas[i] = updatedCar
			CarDatas[i].CarID = carID
			break
		}
	}

	// Jika data mobilnya tidak dapat,
	// maka kita langsung melempar error
	// dengan status Data Not Found(404)
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been suuccessfully updated", carID),
	})
}

// ka data mobilnya ada, maka kita langsung mengupdate
// data mobilnya dan menghentikan proses loopingnya.
func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData Car

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
	})
}
