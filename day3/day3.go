package main

import (
	"fmt"
	"regexp"
	"strconv"

	helper "github.com/ifan-t/adventofcode2024"
)

func main() {
	lines, _ := helper.ParseInput("input.txt")
	pattern := regexp.MustCompile(`\bmul\((\d+),(\d+)\)`)
	var part1Answer = part1(lines, pattern)
	fmt.Println("Answer to Part 1 is: ", part1Answer)
}

func part1(lines []string, pattern *regexp.Regexp) int {
	numberOfMatches := 0
	for _, str := range lines {
		matches := pattern.FindAllStringSubmatch(str, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			numberOfMatches += num1 * num2
		}
	}
	return numberOfMatches
}
