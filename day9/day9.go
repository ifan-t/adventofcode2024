package main

import (
	"container/list"
	"fmt"
	"strconv"

	helper "github.com/ifan-t/adventofcode2024"
)

func main() {
	lines, _ := helper.ParseInput("input.txt")
	part1Answer := part1(lines[0])
	fmt.Println("Answer to part 1 is: ", part1Answer)
}

func part1(line string) int {
	queue := list.New()
	var filesystemArr []string
	for index, char := range line {
		if index%2 == 0 {
			fileIdStr := strconv.Itoa(index / 2)
			noOfRepetitions, _ := strconv.Atoi(string(char))
			for range noOfRepetitions {
				filesystemArr = append(filesystemArr, fileIdStr)
			}
		} else {
			noOfRepetitions, _ := strconv.Atoi(string(char))
			queue.PushBack(noOfRepetitions)
			for range noOfRepetitions {
				filesystemArr = append(filesystemArr, ".")
			}
		}
	}
	// filesystem := []rune(filesystemStr)
	for i := len(filesystemArr) - 1; i >= 0; i-- {
		if filesystemArr[i] == "." {
			continue
		}
		digitsStr := filesystemArr[i]
		for index, str := range filesystemArr {
			if index >= i {
				break
			}
			if str == "." {
				filesystemArr[index] = digitsStr
				break
			}
		}
	}
	sum := 0
	for i, id := range filesystemArr {
		if id == "." {
			break
		}
		idValue, _ := strconv.Atoi(id)
		sum += i * idValue
	}
	return sum
}
