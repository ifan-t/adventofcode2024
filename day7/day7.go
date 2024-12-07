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
	}
	part1Answer := part1(arr)
	fmt.Println("Answer to part 1 is: ", part1Answer)
}

func part1(inputs [][]int) int {
	sum := 0
	for _, input := range inputs {
		if tryCalcTarget(input[0], input[1:], 0, true, 0) || tryCalcTarget(input[0], input[1:], 0, false, 0) {
			sum += input[0]
		}
	}
	return sum
}

func tryCalcTarget(target int, inputs []int, index int, add bool, sum int) bool {
	if index == len(inputs)-1 {
		return target == sum
	}
	if sum > target {
		return false
	}
	if index == 0 {
		sum = inputs[0]
	}
	if add {
		sum += inputs[index+1]
	} else {
		sum *= inputs[index+1]
	}
	tryAdd := tryCalcTarget(target, inputs, index+1, true, sum)
	tryMultiply := tryCalcTarget(target, inputs, index+1, false, sum)
	return tryAdd || tryMultiply
}
