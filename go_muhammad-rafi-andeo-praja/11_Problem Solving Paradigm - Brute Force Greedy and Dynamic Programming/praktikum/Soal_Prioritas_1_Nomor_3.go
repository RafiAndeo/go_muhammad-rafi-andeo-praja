package main

import "fmt"

func fibo(n int) int {

	var result = map[int]int{}
	if n < 2 {
		return n
	} else {
		result[n] = fibo(n-2) + fibo(n-1)
		return result[n]
	}

}

func main() {

	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 34
	fmt.Println(fibo(10)) // 55

}
