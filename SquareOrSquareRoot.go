package main

import (
	"fmt"
	"math"
)

func main() {
	type data struct {
		input  []int
		output []int
	}

	task0 := data{[]int{4, 3, 9, 7, 2, 1}, []int{2, 9, 3, 49, 4, 1}}
	task1 := data{[]int{100, 101, 5, 5, 1, 1}, []int{10, 10201, 25, 25, 1, 1}}
	task2 := data{[]int{1, 2, 3, 4, 5, 6}, []int{1, 4, 9, 2, 25, 36}}

	var tasks [3]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = SquareOrSquareRoot(tasks[i].input)
		fmt.Println(ans, "->", tasks[i].output)
		// if ans == tasks[i].output {
		// 	correctCount++
		// }
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func SquareOrSquareRoot(arr []int) []int {
	var res []int
	for _, v := range arr {
		n := math.Sqrt(float64(v))
		if math.Remainder(n, 1) == 0 {
			res = append(res, int(n))
		} else {
			res = append(res, v*v)
		}
	}
	return res
}

func SquareOrSquareRoot2(arr []int) []int {
	results := make([]int, len(arr))

	for i, x := range arr {
		sqrt := math.Sqrt(float64(x))

		if sqrt == math.Floor(sqrt) {
			results[i] = int(sqrt)
		} else {
			results[i] = x * x
		}
	}

	return results
}

func SquareOrSquareRoot3(arr []int) []int {
	var r []int
	for _, num := range arr {
		s := math.Sqrt(float64(num))
		if math.Mod(s, 1.0) == 0 {
			r = append(r, int(s))
		} else {
			r = append(r, num*num)
		}
	}

	return r
}
