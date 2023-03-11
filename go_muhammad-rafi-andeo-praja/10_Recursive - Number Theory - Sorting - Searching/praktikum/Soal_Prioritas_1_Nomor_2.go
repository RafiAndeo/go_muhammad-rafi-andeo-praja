package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {

	if len(items) == 0 {
		return []pair{}
	}

	last := items[len(items)-1]
	rest := items[:len(items)-1]

	pairs := MostAppearItem(rest)
	found := false

	for i := range pairs {
		if pairs[i].name == last {
			pairs[i].count++
			found = true
			break
		}
	}

	if !found {
		pairs = append(pairs, pair{last, 1})
	}

	return pairs

}

func main() {

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1, ruby->2, js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1, D->2, B->3, A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1, basketball->1, tenis->1

}
