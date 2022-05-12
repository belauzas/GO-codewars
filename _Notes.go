package main

func EvenOrOdd(number int) string {
	return []string{"Even", "Odd"}[number&1]
}

func BoolToWord(word bool) string {
	return map[bool]string{false: "No", true: "Yes"}[word]
}

// import (
//         "net/http"
//         "io/ioutil"
//         "strings"
//         )

func greet() string {
	URL := "https://raw.githubusercontent.com/dmfed/go-exercises/master/hello.txt"
	resp, err := http.Get(URL)
	defer resp.Body.Close()
	if err != nil {
		return "Oh, noooo. Could not connect to the internet!"
	}
	out, _ := ioutil.ReadAll(resp.Body)
	return strings.Trim(string(out), "\n")
}

func MultiTable(number int) string {
	superstring := ""
	for i := 1; i < 11; i++ {
		superstring += fmt.Sprintf("%d * %d = %d\n", i, number, number*i)
	}
	return strings.TrimRight(superstring, "\n")
}

func seatsInTheater(nCols, nRows, col, row int) int {
	// return (nCols - col + 1) * (nRows - row)
	return (^nCols + col) * (row - nRows)
}
