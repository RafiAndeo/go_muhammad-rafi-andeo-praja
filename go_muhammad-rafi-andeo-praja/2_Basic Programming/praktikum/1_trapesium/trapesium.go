package main

import "fmt"

func main() {
	var a, b, tinggi float64
	fmt.Print("Masukkan panjang sisi a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan panjang sisi b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)
	var luas float64 = (a + b) * tinggi / 2
	fmt.Println("Luas trapesium adalah", luas)
}
