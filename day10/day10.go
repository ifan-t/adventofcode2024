package main

import (
	"fmt"

	helper "github.com/ifan-t/adventofcode2024"
)

type Vector struct {
	x int
	y int
}

var directions = []Vector{Vector{x: -1, y: 0}, Vector{x: 1, y: 0}, Vector{x: 0, y: 1}, Vector{x: 0, y: -1}}

func main() {
	lines, _ := helper.ParseInput("input.txt")
	part1Answer := part1(lines)
	fmt.Println("Answer to part1: ", part1Answer)

}

func part1(input []string) int {
	sum := 0
	for y, _ := range input {
		for x, _ := range input[0] {
			if input[y][x] == '0' {
				visited := make(map[Vector]bool)
				endpoints := dfs(Vector{x: x, y: y}, 0, input, visited)
				sum += len(endpoints)
			}
		}
	}

	return sum
}

func dfs(position Vector, current int, input []string, visited map[Vector]bool) map[Vector]bool {
	if position.y < 0 || position.y >= len(input) || position.x < 0 || position.x >= len(input[0]) {
		return nil
	}
	length := int(input[position.y][position.x] - '0')
	if length != current {
		return nil
	}

	if length == 9 {
		return map[Vector]bool{position: true}
	}

	if visited[position] {
		return nil
	}
	visited[position] = true

	result := make(map[Vector]bool)
	for _, v := range directions {
		next := Vector{x: position.x + v.x, y: position.y + v.y}
		matches := dfs(next, current+1, input, visited)
		for match := range matches {
			result[match] = true
		}
	}
	return result
}
