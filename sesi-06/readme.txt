Go telah menyediakan sebuah package bernama net/http
untuk berbagai macam keperluan dalam membuat 
sebuah aplikasi berbasis web seperti 
contohnya routing, templating, web server dan lain-lain. 

Dan pada materi kali ini kita akan
belajar bagaimana cara membuat web server 
beserta routingnya pada bahasa Go

####################################

mkdir hello-gin
    cd hello-gin
        go mod init hello-gin
        gin framework > go get -u github.com/gin-gonic/gin
mkdir controllers
    touch carCountroller.go
    cd ..
mkdir routers
    touch carRouter.go
    cd ..
touch main.go
go run main.go

####################################

tools :
    - postman
    - insomnia
    - testmace
    - http master

    - plugsin vscode 
        - rest client
        - post code

####################################

sample api :    
    https://jsonplaceholder.typicode.com
    https://reqres.in
    https://gorest.co.in

convert json :
    http://jsonviewer.stack.hu/
    https://jsoneditoronline.org
    https://jsonformatter.org

akses api golang to local :
    http://localhost:8080/
    http://localhost:8080/cars (create|delete|put)
    http://localhost:8080/cars/c1 (get params)
    http://localhost:8080/cars/all (get all)

####################################

request body code :
create / put(update)
{
    "brand": "Toyota",
    "model": "Avanza",
    "price": 1500
}

database server :
    - postgree
    - laragon
    - mariadb (xampp 7.4 / 5.6)

tools database :
    - DBeaver > all sql
    - HeidiSQL > all
    - Adminer / Adminer Custom