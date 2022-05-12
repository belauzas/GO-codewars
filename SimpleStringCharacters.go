package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	type data struct {
		input  string
		output []int
	}

	task0 := data{"Codewars@codewars123.com", []int{1, 18, 3, 2}}
	task1 := data{"bgA5<1d-tOwUZTS8yQ", []int{7, 6, 3, 2}}
	task2 := data{"P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H", []int{9, 9, 6, 9}}
	task3 := data{"RYT'>s&gO-.CM9AKeH?,5317tWGpS<*x2ukXZD", []int{15, 8, 6, 9}}
	task4 := data{"$Cnl)Sr<7bBW-&qLHI!mY41ODe", []int{10, 7, 3, 6}}

	var tasks [5]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = Solve(tasks[i].input)
		fmt.Println(ans, "->", tasks[i].output)
		// if ans == tasks[i].output {
		// 	correctCount++
		// }
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

const UPPERCASE = "QWERTYUIOPASDFGHJKLZXCVBNM"
const LOWERCASE = "qwertyuiopasdfghjklzxcvbnm"
const NUMBERS = "0123456789"

func Solve(s string) []int {
	res := []int{0, 0, 0, 0}
	var isUpperCase bool
	var isLowerCase bool
	var isNumber bool
	var isSpecial bool
	for _, v := range s {
		isUpperCase = strings.Index(UPPERCASE, string(v)) >= 0
		isLowerCase = strings.Index(LOWERCASE, string(v)) >= 0
		isNumber = strings.Index(NUMBERS, string(v)) >= 0
		isSpecial = !isUpperCase && !isLowerCase && !isNumber
		if isUpperCase {
			res[0]++
		}
		if isLowerCase {
			res[1]++
		}
		if isNumber {
			res[2]++
		}
		if isSpecial {
			res[3]++
		}
	}
	return res
}

func Solve2(s string) []int {
	r := make([]int, 4)
	for _, c := range s {
		switch {
		case c >= 'A' && c <= 'Z':
			r[0]++
		case c >= 'a' && c <= 'z':
			r[1]++
		case c >= '0' && c <= '9':
			r[2]++
		default:
			r[3]++
		}
	}
	return r
}

func Solve3(s string) []int {
	uppercase := 0
	lowercase := 0
	number := 0
	character := 0
	for _, v := range s {
		if unicode.IsUpper(v) {
			uppercase++
		} else if unicode.IsLower(v) {
			lowercase++
		} else if unicode.IsNumber(v) {
			number++
		} else if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			character++
		}
	}
	return []int{uppercase, lowercase, number, character}
}
