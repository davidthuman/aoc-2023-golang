package main

import (
	"aoc-2023/day1"
	"aoc-2023/day2"
	"aoc-2023/day3"
	"fmt"
	"os"
)

var DAY = "3"
var SOLVE = "puzzle"
var PUZZLE = "1"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.ReadFile(fmt.Sprintf("./day%s/%s%s.txt", DAY, SOLVE, PUZZLE))
	check(err)

	switch DAY {
	case "1":
		if PUZZLE == "1" {
			day1.Solve1(string(file))
		} else {
			day1.Solve2(string(file))
		}
	case "2":
		if PUZZLE == "1" {
			day2.Solve1(string(file))
		} else {
			day2.Solve2(string(file))
		}
	case "3":
		if PUZZLE == "1" {
			day3.Solve1(string(file))
		} else {
			day3.Solve2(string(file))
		}
	}
}
