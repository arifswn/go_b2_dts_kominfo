package routers

import (
	"hello-gin/controllers"

	"github.com/gin-gonic/gin"
)

// Function StartServer ini mengembalikan suatu data
// dengan tipe data struct *gin.Engine
// yang berasal dari Gin, dan kita gunakan
// untuk menjalankan server, sebagai multiplexer
// dari routing dan lain-lain
// multiplexer mengubah input (banyak input) menjadi satu output
func StartServer() *gin.Engine {
	// variable router kita gunakan sebagai penampung
	// untuk engine dari router yang kita dapatkan
	// dari pemanggilan function gin.Default.
	router := gin.Default()

	// membuat endpoint untuk membuat data mobil baru,
	// maka kita perlu membuat routing yang akan menghubungkan
	//  client kepada endpoint tersebut
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
