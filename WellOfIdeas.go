package main

import (
	"fmt"
)

func main() {
	type data struct {
		input  []string
		output string
	}

	task0 := data{[]string{"bad", "bad"}, "Fail!"}
	task1 := data{[]string{"bad", "bad", "bad", "bad", "bad"}, "Fail!"}
	task2 := data{[]string{"bad", "bad", "bad", "bad", "good", "bad"}, "Publish!"}
	task3 := data{[]string{"bad"}, "Fail!"}
	task4 := data{[]string{"bad", "bad", "bad", "bad", "bad", "good", "bad", "good", "good", "good", "bad", "bad", "good"}, "I smell a series!"}
	task5 := data{[]string{"bad", "bad", "bad", "bad", "good", "bad", "bad"}, "Publish!"}
	task6 := data{[]string{"bad", "bad", "good", "bad", "good", "good", "good", "bad", "good"}, "I smell a series!"}

	var tasks [7]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4
	tasks[5] = task5
	tasks[6] = task6

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = Well(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func Well(x []string) string {
	var count int = 0

	for _, v := range x {
		if v == "good" {
			count++
		}
	}

	if count == 0 {
		return "Fail!"
	}
	if count < 3 {
		return "Publish!"
	}
	return "I smell a series!"
}

func Well2(x []string) string {
	good := 0
	for _, idea := range x {
		switch idea {
		case "good":
			good++
		}
	}

	switch {
	case good > 2:
		return "I smell a series!"
	case good > 0:
		return "Publish!"
	default:
		return "Fail!"
	}
}
