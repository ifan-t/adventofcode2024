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
