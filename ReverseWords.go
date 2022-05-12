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

	task0 := data{"hello world!", "world! hello"}
	task1 := data{"yoda doesn't speak like this", "this like speak doesn't yoda"}
	task2 := data{"foobar", "foobar"}
	task3 := data{"kata editor", "editor kata"}
	task4 := data{"row row row your boat", "boat your row row row"}
	task5 := data{"", ""}

	var tasks [6]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4
	tasks[5] = task5

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = ReverseWords(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func ReverseWords(str string) string {
	arr := strings.Split(str, " ")
	var rev []string

	for _, v := range arr {
		rev = append([]string{v}, rev...)
	}

	return strings.Join(rev, " ")
}

func ReverseWords2(str string) string {
	words := strings.Split(str, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func ReverseWords3(str string) string {
	sli := strings.Split(str, " ")
	out := make([]string, len(sli))
	for i := len(sli) - 1; i >= 0; i-- {
		out[len(sli)-1-i] = sli[i]
	}
	return strings.Join(out, " ") // reverse those words
}

func ReverseWords4(str string) string {
	strSplit := strings.Split(str, " ")
	var res []string
	for i := len(strSplit) - 1; i >= 0; i-- {
		res = append(res, strSplit[i])
	}
	return strings.Join(res, " ")
}

func ReverseWords5(str string) string {
	words := strings.Split(str, " ")
	count := len(words)
	res := make([]string, 0, count)
	for i := count - 1; i >= 0; i-- {
		res = append(res, words[i])
	}

	return strings.Join(res, " ")
}

func ReverseWords6(str string) (result string) {
	arr := strings.Fields(str)
	for _, v := range arr {
		result = string(v) + " " + result
	}
	return strings.TrimSuffix(result, " ")
}
