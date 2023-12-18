package main

import (
	"fmt"

	"github.com/marcomaiermm/aoc2023day2/internal"
	lib "github.com/marcomaiermm/aoc2023day2/pkg"
)

func main() {
	lines := internal.ReadFromFile("data")
	bag := lib.Balls{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	possibleGames := lib.GetPossibleGames(lines, bag)

	fmt.Println(possibleGames)
}
