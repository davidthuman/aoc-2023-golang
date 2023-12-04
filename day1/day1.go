package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type RowNumber struct {
	number   int64
	assigned bool
}

func Solve1(input string) {

	data := strings.Split(input, "\n")

	var lineNumber []int

	for _, d := range data {

		first := RowNumber{0, false}
		last := RowNumber{0, false}

		for _, s := range strings.Split(d, "") {

			num, err := strconv.ParseInt(s, 10, 64)
			if err == nil {

				if !first.assigned {
					first.number = num
					first.assigned = true
				}
				last.number = num
				last.assigned = true

			}
		}

		if !last.assigned {
			last.number = first.number
			last.assigned = true
		}

		d1 := strconv.FormatInt(first.number, 10)
		d2 := strconv.FormatInt(last.number, 10)

		finalNumber, err := strconv.ParseInt(d1+d2, 10, 64)
		if err != nil {
			panic(err)
		}
		lineNumber = append(lineNumber, int(finalNumber))

	}

	result := 0

	for _, v := range lineNumber {
		result = result + v
	}

	fmt.Println(result)
}

type NumberPosition struct {
	Value    int
	Position int
}

type ByPosition []NumberPosition

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

func Indices(s string, substr string, digits map[string]int, words map[string]int) []NumberPosition {

	var result []NumberPosition
	index := strings.Index(s, substr)
	for index != -1 {

		var number int
		if len(substr) == 1 {
			number = digits[substr]
		} else {
			number = words[substr]
		}

		result = append(result, NumberPosition{number, index})
		s = s[:index] + "_" + s[index+1:]

		index = strings.Index(s, substr)
	}
	return result
}

func Solve2(input string) {

	words := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	digits := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

	lines := strings.Split(input, "\n")

	var lineNumber []int

	for _, line := range lines {

		var lineInfo []NumberPosition

		for word, _ := range words {
			lineInfo = append(lineInfo, Indices(line, word, digits, words)...)
		}

		for digit, _ := range digits {
			lineInfo = append(lineInfo, Indices(line, digit, digits, words)...)
		}

		sort.Sort(ByPosition(lineInfo))
		d1 := strconv.FormatInt(int64(lineInfo[0].Value), 10)
		d2 := strconv.FormatInt(int64(lineInfo[len(lineInfo)-1].Value), 10)

		num, err := strconv.ParseInt(d1+d2, 10, 64)
		if err != nil {
			panic(err)
		}
		lineNumber = append(lineNumber, int(num))
	}

	result := 0
	for _, v := range lineNumber {
		result = result + v
	}

	fmt.Println(result)

}
