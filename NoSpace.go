package main

import (
	"fmt"
	"strings"
)

func main() {
	task0 := [2]string{
		"8 j 8   mBliB8g  imjB8B8  jl  B",
		"8j8mBliB8gimjB8B8jlB"}
	task1 := [2]string{
		"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd",
		"88Bifk8hB8BB8BBBB888chl8BhBfd"}
	task2 := [2]string{
		"8aaaaa dddd r     ",
		"8aaaaaddddr"}
	task3 := [2]string{
		"jfBm  gk lf8hg  88lbe8 ",
		"jfBmgklf8hg88lbe8"}
	task4 := [2]string{
		"8j aam",
		"8jaam"}
	task5 := [2]string{
		"",
		""}

	var tasks [6][2]string
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4
	tasks[5] = task5

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		if NoSpace(tasks[i][0]) == tasks[i][1] {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func NoSpace(word string) string {
	var newWord string = ""

	for i := 0; i < len(word); i++ {
		var letter string = string(word[i])
		if letter != " " {
			newWord += letter
		}

	}

	return newWord
}

func NoSpace2(word string) string {
	stArray := strings.Split(word, " ")
	return strings.Join(stArray, "")
}

func NoSpace3(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func NoSpace4(word string) string {
	arr := strings.Fields(word)
	return strings.Join(arr, "")
}

func NoSpace5(word string) (result string) {
	for _, v := range word {
		if v == ' ' {
			continue
		}
		result += string(v)
	}
	return result
}
