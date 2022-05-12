package main

import (
	"fmt"
	"strings"
)

type MyString string

func main() {

	type data struct {
		input  MyString
		output bool
	}

	task0 := data{"a", false}
	task1 := data{"A", true}

	var tasks [2]data
	tasks[0] = task0
	tasks[1] = task1

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = MyString(tasks[i].input).IsUpperCase()
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func (s MyString) IsUpperCase() bool {
	str := string(s)
	return str == strings.ToUpper(str)
}
