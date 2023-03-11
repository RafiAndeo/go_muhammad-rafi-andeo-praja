package main

import (
	"fmt"
	"time"
)

func kelipatan(number chan int) {
	data := <-number
	if data%3 == 0 {
		fmt.Println("Bilangan", data, "adalah bilangan kelipatan 3")
	} else {
		fmt.Println("Bilangan", data, "bukan bilangan kelipatan 3")
	}
}

func main() {
	x := make(chan int, 2)
	y := make(chan int, 2)
	go kelipatan(x)
	x <- 3
	go kelipatan(y)
	y <- 5
	time.Sleep(3 * time.Second)
}
