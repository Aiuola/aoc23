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

func printCubes(cubes *Cubes) string {
	return fmt.Sprintf("{ r - %d g - %d b - %d }", cubes.red, cubes.green, cubes.blue)
}

func NewCubes() *Cubes {
	return &Cubes{red: 0, green: 0, blue: 0}
}

func addCube(cubes *Cubes, size int, color string) {
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

func nCubesPossible(gameCubes *Cubes) bool {
	return gameCubes.red <= maxCubes.red &&
		gameCubes.green <= maxCubes.green &&
		gameCubes.blue <= maxCubes.blue
}

var maxCubes = Cubes{
	red:   12,
	green: 13,
	blue:  14,
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

		cubes = cubesShownInGame(line, i)
		fmt.Printf("\nFrom line: %s\nI got cube %s\n", line, printCubes(cubes))
		if nCubesPossible(cubes) {
			fmt.Println("Which is valid!")
			aggregator = aggregator + i + 1
		}
	}

	return aggregator
}

func cubesShownInGame(line string, i int) *Cubes {
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
			addCube(cubes, shownCubes, color)
			continue
		}

		parsedWord = append(parsedWord, char)
	}

	color = string(parsedWord)
	addCube(cubes, shownCubes, color)

	return cubes
}

func day2partTwo(path string) int {
	return len(path)
}
