package main

import (
	"fmt"
)

func main() {
	type nums struct {
		num  int
		div1 int
		div2 int
	}
	type data struct {
		input  nums
		output bool
	}

	task0 := data{nums{3, 1, 3}, true}
	task1 := data{nums{12, 2, 6}, true}
	task2 := data{nums{100, 5, 3}, false}
	task3 := data{nums{12, 7, 5}, false}
	task4 := data{nums{3, 3, 4}, false}
	task5 := data{nums{20, 4, 5}, true}

	var tasks [6]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4
	tasks[5] = task5

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var n1 int = tasks[i].input.num
		var n2 int = tasks[i].input.div1
		var n3 int = tasks[i].input.div2
		var ans = IsDivisible(n1, n2, n3)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func IsDivisible2(n, x, y int) bool {
	var resX bool = n%x == 0
	var resY bool = n%y == 0
	return resY && resX
}
