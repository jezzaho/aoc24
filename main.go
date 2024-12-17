package main

import (
	"github.com/jezzaho/aoc24/helpers"
)

func main() {

	D1Input := helpers.ReadInputFile("input/day1a.txt")
	a, b := Day1(D1Input)
	helpers.PrintResult(1, a, b)

	D2Input := helpers.ReadInputFile("input/day2.txt")
	a, b = Day2a(D2Input), Day2b(D2Input)
	helpers.PrintResult(2, a, b)

	D3Input := helpers.ReadInputFileAsString("input/day3.txt")
	a, b = Day3a(D3Input), Day3b(D3Input)
	helpers.PrintResult(3, a, b)

	D4Input := helpers.ReadInputFileAsString("input/day4.txt")

}
