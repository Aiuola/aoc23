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

func printArray(arr []int) {
	res := "{ "
	for _, val := range arr {
		res = fmt.Sprintf("%s, %d", res, val)
	}
	res = fmt.Sprintf("%s }", res)
}

func main() {
	allPtr := flag.Bool("all", false, "Run all days")
	dirsPtr := flag.Bool("dirs", false, "Create directories and empty files for input")
	dayPtr := flag.Int("day", 0, "Day to execute")
	partPtr := flag.Int("part", 0, "What part of the day do you want to execute")
	flag.Parse()

	if *dirsPtr {
		generateDirs()
	}

	if *dayPtr == 0 {
		fmt.Println("No day specified please provide the day you want to execute")
		return
	}

	if !isInArray(*partPtr, validParts) {
		panic(fmt.Sprintf("%d is not in the valid parts array", *partPtr))
	}

	paths := generateInputPaths()

	switch *partPtr {
	case 0:

	}

	fmt.Printf("Chosen day %d, run all: %t\n", *dayPtr, *allPtr)

	var testResult int
	var mainResult int
	switch *dayPtr {
	case 1:
		if *partPtr == 0 || *partPtr == 1 {
			testResult = day1FirstPart(paths[*dayPtr-1])
			mainResult = day1FirstPart(paths[*dayPtr])
		}
		if *partPtr == 0 || *partPtr == 2 {
			testResult = day1FirstPart(paths[*dayPtr-1])
			mainResult = day1FirstPart(paths[*dayPtr])
		}
	}

	fmt.Printf("Test result: %d\nMain result: %d\n", testResult, mainResult)
}
