package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println(x, "adalah bilangan genap")
	} else {
		fmt.Println(x, "adalah bilangan ganjil")
	}
}
