package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama : ", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	// Jika eksekusi method diikuti pengisian parameter,
	// maka parameternya harus ditulis dalam bentuk
	// array[]reflect.Value berurutan sesuai urutan
	// deklarasi parameter-nya.Dan nilai yang dimasukkan
	//  ke array tersebut harus dalam bentuk reflect.Value
	//  (gunakan reflect.ValueOf() untukpengambilannya).
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})
	fmt.Println("nama : ", s1.Name)
}
