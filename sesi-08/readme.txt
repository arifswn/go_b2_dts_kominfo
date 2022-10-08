summary sesi 08
---

step swago :
mkdir swaggo
cd swaggo
go mod init swaggo

library :
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/http-swagger
go get -u github.com/alecthomas/template
go get -u github.com/gorilla/mux

generate swagger/docs : 
    swag init -g main.go
    go run main.go

global :
go install github.com/swaggo/swag/cmd/swag

add path :
export PATH=$(go env GOPATH)/bin:$PATH

custom :
go install github.com/swaggo/swag/cmd/swag@v1.8.6

reference :
https://www.soberkoder.com/go-rest-api-gorilla-mux/
https://www.soberkoder.com/swagger-go-api-swaggo/
https://github.com/soberkoder/swaggo-orders-api/
https://github.com/soberkoder/go-orders-api/
https://github.com/jonnylangefeld/go-api/

how to access swagger :
http://localhost:8080/swagger/index.html


practice :
go get -u "github.com/gin-gonic/gin"
go get -u "github.com/jinzhu/gorm"
go get -u "github.com/jinzhu/gorm/dialects/mysql"
go get -u "github.com/go-sql-driver/mysql"
