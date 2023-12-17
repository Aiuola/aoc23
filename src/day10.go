package main

import (
	"os"
	"strings"
)

func day10PartOne(path string) int {
	return day10Core(path)
}

func day10Core(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	var left uint8
	initPipeMap()
	for i, line := range lines {
		for j, char := range line {
			if char != 'S' {
				continue
			}
			if j == 0 {
				left = '.'
			} else {
				left = line[j-1]
			}
			startingPipe := determineStartingSymbol(
				// North
				lines[i-1][j],
				// South
				lines[i+1][j],
				left,
				// Right
				line[j+1],
			)
			return startingPipe.LoopSize(lines,
				NewIndexes(i, j, -1, -1))
		}
	}
	panic("Starting pipe not found")
}

func parseLoop(lines []string) Loop {
	var left uint8
	initPipeMap()
	for i, line := range lines {
		for j, char := range line {
			if char != 'S' {
				continue
			}
			if j == 0 {
				left = '.'
			} else {
				left = line[j-1]
			}
			startingPipe := determineStartingSymbol(
				// North
				lines[i-1][j],
				// South
				lines[i+1][j],
				left,
				// Right
				line[j+1],
			)
			return startingPipe.ExploreLoop(lines,
				NewIndexes(i, j, -1, -1))
		}
	}
	panic("Starting pipe not found")
}

func day10PartTwo(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")

	return parseLoop(lines).GenerateRanges().TilesEnclosed(lines)
}
