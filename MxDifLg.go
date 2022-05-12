package main

import (
	"fmt"
	"sort"
)

func main() {
	type data struct {
		input1 []string
		input2 []string
		output int
	}
	var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	task0 := data{s1, s2, 13}

	var s3 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	var s4 = []string{"bbbaaayddqbbrrrv"}
	task1 := data{s3, s4, 10}

	var tasks [2]data
	tasks[0] = task0
	tasks[1] = task1

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = MxDifLg(tasks[i].input1, tasks[i].input2)
		fmt.Println(ans, "->", tasks[i].output)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	var a1Size []int
	for _, v := range a1 {
		a1Size = append(a1Size, len(v))
	}
	var a2Size []int
	for _, v := range a2 {
		a2Size = append(a2Size, len(v))
	}
	sort.Ints(a1Size)
	sort.Ints(a2Size)
	dif1 := a1Size[len(a1Size)-1] - a2Size[0]
	dif2 := a2Size[len(a2Size)-1] - a1Size[0]
	if dif1 > dif2 {
		return dif1
	}
	return dif2
}
