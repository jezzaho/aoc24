package main

import (
	"strconv"

	"github.com/jezzaho/aoc24/helpers"
)

func Day3a(input string) int {

	pattern := `mul\(\b\d{1,3},\d{1,3}\)`

	matches := helpers.AllOccurencesOfRegex(pattern, input)

	totalSum := 0
	patternDigits := `\d{1,3}`

	for _, match := range matches {
		numbers := helpers.AllOccurencesOfRegex(patternDigits, match)

		n1, err := strconv.Atoi(numbers[0])
		if err != nil {
			return -1
		}
		n2, err := strconv.Atoi(numbers[1])
		if err != nil {
			return -1
		}

		totalSum += (n1 * n2)
	}
	return totalSum

}

func Day3b(input string) int {

	pattern := `do\(\)|don\'t\(\)|mul\(\b\d{1,3},\d{1,3}\)`

	matches := helpers.AllOccurencesOfRegex(pattern, input)

	totalSum := 0
	patternDigits := `\d{1,3}`

	enabled := true

	for _, match := range matches {
		numbers := helpers.AllOccurencesOfRegex(patternDigits, match)
		if len(numbers) != 0 && enabled {
			n1, err := strconv.Atoi(numbers[0])
			if err != nil {
				return -1
			}
			n2, err := strconv.Atoi(numbers[1])
			if err != nil {
				return -1
			}
			totalSum += (n1 * n2)
		} else if match == "do()" {
			enabled = true
		} else {
			enabled = false
		}

	}

	return totalSum
}
