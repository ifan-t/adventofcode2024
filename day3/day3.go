package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	helper "github.com/ifan-t/adventofcode2024"
)

func main() {
	lines, _ := helper.ParseInput("input.txt")
	part1Pattern := regexp.MustCompile(`\bmul\((\d+),(\d+)\)`)
	var part1Answer = part1(lines, part1Pattern)
	fmt.Println("Answer to Part 1 is: ", part1Answer)
	part2Pattern := regexp.MustCompile(`\bmul\((\d+),(\d+)\)|\bdo\(\)|\bdon't\(\)`)
	var part2Answer = part2(lines, part2Pattern)
	fmt.Println("Answer to Part2 is: ", part2Answer)
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

func part2(lines []string, pattern *regexp.Regexp) int {
	numberOfMatches := 0
	var do = true
	for _, str := range lines {
		matches := pattern.FindAllStringSubmatch(str, -1)
		for _, match := range matches {
			num1, err := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			if do && err == nil {
				numberOfMatches += num1 * num2
				continue
			}
			if err != nil && strings.Contains(match[0], "do()") {
				do = true
				continue
			}
			if err != nil && strings.Contains(match[0], "don't()") {
				do = false
				continue
			}
		}
	}
	return numberOfMatches
}
