package main

import (
	"os"
	"strings"
)

func day6PartOne(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")

	var raceTimes []int
	var recordLengths []int

	for i, line := range lines {
		line = line[strings.IndexByte(line, ':')+2:]
		line = strings.TrimSpace(line)

		if i == 0 {
			raceTimes = parseNumbers(line)
		} else {
			recordLengths = parseNumbers(line)
		}
	}

	aggregator := 1
	for i := 0; i < len(recordLengths); i++ {
		aggregator *= waysToBeatRecord(raceTimes[i], recordLengths[i])
	}

	return aggregator
}

func parseNumbers(line string) []int {
	ret := make([]int, 0)
	currNumber := make([]int32, 0)
	for _, char := range line {
		if char == ' ' {
			if len(currNumber) == 0 {
				continue
			}
			ret = append(ret, intArrayToNumber(currNumber))
			currNumber = make([]int32, 0)
			continue
		}
		currNumber = append(currNumber, char)
	}
	ret = append(ret, intArrayToNumber(currNumber))
	return ret
}

func waysToBeatRecord(time int, recordLength int) int {
	var attemptLength int
	cont := 0
	for i := 1; i < time; i++ {
		attemptLength = (time - i) * i
		if attemptLength > recordLength {
			cont++
		}
	}
	return cont
}

func day6PartTwo(path string) int {
	return len(path)
}
