package main

import (
	"fmt"
)

func main() {
	type data struct {
		input  int
		output int
	}

	task0 := data{15, 7}
	task1 := data{15023, 7511}

	var tasks [2]data
	tasks[0] = task0
	tasks[1] = task1

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = OddCount(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func OddCount(n int) int {
	return (n - (n % 2)) / 2
}

func OddCount2(n int) int {
	return n / 2
}

func OddCount3(n int) int {
	return n >> 1
}
