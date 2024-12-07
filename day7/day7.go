package main

import (
	"fmt"
	"strconv"
	"strings"

	helper "github.com/ifan-t/adventofcode2024"
)

func main() {
	lines, _ := helper.ParseInput("input.txt")
	var arr [][]int
	for _, line := range lines {
		var nums []int
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(parts[0])
		nums = append(nums, target)
		inputs := strings.Split(parts[1], " ")
		for _, input := range inputs {
			num, err := strconv.Atoi(input)
			if err == nil {
				nums = append(nums, num)
			}
		}
		arr = append(arr, nums)
		concat(nums[1], nums[2])
	}
	part1Answer := part1(arr)
	fmt.Println("Answer to part 1 is: ", part1Answer)
	part2Answer := part2(arr)
	fmt.Println("Answer to part 2 is: ", part2Answer)
}

func part1(inputs [][]int) int {
	sum := 0
	for _, input := range inputs {
		if tryCalcTarget(input[0], input[1:], 0, add, 0) || tryCalcTarget(input[0], input[1:], 0, multiply, 0) {
			sum += input[0]
		}
	}
	return sum
}

func tryCalcTarget(target int, inputs []int, index int, operator func(int, int) int, sum int) bool {
	if index == len(inputs)-1 {
		return target == sum
	}
	if sum > target {
		return false
	}
	if index == 0 {
		sum = inputs[0]
	}
	sum = operator(sum, inputs[index+1])
	tryAdd := tryCalcTarget(target, inputs, index+1, add, sum)
	tryMultiply := tryCalcTarget(target, inputs, index+1, multiply, sum)
	return tryAdd || tryMultiply
}

func part2(inputs [][]int) int {
	sum := 0
	for _, input := range inputs {
		if tryCalcTargetWithConcat(input[0], input[1:], 0, add, 0) || tryCalcTargetWithConcat(input[0], input[1:], 0, multiply, 0) || tryCalcTargetWithConcat(input[0], input[1:], 0, concat, 0) {
			sum += input[0]
		}
	}
	return sum
}

func tryCalcTargetWithConcat(target int, inputs []int, index int, operator func(int, int) int, sum int) bool {
	if index == len(inputs)-1 {
		return target == sum
	}
	if sum > target {
		return false
	}
	if index == 0 {
		sum = inputs[0]
	}
	sum = operator(sum, inputs[index+1])
	tryAdd := tryCalcTargetWithConcat(target, inputs, index+1, add, sum)
	tryMultiply := tryCalcTargetWithConcat(target, inputs, index+1, multiply, sum)
	tryConcat := tryCalcTargetWithConcat(target, inputs, index+1, concat, sum)
	return tryAdd || tryMultiply || tryConcat
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func concat(num1 int, num2 int) int {
	i, _ := strconv.Atoi(strconv.Itoa(num1) + strconv.Itoa(num2))
	return i
}
