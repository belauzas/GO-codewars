package main

import (
	"fmt"
)

func main() {
	type data struct {
		input  int
		output []int
	}

	task0 := data{1, []int{1}}
	task1 := data{2, []int{2, 1}}
	task2 := data{3, []int{3, 2, 1}}
	task3 := data{4, []int{4, 3, 2, 1}}
	task4 := data{5, []int{5, 4, 3, 2, 1}}

	var tasks [5]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = ReverseSeq(tasks[i].input)
		fmt.Println(ans, "->", tasks[i].output)
		// if ans == tasks[i].output {
		// 	correctCount++
		// }
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func ReverseSeq(n int) (res []int) {
	res = make([]int, n)

	for i := n; i > 0; i-- {
		res[i-1] = n - i + 1
	}

	return res
}

func ReverseSeq2(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i
	}
	return arr
}

func ReverseSeq3(n int) (result []int) {
	for i := n; i > 0; i-- {
		result = append(result, i)
	}
	return
}

func ReverseSeq4(n int) []int {
	ans := []int{}
	for n > 0 {
		ans = append(ans, n)
		n--
	}

	return ans
}

func ReverseSeq5(n int) []int {
	if n == 1 {
		return []int{1}
	}

	return append([]int{n}, ReverseSeq(n-1)...)
}
