package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	helper "github.com/ifan-t/adventofcode2024"
)

var rules = make(map[int][]int)

func main() {

	lines, _ := helper.ParseInput("input.txt")
	var updates [][]int
	isFirstSection := true
	for _, line := range lines {
		if line == "" {
			isFirstSection = false
			continue
		}

		if isFirstSection {
			parts := strings.Split(line, "|")
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])
			rules[num1] = append(rules[num1], num2)
		} else {
			strNumbers := strings.Split(line, ",")
			var intArray []int
			for _, strNum := range strNumbers {
				num, _ := strconv.Atoi(strings.TrimSpace(strNum))
				intArray = append(intArray, num)
			}
			updates = append(updates, intArray)
		}
	}
	part1Answer := part1(updates)
	fmt.Println("Answer to Part 1: ", part1Answer)
}

func part1(updates [][]int) int {
	result := 0
	for _, update := range updates {
		l := len(update)
		if isOrdered(update) {
			result += update[l/2]
		}
	}
	return result
}

func isOrdered(update []int) bool {
	for i, num1 := range update {
		for j := i + 1; j < len(update); j++ {
			num2 := update[j]
			if !AreInOrder(num1, num2) {
				return false
			}
		}
	}
	return true
}

func AreInOrder(num1 int, num2 int) bool {
	arr, ok := rules[num2]
	if ok && slices.Contains(arr, num1) {
		return false
	}
	return true
}
