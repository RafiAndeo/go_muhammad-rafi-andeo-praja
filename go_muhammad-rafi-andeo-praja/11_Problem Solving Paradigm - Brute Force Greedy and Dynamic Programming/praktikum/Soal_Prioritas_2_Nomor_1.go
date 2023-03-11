package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {

	var firstJump, secondJump int
	var min int
	cost := make([]int, len(jumps))
	cost[0] = 0

	for i := 1; i < len(jumps); i++ {
		firstJump = cost[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))

		if i > 1 {
			secondJump = cost[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		} else {
			secondJump = firstJump
		}
		min = int(math.Min(float64(firstJump), float64(secondJump)))
		cost[i] = min
		secondJump = firstJump
		firstJump = min
	}

	return min

}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
