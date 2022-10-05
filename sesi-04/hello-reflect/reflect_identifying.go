package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23 //default number assign 23
	var reflectvalue = reflect.ValueOf(number)

	fmt.Printf("reflect value : %+v\n", number)
	fmt.Printf("tipe variabel : %+v\n", reflectvalue.Type())
	//cek interface
	fmt.Printf("nilai variabel : %+v\n", reflectvalue.Interface())

	if reflectvalue.Kind() == reflect.Int {
		fmt.Printf("nilai variabel : %d\n", reflectvalue.Int())
	}
}
