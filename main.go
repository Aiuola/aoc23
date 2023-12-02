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

func stringRepresentation(arr []int) string {
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
			"\nAvaiaible inputs are: %s", *partPtr, stringRepresentation(validParts)))
	}

	inputs := generateInputPaths()

	fmt.Printf("Chosen day %d part %d, debug: %t\n", *dayPtr, *partPtr, *testPtr)

	var testResult int
	var mainResult int
	var firstPart = *partPtr == 0 || *partPtr == 1
	var secondPart = *partPtr == 0 || *partPtr == 2
	switch *dayPtr {
	case 1:
		var dayInput = inputs[*dayPtr-1]
		if firstPart {
			testResult = day1FirstPart(dayInput.exampleOne)
			if *testPtr {
				break
			}
			mainResult = day1FirstPart(dayInput.mainInput)
		}
		if secondPart {
			testResult = day1SecondPart(dayInput.exampleTwo)
			if *testPtr {
				break
			}
			mainResult = day1SecondPart(dayInput.mainInput)
		}
	}

	fmt.Printf("Test result: %d\nMain result: %d\n", testResult, mainResult)
}
