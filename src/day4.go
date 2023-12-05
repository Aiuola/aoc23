package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	two = 2
)

func day4PartOne(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	points := 0

	for i, line := range lines {
		line := line[strings.IndexByte(line, ':')+2:]
		points += parseLine(line)
		fmt.Printf("\nPoints after line %d - %d\n\n", i+1, points)
	}
	return points
}

func intArrayToNumber(arr []int32) int {
	kek, err := strconv.Atoi(string(arr))
	check(err)
	return kek
}

func parseLine(line string) int {
	currNumber := make([]int32, 0)
	var parsedNumber int
	foundPipe := false

	winningNumbers := NewSortedArray()

	points, cardValue := 0, 1

	for _, char := range line {
		if char == '|' {
			foundPipe = true
			continue
		}

		if char != ' ' {
			currNumber = append(currNumber, char)
			continue
		}

		if len(currNumber) != 0 {
			parsedNumber = intArrayToNumber(currNumber)
			currNumber = make([]int32, 0)

			if !foundPipe {
				winningNumbers.Insert(parsedNumber)
				continue
			}

			if !winningNumbers.Contains(parsedNumber) {
				continue
			}

			fmt.Printf("Match found! %d is in %s \n", parsedNumber, winningNumbers.ToString())
			points = cardValue
			cardValue *= two
		}
	}

	if len(currNumber) == 0 {
		panic("Line didn't end with number")
	}

	parsedNumber = intArrayToNumber(currNumber)
	if winningNumbers.Contains(parsedNumber) {
		fmt.Printf("Match found! %d is in %s \n", parsedNumber, winningNumbers.ToString())
		points = cardValue
		cardValue *= two
	}

	return points
}

func day4PartTwo(path string) int {
	return len(path)
}
