package day3

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Number struct {
	Value int
	Row   int
	Start int
	End   int
}

type Symbol struct {
	Value  string
	Row    int
	Column int
}

func Solve1(input string) {

	var numbers []Number
	var symbols []Symbol

	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	rows := strings.Split(input, "\n")

	for i, row := range rows {

		split := strings.Split(row, "")
		numberFlag := false
		number := ""

		for j, el := range split {

			if !slices.Contains(digits, el) && numberFlag {
				// Number finished, must parse
				value, err := strconv.ParseInt(number, 10, 64)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, Number{int(value), i, j - len(number), j - 1})
				number = ""
				numberFlag = false
			}
			if slices.Contains(digits, el) {
				number = number + el
				numberFlag = true
			}
			if !slices.Contains(digits, el) && el != "." {
				symbols = append(symbols, Symbol{el, i, j})
			}
		}
	}

	var total int

	for _, number := range numbers {

		partNumber := false

		for _, symbol := range symbols {

			// Same row
			if symbol.Row == number.Row {

				// Adjacent Left or Right
				if (symbol.Column-1 == number.End) || (symbol.Column+1 == number.Start) {
					partNumber = true
				}
			}

			// Above or Below
			if (symbol.Row-1 == number.Row) || (symbol.Row+1 == number.Row) {

				// Adjacent Above or Below
				if (symbol.Column >= number.Start) && (symbol.Column <= number.End) {
					partNumber = true
				}

				// Adjacent Diagonal Left
				if (symbol.Column-1 >= number.Start) && (symbol.Column-1 <= number.End) {
					partNumber = true
				}

				// Adjacent Diagonal Right
				if (symbol.Column+1 >= number.Start) && (symbol.Column+1 <= number.End) {
					partNumber = true
				}
			}
		}

		if partNumber {
			total = total + number.Value
		}
	}

	fmt.Println(total)
}

func Solve2(input string) {

}
