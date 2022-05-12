package main

import (
	"fmt"
	"strings"
)

func main() {
	type data struct {
		input  string
		output bool
	}

	task0 := data{"a", true}
	task1 := data{"aba", true}
	task2 := data{"Abba", true}
	task3 := data{"hello", false}

	var tasks [4]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = IsPalindrome(tasks[i].input)

		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func IsPalindrome(str string) bool {
	var lowerStr string = strings.ToLower(str)
	var revStr string = ReverseString(lowerStr)
	return lowerStr == revStr
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome2(str string) bool {
	str = strings.ToLower(str)
	n := len(str)
	for i := 0; i < n; i++ {
		n -= 1
		if str[i] != str[n] {
			return false
		}
	}
	return true
}
