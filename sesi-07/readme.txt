
step 1 :
---
mkdir go_psql
cd go_psql
go mod init go-psql
go get -u github.com/lib/pq

create database : 
---
create db_go

create table :
---
create table employees(
	id serial primary key,
	full_name varchar(50) not null,
	email text unique not null,
	age int not null,
	division varchar(20) not null
)

mysql : 
go get -u github.com/go-sql-driver/mysql