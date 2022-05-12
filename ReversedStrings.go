package main

import (
	"fmt"
	"strings"
)

func main() {
	type data struct {
		input  string
		output string
	}

	task0 := data{"world", "dlrow"}
	task1 := data{"", ""}

	var tasks [2]data
	tasks[0] = task0
	tasks[1] = task1

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = Solution(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func Solution(word string) (res string) {
	for _, letterRune := range word {
		res = string(letterRune) + res
	}
	return
}

func Solution2(word string) string {
	var b strings.Builder
	for i := len(word) - 1; i > -1; i-- {
		b.WriteByte(word[i])
	}
	return b.String()
}

func Solution3(word string) (reversed string) {
	for i := len(word) - 1; i >= 0; i-- {
		reversed += word[i : i+1]
	}
	return
}

func Solution4(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
