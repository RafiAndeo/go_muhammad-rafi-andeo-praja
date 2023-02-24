package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var result []int
	var size = len(angka)
	var n int
	for i := 0; i < size; i++ {
		n = 0
		for j := 0; j < size; j++ {
			if angka[i] == angka[j] {
				n++
			}
		}
		if n == 1 {
			angkaInt, _ := strconv.Atoi(string(angka[i]))
			result = append(result, angkaInt)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
