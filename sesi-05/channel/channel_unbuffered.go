package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutines starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutines after sending data into the channel")
	}(c1)

	fmt.Println("main gorotine sleeps for 2 second")
	time.Sleep(time.Second * 2)

	fmt.Println("main gorotine starts receiving data")
	d := <-c1
	fmt.Println("main gorotine received data : ", d)

	close(c1)
	time.Sleep(time.Second)
}
