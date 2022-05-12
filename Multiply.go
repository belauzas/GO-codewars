package main

import (
	"fmt"
)

func main() {
	task0 := [3]int{2, 2, 4}
	task1 := [3]int{4, 5, 20}
	task2 := [3]int{0, 100, 0}
	task3 := [3]int{1, 100, 100}

	var tasks [4][3]int
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		if Multiply(tasks[i][0], tasks[i][1]) == tasks[i][2] {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func Multiply(a int, b int) int {
	return a * b
}

func Multiply2(a, b int) int {
	return int(a * b)
}
