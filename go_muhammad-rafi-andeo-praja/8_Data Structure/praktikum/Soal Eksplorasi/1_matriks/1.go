package main

import "fmt"

func selisihAbsolutDiagonal(arr [][]int) int {
	var n = len(arr)
	var diagonal1, diagonal2 int
	for i := 0; i < n; i++ {
		diagonal1 = diagonal1 + arr[i][i]
		diagonal2 = diagonal2 + arr[i][n-i-1]
	}
	if diagonal1 > diagonal2 {
		return diagonal1 - diagonal2
	}
	return diagonal2 - diagonal1
}

func main() {
	fmt.Println(selisihAbsolutDiagonal([][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}})) // 2
}
