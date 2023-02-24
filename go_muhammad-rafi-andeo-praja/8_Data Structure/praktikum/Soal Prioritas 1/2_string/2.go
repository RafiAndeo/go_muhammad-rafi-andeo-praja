package main

import "fmt"

func Mapping(slice []string) map[string]int {
	var size = len(slice)
	var n int
	var result = make(map[string]int)
	for i := 0; i < size; i++ {
		n = 0
		for j := 0; j < size; j++ {
			if slice[i] == slice[j] {
				n++
			}
		}
		result[slice[i]] = n
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         // map[]
}
