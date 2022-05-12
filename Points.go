package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	type data struct {
		input  []string
		output int
	}

	task0 := data{[]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}, 30}
	task1 := data{[]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}, 10}
	task2 := data{[]string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}, 0}
	task3 := data{[]string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"}, 15}
	task4 := data{[]string{"1:0", "2:0", "3:0", "4:4", "2:2", "3:3", "1:4", "2:3", "2:4", "3:4"}, 12}

	var tasks [5]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = Points(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func Points(games []string) (res int) {
	res = 0

	for _, game := range games {
		var points = strings.Split(game, ":")
		first, firstErr := strconv.Atoi(points[0])
		second, secondErr := strconv.Atoi(points[1])
		if firstErr == nil && secondErr == nil {
			if first > second {
				res += 3
			}
			if first == second {
				res += 1
			}
		}
	}

	return res
}

func Points2(games []string) int {
	result := 0
	for _, game := range games {
		str := strings.Split(game, ":")
		x, y := str[0], str[1]
		switch {
		case x > y:
			result += 3
		case x == y:
			result += 1
		}
	}
	return result
}

func Points3(games []string) int {
	points := 0
	for _, g := range games {
		if g[0] > g[2] {
			points += 3
		} else if g[0] == g[2] {
			points += 1
		}
	}
	return points
}

func Points4(games []string) int {
	const WIN_POINTS = 3
	const DRAW_POINTS = 1
	var totalPoints int
	for _, val := range games {
		x := string(val[:1])
		y := string(val[2:])

		if x > y {
			totalPoints += WIN_POINTS
		} else if x == y {
			totalPoints += DRAW_POINTS
		}
	}
	return totalPoints
}
