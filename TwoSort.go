package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	type data struct {
		input  []string
		output string
	}

	task0 := data{[]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}, "b***i***t***c***o***i***n"}
	task1 := data{[]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}, "a***r***e"}
	task2 := data{[]string{"lets", "talk", "about", "javascript", "the", "best", "language"}, "a***b***o***u***t"}
	task3 := data{[]string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}, "c***o***d***e"}
	task4 := data{[]string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}, "L***e***t***s"}

	var tasks [5]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = TwoSort(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func TwoSort(arr []string) string {
	sort.Strings(arr)
	return strings.Join(strings.Split(arr[0], ""), "***")
}
