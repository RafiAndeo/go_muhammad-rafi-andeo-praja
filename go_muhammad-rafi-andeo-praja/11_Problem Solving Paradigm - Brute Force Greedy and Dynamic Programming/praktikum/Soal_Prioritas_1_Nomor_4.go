package main

import "fmt"

func SimpleEquations(a, b, c int) {

	var x int = 1
	var y, z int
	var found bool

	for _, x = range []int{1, 2, 3, 4, 5, 6} {
		y = x
		for _, y = range []int{1, 2, 3, 4, 5, 6} {
			z = a - x - y
			if x*y*z == b && x*x+y*y+z*z == c {
				found = true
				break
			}
		}
		if found {
			fmt.Println(x, y, z)
			break
		}
	}
	if !found {
		fmt.Println("no solution")
	}

}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
