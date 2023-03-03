package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	var result []pair
	var temp []string
	var found bool
	for i := 0; i < len(items); i++ {
		found = false
		for j := 0; j < len(temp); j++ {
			if items[i] == temp[j] {
				found = true
				break
			}
		}
		if !found {
			temp = append(temp, items[i])
		}
	}
	for i := 0; i < len(temp); i++ {
		var count int
		for j := 0; j < len(items); j++ {
			if temp[i] == items[j] {
				count++
			}
		}
		result = append(result, pair{temp[i], count})
	}
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].count < result[j].count {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}

func main() {

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1, ruby->2, js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1, D->2, B->3, A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1, basketball->1, tenis->1

}
