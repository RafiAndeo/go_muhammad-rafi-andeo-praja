package main

import "fmt"

func pascalTriangle(n int) [][]int {

	var arr [][]int = make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
			}
		}
	}
	return arr

}

func main() {
	fmt.Println(pascalTriangle(5))
}
