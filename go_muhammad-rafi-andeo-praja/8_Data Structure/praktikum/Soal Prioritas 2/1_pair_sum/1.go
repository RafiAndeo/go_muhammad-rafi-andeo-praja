package main

import "fmt"

func PairSum(arr []int, target int) []int {
	var result []int
	var n = len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
