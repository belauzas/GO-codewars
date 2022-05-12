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

	task0 := data{"GCAT", "GCAU"}

	var tasks [1]data
	tasks[0] = task0

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = DNAtoRNA(tasks[i].input)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func DNAtoRNA2(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}
