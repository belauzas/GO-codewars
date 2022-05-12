package main

import (
	"fmt"
)

func main() {
	type data struct {
		input  []int
		output int
	}

	task0 := data{[]int{1, 2, 3, 4, 5}, 15}
	task1 := data{[]int{1, -2, 3, 4, 5}, 13}
	task2 := data{[]int{}, 0}
	task3 := data{[]int{-1, -2, -3, -4, -5}, 0}
	task4 := data{[]int{-1, 2, 3, 4, -5}, 9}

	var tasks [5]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = PositiveSum(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func PositiveSum(numbers []int) int {
	var sum int = 0
	for _, v := range numbers {
		if v > 0 {
			sum += v
		}
	}
	return sum
}
