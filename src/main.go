package main

import (
	"flag"
	"fmt"
)

var validParts = []int{0, 1, 2}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isInArray(target int, arr []int) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}

func arrToString(arr []int) string {
	res := "{"
	for i, val := range arr {
		if i == len(arr)-1 {
			res = fmt.Sprintf("%s%d", res, val)
			continue
		}
		res = fmt.Sprintf("%s%d, ", res, val)
	}
	return fmt.Sprintf("%s}", res)
}

type Day struct {
	partOne func(string) int
	partTwo func(string) int
}

func main() {
	dirsPtr := flag.Bool("dirs", false, "Create directories and empty files for input")
	dayPtr := flag.Int("day", 0, "Day to execute")
	partPtr := flag.Int("part", 0, "What part of the day do you want to execute")
	testPtr := flag.Bool("test", false, "Should only tests be run")
	flag.Parse()

	if *dirsPtr {
		generateDirs()
	}

	if *dayPtr == 0 {
		fmt.Println("No day specified please provide the day you want to execute")
		return
	}

	if !isInArray(*partPtr, validParts) {
		panic(fmt.Sprintf("%d is not a valid part number, each day has only 2 parts"+
			"\nAvaiaible inputs are: %s", *partPtr, arrToString(validParts)))
	}

	inputs := generateInputPaths()

	fmt.Printf("Chosen day %d part %d, test only: %t\n", *dayPtr, *partPtr, *testPtr)

	var testResult int
	var mainResult int
	var day Day
	var firstPart = *partPtr == 0 || *partPtr == 1
	var secondPart = *partPtr == 0 || *partPtr == 2
	var dayInput = inputs[*dayPtr-1]
	switch *dayPtr {
	case 1:
		day = Day{
			partOne: func(s string) int {
				return day1PartOne(s)
			},
			partTwo: func(s string) int {
				return day1PartTwo(s)
			},
		}
		break
	case 2:
		day = Day{
			partOne: func(s string) int {
				return day2partOne(s)
			},
			partTwo: func(s string) int {
				return day2partTwo(s)
			},
		}
		break
	case 3:
		day = Day{
			partOne: func(s string) int {
				return day3PartOne(s)
			},
			partTwo: func(s string) int {
				return day3PartTwo(s)
			},
		}
		break
	case 4:
		day = Day{
			partOne: func(s string) int {
				return day4PartOne(s)
			},
			partTwo: func(s string) int {
				return day4PartTwo(s)
			},
		}
		break
	case 5:
		day = Day{
			partOne: func(s string) int {
				return day5PartOne(s)
			},
			partTwo: func(s string) int {
				return day5PartTwo(s)
			},
		}
		break
	default:
		panic(fmt.Sprintf("Unknown day number %d please provide a number within 0-25", *dayPtr))
	}

	if firstPart {
		testResult = day.partOne(dayInput.exampleOne)
		if !*testPtr {
			mainResult = day.partOne(dayInput.mainInput)
		}
		fmt.Printf("Test result: %d\nMain result: %d\n", testResult, mainResult)
	}
	if secondPart {
		testResult = day.partTwo(dayInput.exampleTwo)
		if !*testPtr {
			mainResult = day.partTwo(dayInput.mainInput)
		}
		fmt.Printf("Test result: %d\nMain result: %d\n", testResult, mainResult)
	}
}
