package tests

import (
	"aocday1first/lib"
	"testing"
)

func TestMain(t *testing.T) {
	test_data := []string{
		"two1nine",
		"eightwo",
		"threeight",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	sum := lib.SumUp(test_data)

	expected := 281 + 38 + 82
	if sum != expected {
		t.Errorf("expected sum to be %d; got %d", expected, sum)
	}
}
