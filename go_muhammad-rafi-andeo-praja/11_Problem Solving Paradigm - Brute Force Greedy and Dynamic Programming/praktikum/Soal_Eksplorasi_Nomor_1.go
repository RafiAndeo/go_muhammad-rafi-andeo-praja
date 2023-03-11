package main

import "fmt"

func romanNumbers(number int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	decimal := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumber := ""
	for i := 0; i < len(roman); i++ {
		for number >= decimal[i] {
			number -= decimal[i]
			romanNumber += roman[i]
		}
	}
	return romanNumber
}

func main() {
	fmt.Println(romanNumbers(4))    // IV
	fmt.Println(romanNumbers(9))    // IX
	fmt.Println(romanNumbers(23))   // XXIII
	fmt.Println(romanNumbers(2021)) // MMXXI
	fmt.Println(romanNumbers(1646)) // MDCXLVI
}
