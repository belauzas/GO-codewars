package main

import (
	"fmt"
	"sort"
)

func main() {
	type data struct {
		input  []int
		output []int
	}

	task0 := data{[]int{1, 2, 10, 50, 5}, []int{1, 2, 5, 10, 50}}
	task1 := data{[]int{}, []int{}}

	var tasks [2]data
	tasks[0] = task0
	tasks[1] = task1

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = SortNumbers(tasks[i].input)
		fmt.Println(ans, "->", tasks[i].output)
		// if ans == tasks[i].output {
		// 	correctCount++
		// }
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func SortNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return numbers
	}
	sort.Ints(numbers)
	return numbers
}

func SortNumbers2(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[j] > numbers[i] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	return numbers
}
