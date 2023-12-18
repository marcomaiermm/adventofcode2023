package tests

import (
	"testing"

	"github.com/marcomaiermm/aoc2023day2/internal"
	lib "github.com/marcomaiermm/aoc2023day2/pkg"
)

func TestMain(t *testing.T) {
	lines := internal.ReadFromFile("../data_test")

	bag := lib.Balls{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	possibleGames := lib.GetPossibleGames(lines, bag)
	expected := 8
	if possibleGames != expected {
		t.Errorf("expected possible games to be %d; got %d", expected, possibleGames)
	}
}
