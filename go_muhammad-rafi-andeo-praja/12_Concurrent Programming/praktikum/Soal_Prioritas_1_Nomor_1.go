package main

import (
	"fmt"
	"time"
)

func kelipatan(number int) {

	if number > 0 {
		if number%2 == 0 {
			fmt.Println("Bilangan", number, "adalah bilangan kelipatan 2")
		} else if number%3 == 0 {
			fmt.Println("Bilangan", number, "adalah bilangan kelipatan 3")
		} else if number%5 == 0 {
			fmt.Println("Bilangan", number, "adalah bilangan kelipatan 5")
		} else if number%7 == 0 {
			fmt.Println("Bilangan", number, "adalah bilangan kelipatan 7")
		} else {
			fmt.Println("Bilangan", number, "bukan bilangan kelipatan 2, 3, atau 5")
		}
	} else {
		fmt.Println("Bilangan harus lebih besar dari 0")
	}

}

func main() {

	go kelipatan(2)
	go kelipatan(3)
	go kelipatan(5)
	go kelipatan(7)
	go kelipatan(11)
	go kelipatan(-1)

	time.Sleep(3 * time.Second)

}
