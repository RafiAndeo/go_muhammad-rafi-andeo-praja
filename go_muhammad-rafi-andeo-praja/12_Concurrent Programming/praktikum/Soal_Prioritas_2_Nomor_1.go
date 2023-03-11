package main

import (
	"fmt"
)

func characterFrequency(str string) {
	var freq [256]int
	for _, v := range str {
		freq[v]++
	}

	for i, v := range freq {
		if v != 0 {
			fmt.Printf("%c = %d times \n", i, v)
		}
	}
}

func main() {
	characterFrequency("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
}
