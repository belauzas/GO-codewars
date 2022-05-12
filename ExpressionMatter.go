package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	type data struct {
		input  []int
		output int
	}

	task0 := data{[]int{2, 1, 2}, 6}
	task1 := data{[]int{2, 1, 1}, 4}
	task2 := data{[]int{1, 1, 1}, 3}
	task3 := data{[]int{1, 2, 3}, 9}
	task4 := data{[]int{1, 3, 1}, 5}
	task5 := data{[]int{2, 2, 2}, 8}
	task6 := data{[]int{5, 1, 3}, 20}
	task7 := data{[]int{3, 5, 7}, 105}
	task8 := data{[]int{5, 6, 1}, 35}
	task9 := data{[]int{1, 6, 1}, 8}
	task10 := data{[]int{2, 6, 1}, 14}
	task11 := data{[]int{6, 7, 1}, 48}
	task12 := data{[]int{2, 10, 3}, 60}
	task13 := data{[]int{1, 8, 3}, 27}
	task14 := data{[]int{9, 7, 2}, 126}
	task15 := data{[]int{1, 1, 10}, 20}
	task16 := data{[]int{9, 1, 1}, 18}
	task17 := data{[]int{10, 5, 6}, 300}
	task18 := data{[]int{1, 10, 1}, 12}

	var tasks [19]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4
	tasks[5] = task5
	tasks[6] = task6
	tasks[7] = task7
	tasks[8] = task8
	tasks[9] = task9
	tasks[10] = task10
	tasks[11] = task11
	tasks[12] = task12
	tasks[13] = task13
	tasks[14] = task14
	tasks[15] = task15
	tasks[16] = task16
	tasks[17] = task17
	tasks[18] = task18

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = ExpressionMatter(tasks[i].input[0], tasks[i].input[1], tasks[i].input[2])
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func ExpressionMatter(a, b, c int) (res int) {
	n1 := a + b + c
	n2 := a + b*c
	n3 := a*b + c
	n4 := (a + b) * c
	n5 := a * (b + c)
	n6 := a * b * c
	all := [...]int{n1, n2, n3, n4, n5, n6}
	res = all[0]

	for _, v := range all {
		if v > res {
			res = v
		}
	}

	return
}

func ExpressionMatter2(a int, b int, c int) int {
	arr := []int{a * (b + c), a * b * c, a + b + c, (a + b) * c}
	sort.Ints(arr)
	return arr[len(arr)-1]
}

func ExpressionMatter3(a int, b int, c int) int {
	return int(
		math.Max(
			math.Max(
				math.Max(
					float64(a*b*c),
					float64(a+b+c)),
				float64((a+b)*c),
			),
			float64(a*(b+c)),
		),
	)
}
