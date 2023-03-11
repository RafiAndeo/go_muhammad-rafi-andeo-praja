package main

import (
	"fmt"
)

func primeX(number int) int {

	if number == 1 {
		return 2
	}

	var prime int = 2
	var count int = 1

	for count < number {
		prime++
		if isPrime(prime) {
			count++
		}
	}

	return prime

}

func isPrime(number int) bool {

	if number == 2 {
		return true
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
