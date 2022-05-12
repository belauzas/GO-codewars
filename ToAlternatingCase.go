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

	task0 := data{"hello world", "HELLO WORLD"}
	task1 := data{"HELLO WORLD", "hello world"}
	task2 := data{"HeLLo WoRLD", "hEllO wOrld"}
	task3 := data{"12345", "12345"}
	task4 := data{"1a2b3c4d5e", "1A2B3C4D5E"}
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
		var ans = ToAlternatingCase(tasks[i].input)
		fmt.Println(ans, "->", tasks[i].output)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func ToAlternatingCase(str string) string {
	arr := strings.Split(str, "")
	for i, v := range arr {
		if v == strings.ToLower(v) {
			arr[i] = strings.ToUpper(v)
		} else {
			arr[i] = strings.ToLower(v)
		}
	}
	return strings.Join(arr, "")
}

func ToAlternatingCase2(str string) string {
	result := []rune{}
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			result = append(result, unicode.ToLower(ch))
		} else if unicode.IsLower(ch) {
			result = append(result, unicode.ToUpper(ch))
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}

func ToAlternatingCase3(str string) (res string) {
	for _, c := range str {
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			c ^= 32
		}
		res += string(c)
	}
	return
}

func ToAlternatingCase4(str string) string {
	alternate := func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		} else {
			return unicode.ToLower(r)
		}
	}

	return strings.Map(alternate, str)
}

func ToAlternatingCase5(str string) string {
	var alternate string = ""

	for i := 0; i < len(str); i++ {
		var c int = int(str[i])

		if c >= 65 && c <= 90 {
			c += 32
		} else if c >= 97 && c <= 122 {
			c -= 32
		}

		alternate += string(c)
	}

	return alternate
}

func ToAlternatingCase6(str string) string {
	s := make([]string, len(str))
	for i, c := range str {
		w := string(c)
		if w == strings.ToUpper(w) {
			s[i] = strings.ToLower(w)
		} else {
			s[i] = strings.ToUpper(w)
		}
	}
	return strings.Join(s, "")
}
