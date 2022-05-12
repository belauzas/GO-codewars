package main

import (
	"fmt"
)

func main() {
	type data struct {
		input  []int
		output []int
	}

	task0 := data{[]int{1, 2, 3}, []int{2, 4, 6}}
	task1 := data{[]int{4, 1, 1, 1, 4}, []int{8, 2, 2, 2, 8}}
	task2 := data{[]int{2, 2, 2, 2, 2, 2}, []int{4, 4, 4, 4, 4, 4}}

	var tasks [3]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = Maps(tasks[i].input)
		var good bool = true

		for j := 0; j < len(ans); j++ {
			if ans[j] != tasks[i].output[j] {
				good = false
			}
		}

		if good {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func Maps(x []int) []int {
	var ans []int

	for i := 0; i < len(x); i++ {
		ans = append(ans, x[i]*2)
	}

	return ans
}
