Middleware merupakan sebuah fungsi yang akan terkesekusi 
sesudah maupun sebelum mencapai sebuah endpoint. 

Biasa middleware digunakan untuk logging atau untuk mengamankan 
sebuah endpoint seperti contohnya proses autentikasi dan autorisasi.

Untuk membuat middleware pada bahasa Go, 
kita akan menggunakan package net/http dengan 
menggunakan multiplexer nya agar kita dapat melakukan kustomisasi.

go get github.com/asaskevich/govalidator
go get github.com/dgrijalva/jwt-go
go get github.com/gin-gonic/gin
go get golang.org/x/crypto
go get gorm.io/driver/postgres
go get gorm.io/gorm

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required-Title of your product is required"`
	Description string `json:"description" form:"description" valid:"required-Description of your product is required"`
	UserID      uint
	User        *User
}

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required-Your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required-Your email is required,email-Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required-Your password is required,minstringlength(6)-Password has to have a minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

plugins : Live Server

reference : 
https://getbootstrap.com/docs/4.3/getting-started/introduction/
https://getbootstrap.com/docs/4.1/content/tables/
https://jsonplaceholder.typicode.com/todos/2