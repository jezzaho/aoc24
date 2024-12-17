package helpers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Helpers

func ToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return num
}

func AbsIntValue(n1, n2 int) int {
	if n1 >= n2 {
		return n1 - n2
	} else {
		return n2 - n1
	}
}

func FreqInListInt(n int, list []int) int {
	freq := 0
	for _, v := range list {
		if v == n {
			freq++
		}
	}

	return freq
}

func ReadInputFile(fileName string) []string {
	var text []string

	file, err := os.Open(fileName)
	if err != nil {
		return []string{}
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	defer file.Close()

	return text
}
func PrintResult(dayNumber, a, b int) {
	fmt.Printf("Day %v - part1: %v - part2: %v\n", dayNumber, a, b)
}

func ReadInputFileAsString(fileName string) string {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return ""
	}

	content := string(data)
	return content
}

func AllOccurencesOfRegex(pattern, text string) []string {
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(text, -1)

	return matches

}
