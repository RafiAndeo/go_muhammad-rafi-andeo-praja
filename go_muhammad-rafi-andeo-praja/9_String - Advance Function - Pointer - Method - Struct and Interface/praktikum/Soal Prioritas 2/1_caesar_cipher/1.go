package main

import "fmt"

func caesar(offset int, input string) string {

	var convert string

	for _, char := range input {

		if char >= 'A' && char <= 'Z' {
			convert = convert + string('A'+(char-'A'+int32(offset))%26)
		} else if char >= 'a' && char <= 'z' {
			convert = convert + string('a'+(char-'a'+int32(offset))%26)
		} else {
			convert = convert + string(char)
		}

	}

	return convert

}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
