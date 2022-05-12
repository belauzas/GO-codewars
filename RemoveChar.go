package main

import (
	"fmt"
)

func main() {
	type data struct {
		input  string
		output string
	}

	task0 := data{"eloquent", "loquen"}
	task1 := data{"country", "ountr"}
	task2 := data{"person", "erso"}
	task3 := data{"place", "lac"}

	var tasks [4]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = RemoveChar(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func RemoveChar(word string) (res string) {
	size := len(word) - 1
	for i, v := range word {
		if i != 0 && i != size {
			res += string(v)
		}
	}
	return res
}

func RemoveChar2(word string) string {
	return word[1 : len(word)-1]
}

func RemoveChar3(word string) string {
	var newWord = []rune(word)
	return string(newWord[1 : len(newWord)-1])
}
