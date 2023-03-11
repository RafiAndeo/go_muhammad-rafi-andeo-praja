package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {

	for _, card := range cards {
		for _, value := range card {
			for _, deckValue := range deck {
				if value == deckValue {
					return card
				}
			}
		}
	}

	if len(cards) == 0 {
		return "tutup kartu"
	}

	return playingDomino(cards[1:], deck)

}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	// [3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	// [6, 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
	// "tutup kartu"
}
