package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	type data struct {
		input  string
		output string
	}

	task0 := data{"most trees are blue", "Most Trees Are Blue"}
	task1 := data{"All the rules in this world were made by someone no smarter than you. So make your own.", "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."}
	task2 := data{"When I die. then you will realize", "When I Die. Then You Will Realize"}
	task3 := data{"Jonah Hill is a genius", "Jonah Hill Is A Genius"}
	task4 := data{"Dying is mainstream", "Dying Is Mainstream"}

	var tasks [5]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = ToJadenCase(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		letters := strings.Split(word, "")
		letters[0] = strings.ToUpper(letters[0])
		words[i] = strings.Join(letters, "")
	}
	return strings.Join(words, " ")
}

func ToJadenCase2(str string) string {
	return strings.Title(str)
}

func ToJadenCase3(str string) string {
	newWord := true
	ret := []rune(str)
	for i, r := range ret {
		if newWord {
			ret[i] = []rune(strings.ToUpper(string(ret[i])))[0]
			newWord = false
		}

		if !unicode.IsLetter(r) {
			newWord = true
		}
	}

	return string(ret) // your code here...
}

func ToJadenCase4(str string) string {
	arr := strings.Split(str, " ")
	for i, word := range arr {
		wordArr := strings.Split(word, "")
		arr[i] = strings.ToUpper(wordArr[0]) + strings.Join(wordArr[1:], "")
	}

	return strings.Join(arr, " ")
}
