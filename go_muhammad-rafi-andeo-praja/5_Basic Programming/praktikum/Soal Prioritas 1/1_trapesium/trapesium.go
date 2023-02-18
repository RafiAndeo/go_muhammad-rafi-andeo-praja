package main

import "fmt"

func main() {
	var sisi_a, sisi_b, tinggi, luas float64
	fmt.Print("Masukkan panjang sisi a: ")
	fmt.Scan(&sisi_a)
	fmt.Print("Masukkan panjang sisi b: ")
	fmt.Scan(&sisi_b)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)
	luas = (sisi_a + sisi_b) * tinggi / 2
	fmt.Println("Luas trapesium adalah", luas)
}
