package main

import "fmt"

func main() {
	var kata, hasil string
	fmt.Println("Apakah Palindrome?")
	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)
	fmt.Println("Captured: ", kata)
	var panjang = len(kata)
	for i := panjang - 1; i >= 0; i-- {
		hasil += string(kata[i])
	}
	if hasil == kata {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}
