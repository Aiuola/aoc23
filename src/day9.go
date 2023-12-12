package main

import (
	"os"
	"strconv"
)

func day9PartOne(path string) int {
	return coreDay9(path, false)
}

func coreDay9(path string, day2 bool) int {
	dat, err := os.ReadFile(path)
	check(err)

	histories := parseHistories(dat)
	return histories.ExtrapolateSequences(day2)
}

func parseHistories(dat []byte) Sequences {
	histories := make(Sequences, 0)

	currHistory := make([]int, 0)
	currNumber := make([]byte, 0)

	for _, b := range dat {
		if b == ' ' || b == '\n' {
			recording, err := strconv.Atoi(string(currNumber))
			check(err)
			currNumber = make([]byte, 0)

			currHistory = append(currHistory, recording)
			if b == '\n' {
				histories = append(histories, currHistory)
				currHistory = make([]int, 0)
			}
			continue
		}
		currNumber = append(currNumber, b)
	}

	return histories
}

func day9PartTwo(path string) int {
	return coreDay9(path, true)
}
