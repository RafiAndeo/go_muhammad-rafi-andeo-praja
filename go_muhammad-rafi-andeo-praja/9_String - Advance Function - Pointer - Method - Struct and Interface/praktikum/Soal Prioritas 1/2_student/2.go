package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) average() int {

	var total int
	for i := 0; i < len(s.score); i++ {
		total += s.score[i]
	}
	return total / len(s.score)

}

func (s Student) highest() string {

	var max int
	var student string

	for i := 0; i < len(s.score); i++ {

		if s.score[i] > max {

			max = s.score[i]
			student = s.name[i]

		}

	}

	var score = strconv.Itoa(max)
	return student + "(" + score + ")"

}

func (s Student) lowest() string {

	var min = 100
	var student string

	for i := 0; i < len(s.score); i++ {

		if s.score[i] < min {

			min = s.score[i]
			student = s.name[i]

		}

	}

	var score = strconv.Itoa(min)
	return student + "(" + score + ")"

}

func main() {

	var s Student

	for i := 0; i < 5; i++ {

		var name string
		var score int
		index := strconv.Itoa(i + 1)
		fmt.Print("Input " + index + " Student's Name ")
		fmt.Scan(&name)
		fmt.Print("Input " + index + " Student's Score ")
		fmt.Scan(&score)
		s.name = append(s.name, name)
		s.score = append(s.score, score)

	}

	fmt.Println("Average Score: " + strconv.Itoa(s.average()))
	fmt.Println("Min Score of Students: " + s.lowest())
	fmt.Println("Max Score of Students: " + s.highest())

}
