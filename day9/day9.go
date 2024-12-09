package main

import (
	"fmt"
	"strconv"

	helper "github.com/ifan-t/adventofcode2024"
)

func main() {
	lines, _ := helper.ParseInput("input.txt")
	fileSystemArr := GetFileSystem(lines[0])

	part1Answer := part1(fileSystemArr)
	fmt.Println("Answer to part 1 is: ", part1Answer)
	filesystem := make([]string, len(fileSystemArr))
	copy(filesystem, fileSystemArr)
	part2Answer := part2(filesystem)
	fmt.Println("Answer to part 2 is: ", part2Answer)
}

func part1(fileSystem []string) int {

	for i := len(fileSystem) - 1; i >= 0; i-- {
		if fileSystem[i] == "." {
			continue
		}
		digitsStr := fileSystem[i]
		for index, str := range fileSystem {
			if index >= i {
				break
			}
			if str == "." {
				fileSystem[index] = digitsStr
				break
			}
		}
	}
	sum := 0
	for i, id := range fileSystem {
		if id == "." {
			break
		}
		idValue, _ := strconv.Atoi(id)
		sum += i * idValue
	}
	return sum
}

func part2(fileSystem []string) int {

	for i := len(fileSystem) - 1; i >= 0; i-- {
		if fileSystem[i] == "." {
			continue
		}
		digitsStr := fileSystem[i]
		fileBlockLength := 1
		for j := i - 1; j >= 0; j-- {
			if fileSystem[j] != digitsStr {
				break
			}
			fileBlockLength++
		}

		for index, str := range fileSystem {
			if index >= i {
				break
			}
			if str == "." && EnoughSpace(fileBlockLength, index, fileSystem) {
				for k := index; k < index+fileBlockLength; k++ {
					fileSystem[k] = digitsStr
				}
				for l := i; l > i-fileBlockLength; l-- {
					fileSystem[l] = "."
				}
				break
			}
		}
		i = i - fileBlockLength + 1
	}
	sum := 0
	for i, id := range fileSystem {
		if id == "." {
			continue
		}
		idValue, _ := strconv.Atoi(id)
		sum += i * idValue
	}
	return sum
}

func EnoughSpace(fileBlockLength int, currentIndex int, fileSystem []string) bool {
	for i := currentIndex; i < currentIndex+fileBlockLength; i++ {
		if fileSystem[i] != "." {
			return false
		}
	}
	return true
}

func GetFileSystem(line string) []string {
	var filesystem []string
	for index, char := range line {
		if index%2 == 0 {
			fileIdStr := strconv.Itoa(index / 2)
			noOfRepetitions, _ := strconv.Atoi(string(char))
			for range noOfRepetitions {
				filesystem = append(filesystem, fileIdStr)
			}
		} else {
			noOfRepetitions, _ := strconv.Atoi(string(char))
			for range noOfRepetitions {
				filesystem = append(filesystem, ".")
			}
		}
	}
	return filesystem
}
