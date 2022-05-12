package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	type funcInput struct {
		count  int
		letter string
	}
	type data struct {
		input  funcInput
		output string
	}

	task0 := data{funcInput{1, "i"}, "i"}
	task1 := data{funcInput{4, "a"}, "aaaa"}
	task2 := data{funcInput{3, "hello"}, "hellohellohello"}
	task3 := data{funcInput{2, "abc"}, "abcabc"}

	var tasks [4]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = RepeatStr(tasks[i].input.count, tasks[i].input.letter)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func RepeatStr2(repititions int, value string) (result string) {
	for i := 0; i < repititions; i++ {
		result += value
	}
	return
}

func RepeatStr3(rep int, v string) (result string) {
	for ; rep > 0; rep-- {
		result += v
	}
	return result
}

func RepeatStr4(rep int, value string) string {
	var res bytes.Buffer
	for i := 0; i < rep; i++ {
		res.WriteString(value)
	}
	return res.String()
}

func RepeatStr5(repetitions int, value string) string {
	length := len(value)
	if length == 0 {
		return value
	}
	var stb strings.Builder
	stb.Grow(repetitions * length)
	for i := 0; i < repetitions; i++ {
		stb.WriteString(value)
	}
	return stb.String()
}
