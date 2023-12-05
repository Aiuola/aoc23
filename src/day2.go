package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Cubes struct {
	red   int
	green int
	blue  int
}

var maxCubes = Cubes{
	red:   12,
	green: 13,
	blue:  14,
}

func NewCubes() *Cubes {
	return &Cubes{red: 0, green: 0, blue: 0}
}

func (cubes *Cubes) AddCube(size int, color string) {
	switch color {
	case "red":
		if cubes.red < size {
			cubes.red = size
		}
		break
	case "green":
		if cubes.green < size {
			cubes.green = size
		}
		break
	case "blue":
		if cubes.blue < size {
			cubes.blue = size
		}
		break
	default:
		panic(fmt.Sprintf("Unrecognized color %s\n", color))
	}
}

func (cubes Cubes) LowerThanMax() bool {
	return cubes.red <= maxCubes.red &&
		cubes.green <= maxCubes.green &&
		cubes.blue <= maxCubes.blue
}

func (cubes Cubes) Power() int {
	return cubes.red * cubes.blue * cubes.green
}

func (cubes Cubes) Print() string {
	return fmt.Sprintf("{ r - %d g - %d b - %d }", cubes.red, cubes.green, cubes.blue)
}

func day2partOne(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	var cubes *Cubes
	aggregator := 0
	var found bool

	for i, line := range lines {
		line, found = strings.CutPrefix(line, fmt.Sprintf("Game %d: ", i+1))

		if !found {
			panic(fmt.Sprintf("Couldn't remove prefix from %s", line))
		}

		cubes = cubesShownInGame(line)
		fmt.Printf("\nFrom line: %s\nI got cube %s\n", line, cubes.Print())
		if cubes.LowerThanMax() {
			fmt.Println("Which is valid!")
			aggregator = aggregator + i + 1
		}
	}

	return aggregator
}

func cubesShownInGame(line string) *Cubes {
	var numberString string = ""
	var parsedWord []int32 = nil
	var color string
	var shownCubes int
	var err error
	var cubes = NewCubes()

	for _, char := range line {
		if unicode.IsDigit(char) {
			numberString = fmt.Sprintf("%s%c", numberString, char)
			continue
		}

		if char == ' ' {
			if numberString != "" {
				shownCubes, err = strconv.Atoi(numberString)
				check(err)
				numberString = ""
			}
			continue
		}

		if char == ',' || char == ';' {
			color = string(parsedWord)
			parsedWord = make([]int32, 0)
			cubes.AddCube(shownCubes, color)
			continue
		}

		parsedWord = append(parsedWord, char)
	}

	color = string(parsedWord)
	cubes.AddCube(shownCubes, color)

	return cubes
}

func day2partTwo(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	lines := strings.Split(string(dat), "\n")
	var cubes *Cubes
	aggregator := 0
	var found bool

	for i, line := range lines {
		line, found = strings.CutPrefix(line, fmt.Sprintf("Game %d: ", i+1))
		if !found {
			panic(fmt.Sprintf("Couldn't remove prefix from %s", line))
		}

		cubes = cubesShownInGame(line)
		fmt.Printf("\nFrom line: %s\nI got cube %s\n", line, cubes.Print())

		aggregator = aggregator + cubes.Power()
	}

	return aggregator
}
