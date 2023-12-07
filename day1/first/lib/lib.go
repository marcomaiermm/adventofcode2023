package lib

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func addAtPosition(str string, char rune, index int) string {
	runes := []rune(str)

	if index < 0 || index > len(runes) {
		return str
	}

	runes = append(runes[:index], append([]rune{char}, runes[index:]...)...)

	return string(runes)
}

func findAndInsert(str string, key string) string {
	runes := []rune(key)
	index := strings.Index(str, key)

	// nothing found anymore
	if index == -1 {
		return str
	}

	// we found something with the key
	str = addAtPosition(str, runes[0], index)
	str = addAtPosition(str, runes[len(key)-1], index+len(key))

	str = strings.Replace(str, key, digitMap[key], 1)

	return findAndInsert(str, key)
}

func SumUp(numbers []string) int {
	pattern := "\\d"

	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Error compile regex", err)
	}

	digitRex := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")

	sum := 0
	for _, n := range numbers {
		for key := range digitMap {
			n = findAndInsert(n, key)
		}

		n = digitRex.ReplaceAllStringFunc(n, func(m string) string {
			return digitMap[m]
		})

		nums := r.FindAllString(n, -1)

		first := nums[0]
		second := nums[len(nums)-1]

		number, err := strconv.Atoi(first + second)
		if err != nil {
			log.Fatal(err)
		}

		sum += number
	}

	fmt.Println(sum)

	return sum
}
