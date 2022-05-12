package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	type data struct {
		input  []string
		output []int
	}

	task0 := data{[]string{"abode", "ABc", "xyzD"}, []int{4, 3, 1}}
	task1 := data{[]string{"abide", "ABc", "xyz"}, []int{4, 3, 0}}
	task2 := data{[]string{"IAMDEFANDJKL", "thedefgh", "xyzDEFghijabc"}, []int{6, 5, 7}}
	task3 := data{[]string{"encode", "abc", "xyzD", "ABmD"}, []int{1, 3, 1, 3}}

	var tasks [4]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = solve(tasks[i].input)
		fmt.Println(ans, "->", tasks[i].output)
		// if ans == tasks[i].output {
		// 	correctCount++
		// }
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

const ABC = "abcdefghijklmnopqrstuvwxyz"

func solve(words []string) []int {
	var res []int
	for _, word := range words {
		word = strings.ToLower(word)
		count := 0
		for i, letter := range word {
			if strings.Index(ABC, string(letter)) == i {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

func solve2(slice []string) []int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	results := make([]int, len(slice))
	// Loop strings
	for sliceIndex, str := range slice {
		// Loop characters
		for charIndex, character := range str {
			// avoid match case
			if unicode.ToLower(character) == rune(alphabet[charIndex]) {
				results[sliceIndex]++
			}
		}
	}
	return results
}

func solve3(slice []string) []int {
	var output []int = make([]int, len(slice))
	for i, word := range slice {
		for position, char := range word {
			if int(char)%32 == position+1 {
				output[i]++
			}
		}
	}
	return output
}

func solve4(slice []string) (result []int) {
	for _, str := range slice {
		count := 0
		for chidx, ch := range strings.ToLower(str) {
			if int(ch-'a') == chidx {
				count++
			}
		}
		result = append(result, count)
	}
	return
}
