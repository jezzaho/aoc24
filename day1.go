package main

import (
	"slices"
	"strings"

	"github.com/jezzaho/aoc24/helpers"
)

func Day1(input []string) (int, int) {
	numbersLeft := make([]int, len(input))
	numbersRight := make([]int, len(input))

	for i, t := range input {
		nums := strings.Split(t, "   ")
		numbersLeft[i] = helpers.ToInt(nums[0])
		numbersRight[i] = helpers.ToInt(nums[1])
	}

	slices.Sort(numbersLeft)
	slices.Sort(numbersRight)

	distance := 0

	for i := 0; i < len(input); i++ {
		distance += helpers.AbsIntValue(numbersLeft[i], numbersRight[i])
	}

	simScore := 0

	for i := 0; i < len(numbersLeft); i++ {
		freq := helpers.FreqInListInt(numbersLeft[i], numbersRight)
		simScore += numbersLeft[i] * freq
	}

	return distance, simScore
}
