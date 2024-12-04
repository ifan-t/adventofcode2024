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
	var wordMatrix [][]rune
	for _, line := range lines {
		wordMatrix = append(wordMatrix, []rune(line))
	}
	part1Answer := part1(wordMatrix, "XMAS")
	fmt.Println("Answer to part 1:", part1Answer)
}

func part1(wordMatrix [][]rune, wordToMatch string) int {
	sum := 0
	for row, r := range wordMatrix {
		for col, _ := range r {
			if wordMatrix[row][col] == rune(wordToMatch[0]) {
				sum += checkEachDirectionForMatch(wordMatrix, wordToMatch, Vector{x: col, y: row})
			}
		}
	}
	return sum
}

func checkEachDirectionForMatch(wordMatrix [][]rune, wordToMatch string, startPoint Vector) int {
	return checkForMatch(wordMatrix, wordToMatch, startPoint, Vector{x: 0, y: 1}) +
		checkForMatch(wordMatrix, wordToMatch, startPoint, Vector{x: 0, y: -1}) +
		checkForMatch(wordMatrix, wordToMatch, startPoint, Vector{x: 1, y: 0}) +
		checkForMatch(wordMatrix, wordToMatch, startPoint, Vector{x: 1, y: 1}) +
		checkForMatch(wordMatrix, wordToMatch, startPoint, Vector{x: 1, y: -1}) +
		checkForMatch(wordMatrix, wordToMatch, startPoint, Vector{x: -1, y: 0}) +
		checkForMatch(wordMatrix, wordToMatch, startPoint, Vector{x: -1, y: 1}) +
		checkForMatch(wordMatrix, wordToMatch, startPoint, Vector{x: -1, y: -1})
}

func checkForMatch(wordMatrix [][]rune, wordToMatch string, startPoint Vector, direction Vector) int {
	index := 1
	var coordinate Vector = startPoint.Add(direction)
	for index < len(wordToMatch) {
		if isOutOfBounds(wordMatrix, coordinate) || wordMatrix[coordinate.y][coordinate.x] != rune(wordToMatch[index]) {
			return 0
		}
		index++
		coordinate = coordinate.Add(direction)
	}
	return 1
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func isOutOfBounds(wordMatrix [][]rune, coordinate Vector) bool {
	return coordinate.x < 0 || coordinate.x >= len(wordMatrix) || coordinate.y < 0 || coordinate.y >= len(wordMatrix[0])
}
