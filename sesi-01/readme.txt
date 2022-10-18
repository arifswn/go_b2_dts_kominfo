add path golang mac
---
cd ~
echo "export GOPATH=/Users/arif/Documents/go" >> .bash_profile
echo "export GOPATH=/Users/arif/Documents/go" >> .zshrc
cat .bash_profile
cat .zshrc
echo $GOPATH

cek versi golang
----
go version
echo $GOPATH

create folder project
---
mkdir hello
cd hello

first project
---
go mod init hello > create module
touch main.go > create file golang
go run main.go > run golang
go build -o main.exe >> membutuhkan module

instalasi setup
---
tools vscode (visual studio code)
    > plugins/extensi > 
        - code runner
        - golang / go
        - Prettier - Code formatter
tools golang > https://go.dev/dl/

type data golang
---
nil bukan merupakan tipe data, melainkan sebuah nilai. 
Variabel yang isi nilainya nil berarti memiliki nilai kosong.
Semua tipe data yang sudah dibahas sebelumnya memiliki zero value (nilai default tipe data).
Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, tetap akan ada nilai default-nya.

- Zero value dari string adalah "" (string kosong).
- Zero value dari bool adalah false.
- Zero value dari tipe numerik non-desimal adalah 0.
- Zero value dari tipe numerik desimal adalah 0.0.
- Zero value berbeda dengan nil. Nil adalah nilai kosong, benar-benar kosong.

Nil tidak bisa digunakan pada tipe datayang sudah dibahas sebelumnya.
Ada beberapa tipe data yang bisa di-set nilainya dengan nil, diantaranya:●pointer●tipe data function
- slice
- map
- channel
- interface kosong atau interface{}

link referensi
---
https://codeshare.io/golangbatch2-001