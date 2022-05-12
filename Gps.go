package main

import (
	"fmt"
)

func main() {
	type data struct {
		input1 int
		input2 []float64
		output int
	}

	task0 := data{20, []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}, 41}
	task1 := data{12, []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25}, 219}

	var tasks [2]data
	tasks[0] = task0
	tasks[1] = task1

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = Gps(tasks[i].input1, tasks[i].input2)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func Gps(s int, x []float64) int {
	var max float64 = 0
	for i := 1; i < len(x); i++ {
		if x[i]-x[i-1] > max {
			max = x[i] - x[i-1]
		}
	}
	return int(3600*max) / s
}
