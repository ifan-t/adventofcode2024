package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	helper "github.com/ifan-t/adventofcode2024"
)

func main() {
	lines, err := helper.ParseInput("input.txt")
	if err != nil {
		fmt.Println("Failed to parse input")
	}

	var column1 []int
	var column2 []int

	for _, line := range lines {
		columns := strings.Fields(line)

		if len(columns) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		num1, _ := strconv.Atoi(columns[0])
		num2, _ := strconv.Atoi(columns[1])

		column1 = append(column1, num1)
		column2 = append(column2, num2)
	}
	sort.Ints(column1)
	sort.Ints(column2)

	var part1Answer = Part1(column1, column2)
	fmt.Println("Answer to part 1:", part1Answer)

	var part2Answer = Part2(column1, column2)
	fmt.Println("Answer to part 2:", part2Answer)

}

func Part1(column1 []int, column2 []int) int {
	sum := 0
	length := len(column1)
	for i := 0; i < length; i++ {
		distance := column1[i] - column2[i]
		if distance < 0 {
			distance = -distance
		}
		sum += distance
	}
	return sum
}

func Part2(column1 []int, column2 []int) int {
	similarityScore := 0
	previousIndex := 0
	length := len(column1)
	for i := 0; i < length; i++ {
		matches := 0
		for j := previousIndex; j < length; j++ {
			if column1[i] > column2[j] {
				continue
			}
			if column1[i] < column2[j] {
				break
			}
			if column1[i] == column2[j] {
				for j < length && column1[i] == column2[j] {
					matches++
					j++
				}
				previousIndex = j
				break
			}
		}
		similarityScore += matches * column1[i]
	}
	return similarityScore
}
