package main

import "fmt"

func arrayBiner(n int) []int {

	var arr []int = make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = getBinary(i)
	}
	return arr

}

func getBinary(num int) int {

	var binary = num % 2
	var half = num / 2
	if num == 0 {
		return 0
	}
	return binary + 10*getBinary(half)

}

func main() {

	fmt.Println(arrayBiner(2))
	fmt.Println(arrayBiner(3))

}
