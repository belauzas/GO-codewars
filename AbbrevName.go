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

	task0 := data{"Sam Harris", "S.H"}
	task1 := data{"Patrick Feenan", "P.F"}
	task2 := data{"Evan Cole", "E.C"}
	task3 := data{"P Favuzzi", "P.F"}
	task4 := data{"David Mendieta", "D.M"}

	var tasks [5]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2
	tasks[3] = task3
	tasks[4] = task4

	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = AbbrevName(tasks[i].input)
		fmt.Println(ans)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func AbbrevName(name string) string {
	words := strings.Split(name, " ")
	abb := string(words[0][0]) + "." + string(words[1][0])
	return strings.ToUpper(abb)
}

func AbbrevName2(name string) string {
	var x = strings.Index(name, " ")
	return strings.ToUpper(string(name[0]) + "." + string(name[x+1]))
}

func AbbrevName3(name string) string {
	var parts []string
	for _, part := range strings.Split(name, " ") {
		parts = append(parts, strings.ToUpper(part[:1]))
	}
	return strings.Join(parts, ".")
}

func AbbrevName4(name string) string {
	names := strings.Split(name, " ")
	return strings.ToUpper(fmt.Sprintf("%v.%v", names[0][:1], names[1][:1]))
}

func AbbrevName5(name string) string {
	name = strings.ToUpper(name)
	words := strings.Split(name, " ")
	var result []string
	for _, rune_ := range words {
		result = append(result, string(rune_[0]))
	}
	return strings.Join(result, ".")
}
