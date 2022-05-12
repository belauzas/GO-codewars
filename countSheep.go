package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	type data struct {
		input  int
		output string
	}

	task0 := data{2, "1 sheep...2 sheep..."}
	task1 := data{0, ""}
	task2 := data{1, "1 sheep..."}

	var tasks [3]data
	tasks[0] = task0
	tasks[1] = task1
	tasks[2] = task2

	// fmt.Println(tasks)
	var correctCount int = 0

	for i := 0; i < len(tasks); i++ {
		var ans = countSheep(tasks[i].input)
		fmt.Println(ans)
		if ans == tasks[i].output {
			correctCount++
		}
	}
	fmt.Println("Correct:", correctCount, "/", len(tasks))
}

func countSheep(num int) string {
	var ans string = ""
	for i := 1; i <= num; i++ {
		ans += (strconv.Itoa(i) + " sheep...")
	}
	return ans
}

func countSheep2(num int) string {
	var sb strings.Builder

	for count := 1; count <= num; count++ {
		fmt.Fprintf(&sb, "%d sheep...", count)
	}

	return sb.String()
}

func countSheep3(num int) string {
	out := ""
	for i := 1; i <= num; i++ {
		out += fmt.Sprintf("%d sheep...", i)
	}
	return out
}

func countSheep4(num int) string {
	if num > 0 {
		return countSheep(num-1) + fmt.Sprintf("%d sheep...", num)
	}
	return ""
}
