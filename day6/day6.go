package main

import (
	"fmt"

	helper "github.com/ifan-t/adventofcode2024"
)

type Vector struct {
	x int
	y int
}

func main() {
	lines, _ := helper.ParseInput("input.txt")
	var startPosition Vector
	for row, line := range lines {
		for col, char := range line {
			if char == '^' {
				startPosition = Vector{x: col, y: row}
			}
		}
	}
	part1Answer := part1(startPosition, Vector{x: 0, y: -1}, lines)
	fmt.Println("Answer to part1: ", part1Answer)
	part2Answer := part2(startPosition, Vector{x: 0, y: -1}, lines)
	fmt.Println("Answer to part2: ", part2Answer)
}

func part1(position Vector, direction Vector, grid []string) int {
	visited := make(map[Vector]struct{})
	for {
		visited[position] = struct{}{}
		if position.x <= 0 || position.x >= len(grid[0])-1 || position.y <= 0 || position.y >= len(grid)-1 {
			break
		}
		next := position.Add(direction)
		if grid[next.y][next.x] == '#' {
			direction.TurnRight()
			position = position.Add(direction)
			continue
		}
		position = next
	}
	return len(visited)
}

func part2(position Vector, direction Vector, grid []string) int {
	sum := 0
	for row, line := range grid {
		for col, char := range line {
			if char == '#' || char == '^' {
				continue
			}
			if isLoop(position, Vector{x: 0, y: -1}, Vector{x: col, y: row}, grid) {
				sum++
			}
		}
	}
	return sum
}

func isLoop(position Vector, direction Vector, obstaclePosition Vector, grid []string) bool {
	var lastVisitedObstacles []Vector
	for {
		if position.x <= 0 || position.x >= len(grid[0])-1 || position.y <= 0 || position.y >= len(grid)-1 {
			return false
		}
		next := position.Add(direction)
		if grid[next.y][next.x] == '#' || (next.x == obstaclePosition.x && next.y == obstaclePosition.y) {
			lastVisitedObstacles = append(lastVisitedObstacles, next)
			if currentlyInLoop(lastVisitedObstacles) {
				break
			}
			direction.TurnRight()
			position = position.Add(direction)
			continue
		}
		position = next
	}
	return true
}

func currentlyInLoop(slice []Vector) bool {
	if len(slice) < 8 { //loop must have at least 4 obstacles for a loop
		return false
	}
	lastTwo := slice[len(slice)-2:]

	index := len(slice) - 3
	potentialLoop := false
	for index > 0 {
		if slice[index-1] == lastTwo[0] && slice[index] == lastTwo[1] {
			potentialLoop = true
			break
		}
		index--
	}
	if potentialLoop {
		length := len(slice) - index - 1
		if index-length+1 < 0 {
			return false
		}
		firstSlice := slice[index-length+1 : index+1]
		secondSlice := slice[index+1 : index+1+length]
		for i := 0; i < length; i++ {
			if firstSlice[i].x != secondSlice[i].x || firstSlice[i].y != secondSlice[i].y {
				return false
			}
		}
		return true
	}
	return false
}

func (v *Vector) TurnRight() {
	temp := v.x
	v.x = -v.y
	v.y = temp
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}
