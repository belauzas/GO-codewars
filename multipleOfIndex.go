package main

import (
	"fmt"
)

func main() {
	type data struct {
		input  []int
		output []int
	}

	task0 := data{[]int{22, -6, 32, 82, 9, 25}, []int{-6, 32, 25}}
	task1 := data{[]int{68, -1, 1, -7, 10, 10}, []int{-1, 10}}
	task2 := data{[]int{11, -11}, []int{-11}}
	task3 := data{[]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}, []int{-85, 72, 0, 68}}
	task4 := data{[]int{28, 38, -44, -99, -13, -54, 77, -51}, []int{38, -44, -99}}
	task5 := data{[]int{-1, -49, -1, 67, 8, -60, 39, 35}, []int{-49, 8, -60, 35}}

	var tasks [6]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4
	tasks[5] = task5

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = multipleOfIndex(tasks[i].input)
		fmt.Println(ans, "->", tasks[i].output)
		// if ans == tasks[i].output {
		// 	correctCount++
		// }
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func multipleOfIndex(ints []int) (res []int) {
	for i, v := range ints {
		if i > 0 && v%i == 0 {
			res = append(res, v)
		}
	}
	return
}

func multipleOfIndex2(ints []int) []int {
	res := make([]int, 0)

	for i, v := range ints {
		if i > 0 && v%i == 0 {
			res = append(res, v)
		}
	}
	return res
}

func multipleOfIndex3(ints []int) []int {
	var arr []int
	for i, val := range ints[1:] {
		if val%(i+1) == 0 {
			arr = append(arr, val)
		}
	}
	return arr
}

func multipleOfIndex4(ints []int) []int {
	v := append(ints[:0:0], ints...)
	n := 0
	for i, u := range v {
		if i > 0 && u%i == 0 {
			v[n] = u
			n++
		}
	}
	return v[:n]
}
