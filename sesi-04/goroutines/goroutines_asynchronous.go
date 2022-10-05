package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("main execution started")

	go firstProcess(8)
	secondProcess(8)
	fmt.Println("no. of go routines : ", runtime.NumGoroutine())

	time.Sleep(time.Second * 2) //ekseuksi timeout
	fmt.Println("main execution ended")
}


func firstProcess(index int) {
	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("first process func ended")
}

func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process func ended")
}


