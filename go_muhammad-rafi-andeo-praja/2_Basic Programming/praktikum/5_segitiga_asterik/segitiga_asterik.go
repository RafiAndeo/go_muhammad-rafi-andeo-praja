package main

import "fmt"

func main() {
	var n, k int
	fmt.Print("Masukkan nilai kolom: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= (n - i); j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println(" ")
	}
}
