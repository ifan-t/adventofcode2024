package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	helper "github.com/ifan-t/adventofcode2024"
)

func main() {
	lines, err := helper.ParseInput("input.txt")
	if err != nil {
		fmt.Println("Failed to parse input")
	}
	var input [][]float64
	for _, line := range lines {
		parts := strings.Fields(line)

		var report []float64

		for _, part := range parts {
			num, _ := strconv.ParseFloat(part, 64)
			report = append(report, num)
		}
		input = append(input, report)
	}

	var part1Answer = Part1(input)
	fmt.Println("Answer to Part 1 is : ", part1Answer)
	var part2Answer = Part2(input)
	fmt.Println("Answer to Part 2 is : ", part2Answer)
}

func Part1(input [][]float64) float64 {
	result := 0.0
	for _, report := range input {
		if listIsMonotone(report) {
			result++
		}
	}
	return result
}

func Part2(input [][]float64) float64 {
	result := 0.0
	for _, report := range input {
		if listIsMonotoneWithSkip(report) {
			result++
		}
	}
	return result
}

func listIsMonotone(report []float64) bool {
	length := len(report)
	rollingSum := 0.0
	previousSum := 0.0
	for index, _ := range report {
		if index == length-1 {
			continue
		}
		difference := report[index] - report[index+1]
		if math.Abs(difference) > 3 || math.Abs(difference) < 1 {
			return false
		}
		previousSum = rollingSum
		rollingSum += difference
		if math.Abs(rollingSum) < math.Abs((previousSum)) || (rollingSum < 0 && previousSum > 0) || (rollingSum > 0 && previousSum < 0) {
			return false
		}
	}
	return true
}

func listIsMonotoneWithSkip(report []float64) bool {
	length := len(report)
	rollingSum := 0.0
	previousSum := 0.0
	for index, _ := range report {
		if index == length-1 {
			continue
		}
		difference := report[index] - report[index+1]
		if math.Abs(difference) > 3 || math.Abs(difference) < 1 {
			return checkWithSkip(report, index)
		}
		previousSum = rollingSum
		rollingSum += difference
		if math.Abs(rollingSum) < math.Abs((previousSum)) || (rollingSum < 0 && previousSum > 0) || (rollingSum > 0 && previousSum < 0) {
			return checkWithSkip(report, index)
		}
	}
	return true
}

func checkWithSkip(report []float64, index int) bool {
	length := len(report)
	if listIsMonotone(append(append([]float64{}, report[:index]...), report[index+1:]...)) {
		return true
	}
	if index > 0 {
		if listIsMonotone(append(append([]float64{}, report[:index-1]...), report[index:]...)) {
			return true
		}
	}
	if index < length-1 {
		if listIsMonotone(append(append([]float64{}, report[:index+1]...), report[index+2:]...)) {
			return true
		}
	}
	return false
}
