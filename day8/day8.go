package main

import (
	"fmt"
	"math"

	helper "github.com/ifan-t/adventofcode2024"
)

type Vector struct {
	x int
	y int
}

func main() {
	lines, _ := helper.ParseInput("input.txt")
	noOfCols := len(lines[0]) - 1
	noOfRows := len(lines) - 1
	var strToVectorMap = make(map[string][]Vector)
	for col, line := range lines {
		for row, char := range line {
			if char == '.' || char == '#' {
				continue
			}
			str := string(char)
			_, exists := strToVectorMap[str]
			if exists {
				strToVectorMap[str] = append(strToVectorMap[str], Vector{x: col, y: row})
			} else {
				strToVectorMap[str] = []Vector{Vector{x: col, y: row}}
			}
		}
	}
	part1Answer := part1(noOfCols, noOfRows, strToVectorMap)
	fmt.Println("Answer to part 1 is: ", part1Answer)
	part2Answer := part2(noOfCols, noOfRows, strToVectorMap)
	fmt.Println("Answer to part 2 is: ", part2Answer)
}

func part1(noOfCols int, noOfRows int, strToVectorMap map[string][]Vector) int {
	sum := 0
	visited := make(map[Vector]bool)
	for _, vectors := range strToVectorMap {
		for i, vector1 := range vectors {
			for j, vector2 := range vectors {
				if j <= i {
					continue
				}
				difference1 := vector1.GetDifference(vector2)
				potentialAntinode1 := vector1.Add(difference1)
				if potentialAntinode1.x >= 0 && potentialAntinode1.x <= noOfCols && potentialAntinode1.y >= 0 && potentialAntinode1.y <= noOfRows {
					_, haveVisited := visited[potentialAntinode1]
					if !haveVisited {
						sum++
						visited[potentialAntinode1] = true
					}

				}
				difference2 := vector2.GetDifference(vector1)
				potentialAntinode2 := vector2.Add(difference2)
				if potentialAntinode2.x >= 0 && potentialAntinode2.x <= noOfCols && potentialAntinode2.y >= 0 && potentialAntinode2.y <= noOfRows {
					_, haveVisited2 := visited[potentialAntinode2]
					if !haveVisited2 {
						sum++
						visited[potentialAntinode2] = true
					}
				}

			}
		}
	}
	return sum
}

func part2(noOfCols int, noOfRows int, strToVectorMap map[string][]Vector) int {
	sum := 0
	visited := make(map[Vector]bool)
	for _, vectors := range strToVectorMap {
		for i, vector1 := range vectors {
			for j, vector2 := range vectors {
				if j <= i {
					continue
				}
				m, c := getLine(vector1, vector2)
				for x := 0; x <= noOfCols; x++ {
					y := (m * float64(x)) + c
					if !isWholeNumber(y) {
						continue
					}
					yInt := int(math.Round(y))
					potentialAntinode := Vector{x: x, y: yInt}
					if potentialAntinode.y >= 0 && potentialAntinode.y <= noOfRows {
						_, haveVisited := visited[potentialAntinode]
						if !haveVisited {
							sum++
							visited[potentialAntinode] = true
						}
					}
				}
			}
		}
	}
	return sum
}
func isWholeNumber(f float64) bool {
	return math.Abs(f-math.Round(f)) <= 0.001
}

func (v Vector) Add(other Vector) Vector {
	return Vector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vector) GetDifference(other Vector) Vector {
	return Vector{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func getLine(v1 Vector, v2 Vector) (float64, float64) {
	yChange := v1.y - v2.y
	xChange := v1.x - v2.x
	m := float64(yChange) / float64(xChange)
	c := float64(v1.y) - m*float64(v1.x)
	return m, c
}
