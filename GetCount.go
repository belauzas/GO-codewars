package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	type data struct {
		input  string
		output int
	}

	task0 := data{"abracadabra", 5}

	var tasks [1]data
	tasks[0] = task0

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = GetCount(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

const VOWELS = "aeiou"

func GetCount(str string) (count int) {
	for _, v := range str {
		if strings.Index(VOWELS, string(v)) >= 0 {
			count++
		}
	}
	return count
}

func GetCount2(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func GetCoun3t(str string) (count int) {
	for i := range str {
		switch str[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func GetCount4(strn string) int {
	count := 0

	vowels := []string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(strn, vowel)
	}

	return count
}

func GetCount5(str string) (count int) {
	r := regexp.MustCompile("[aeiou]")
	vowels := r.FindAllString(str, -1)
	return len(vowels)
}

func GetCount6(str string) (count int) {
	for _, char := range str {
		if strings.ContainsRune("aeiou", char) {
			count++
		}
	}

	return count
}
