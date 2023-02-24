package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var array []string
	var n = len(arrayA)
	var m = len(arrayB)
	for i := 0; i < n; i++ {
		array = append(array, arrayA[i])
	}
	for j := 0; j < m; j++ {
		array = append(array, arrayB[j])
	}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				array = append(array[:j], array[j+1:]...)
				j--
			}
		}
	}
	return array
}

func main() {

	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "gesse"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "gesse"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []

}
