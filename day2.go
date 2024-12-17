package main

import (
	"strings"

	"github.com/jezzaho/aoc24/helpers"
)

func Day2a(input []string) int {
	numbers := IntInputs(input)
	counter := 0
	for _, item := range numbers {
		if isSafePartA(item) {
			counter++
		}
	}

	return counter
}

func Day2b(input []string) int {
	numbers := IntInputs(input)
	counter := 0
	for _, item := range numbers {
		if isSafePartB(item) {
			counter++
		}
	}

	return counter
}

func IntInputs(input []string) [][]int {
	var numbers [][]int

	for _, inp := range input {
		var row []int
		rowStr := strings.Split(inp, " ")
		for i := 0; i < len(rowStr); i++ {
			row = append(row, helpers.ToInt(rowStr[i]))
		}
		numbers = append(numbers, row)
	}

	return numbers
}

func isSafePartA(input []int) bool {
	return isSortedDescAsc(input) && isGradual(input)
}

func isSafePartB(input []int) bool {
	if isSafePartA(input) {
		return true
	}

	for i := 0; i < len(input); i++ {
		temp := make([]int, 0, len(input)-1)
		temp = append(temp, input[:i]...)
		temp = append(temp, input[i+1:]...)

		if isSortedDescAsc(temp) && isGradual(temp) {
			return true
		}
	}

	return false
}

func isSortedDescAsc(list []int) bool {
	asc := true
	desc := true

	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			asc = false
		}
		if list[i] >= list[i-1] {
			desc = false
		}
	}

	return asc || desc
}

func isGradual(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		for j := 1; j < len(input); j++ {
			diff := helpers.AbsIntValue(input[i], input[j])
			if diff < 1 || diff > 3 {
				return false
			}
			i++
		}
	}
	return true
}
