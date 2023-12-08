package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func day5PartOne(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	seeds := parseSeeds(lines[0])
	//fmt.Printf("\nInitial seeds: %s\n", arrToString(seeds))
	maps := parseMaps(lines[3:])

	for _, m := range maps {
		for i := 0; i < len(seeds); i++ {
			seeds[i] = m.MapValue(seeds[i])
		}
		//fmt.Printf("%s\n%s\n", m.ToString(), arrToString(seeds))
	}

	lowest := seeds[0]
	for i := 1; i < len(seeds); i++ {
		if seeds[i] < lowest {
			lowest = seeds[i]
		}
	}

	return lowest
}

func parseSeeds(line string) []int {
	line = line[strings.IndexByte(line, ':')+2:]
	currNumber := make([]int32, 0)

	seeds := make([]int, 0)

	for _, char := range line {
		if char == ' ' {
			seeds = append(seeds, intArrayToNumber(currNumber))
			currNumber = make([]int32, 0)
			continue
		}

		currNumber = append(currNumber, char)
	}

	return append(seeds, intArrayToNumber(currNumber))
}

func parseMaps(lines []string) []*Map {
	maps := make([]*Map, 0)
	entries := make([]*Entry, 0)
	for _, line := range lines {
		if len(line) == 0 {
			maps = append(maps, NewMap(entries))
			entries = make([]*Entry, 0)
			continue
		}
		if !unicode.IsDigit(rune(line[0])) {
			continue
		}
		entries = append(entries, parseEntry(line))
	}
	maps = append(maps, NewMap(entries))

	return maps
}

func parseEntry(line string) *Entry {
	currNumber := make([]int32, 0)
	values := make([]int, 0)

	for _, char := range line {
		if char == ' ' {
			values = append(values, intArrayToNumber(currNumber))
			currNumber = make([]int32, 0)
			continue
		}
		currNumber = append(currNumber, char)
	}
	values = append(values, intArrayToNumber(currNumber))

	if len(values) != 3 {
		panic(fmt.Sprintf("More than 3 entries... for line %s", line))
	}

	return NewEntry(values[2], values[1], values[0])
}

func day5PartTwo(path string) int {
	return len(path)
}
