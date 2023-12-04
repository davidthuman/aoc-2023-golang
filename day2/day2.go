package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	MinBlue  int
	MinRed   int
	MinGreen int
}

func Solve1(input string) {

	allowedColorNumbers := map[string]int{"red": 12, "green": 13, "blue": 14}
	var allowedGames []int

	games := strings.Split(input, "\n")

	for i, game := range games {

		possibleGame := true

		colonIndex := strings.Index(game, ":")

		sets := strings.Split(game[colonIndex+2:], "; ")

		for _, set := range sets {

			possibleSet := true

			colors := strings.Split(set, ", ")

			for _, color := range colors {

				split := strings.Split(color, " ")

				value, err := strconv.ParseInt(split[0], 10, 64)
				if err != nil {
					panic(err)
				}
				possible := int(value) <= allowedColorNumbers[split[1]]
				possibleSet = possibleSet && possible
			}
			possibleGame = possibleGame && possibleSet
		}
		if possibleGame {
			allowedGames = append(allowedGames, i+1)
		}
	}

	result := 0
	for _, v := range allowedGames {
		result = result + v
	}

	fmt.Println(result)

}

func Solve2(input string) {

	var gameAmounts []Game

	games := strings.Split(input, "\n")

	for _, game := range games {

		gameAmount := Game{0, 0, 0}

		colonIndex := strings.Index(game, ":")

		sets := strings.Split(game[colonIndex+2:], "; ")

		for _, set := range sets {

			colors := strings.Split(set, ", ")

			for _, color := range colors {

				split := strings.Split(color, " ")

				value64, err := strconv.ParseInt(split[0], 10, 64)
				if err != nil {
					panic(err)
				}
				value := int(value64)

				switch split[1] {
				case "red":
					if gameAmount.MinRed < value {
						gameAmount.MinRed = value
					}
				case "green":
					if gameAmount.MinGreen < value {
						gameAmount.MinGreen = value
					}
				case "blue":
					if gameAmount.MinBlue < value {
						gameAmount.MinBlue = value
					}
				}

			}
		}
		gameAmounts = append(gameAmounts, gameAmount)
	}

	result := 0
	for _, v := range gameAmounts {
		result = result + (v.MinRed * v.MinGreen * v.MinBlue)
	}

	fmt.Println(result)

}
