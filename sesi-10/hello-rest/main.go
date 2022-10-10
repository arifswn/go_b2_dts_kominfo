package main

import (
	"hello-rest/database"
	"hello-rest/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()

	r.Run(":8080")
}
