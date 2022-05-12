package main

import "fmt"

func main() {
	var correctCount int = 0

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	res := []int{10, -65}

	fmt.Println("#########")
	fmt.Println(arr)
	fmt.Println(res)
	fmt.Println("#########")
	var ans []int = CountPositivesSumNegatives(arr)
	fmt.Println(ans)
	fmt.Println("#########")
	if ans[0] == res[0] && ans[1] == res[1] {
		correctCount++
	}

	fmt.Println("Correct:", correctCount, "/", 1)
}

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	var positives = 0
	var negatives = 0

	for _, v := range numbers {
		if v > 0 {
			positives++
		} else {
			negatives += v
		}
	}

	res = append(res, positives)
	res = append(res, negatives)

	return res
}

func CountPositivesSumNegatives2(numbers []int) []int {
	res := []int{0, 0}
	for _, v := range numbers {
		if v > 0 {
			res[0] += 1
		} else {
			res[1] += v
		}
	}
	return res // your code here
}

func CountPositivesSumNegatives3(numbers []int) []int {
	var countOfPositives, sumOfNegatives int
	for _, number := range numbers {
		if number < 0 {
			sumOfNegatives += number
		} else if number > 0 {
			countOfPositives++
		}
	}
	return []int{
		countOfPositives,
		sumOfNegatives,
	}
}
