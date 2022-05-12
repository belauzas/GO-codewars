package main

import (
	"fmt"
	"math"
)

func main() {
	type data struct {
		input  []int
		output int
	}

	task0 := data{[]int{36, 7}, 22}
	task1 := data{[]int{55, 30}, 5}
	task2 := data{[]int{42, 21}, 0}
	task3 := data{[]int{22, 1}, 20}
	task4 := data{[]int{29, 0}, 29}
	task5 := data{[]int{36, 7}, 22}

	var tasks [6]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4
	tasks[5] = task5

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = TwiceAsOld(tasks[i].input[0], tasks[i].input[1])
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	return int(math.Abs(float64(dadYearsOld) - 2*float64(sonYearsOld)))
}
