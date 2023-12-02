package main

import (
	"os"
	"strconv"
	"unicode"
)

const (
	base    = 10
	newLine = '\n'
)

func day1FirstPart(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	var accumulator = 0
	var firstDigit = -1
	var lastDigit = -1

	for _, b := range dat {
		if unicode.IsDigit(rune(b)) {
			if firstDigit == -1 {
				firstDigit, err = strconv.Atoi(string(b))
				check(err)
				continue
			}

			lastDigit, err = strconv.Atoi(string(b))
			check(err)
		}

		if b == newLine {
			if lastDigit == -1 {
				lastDigit = firstDigit
			}
			accumulator += firstDigit*base + lastDigit
			firstDigit = -1
			lastDigit = -1
		}
	}

	return accumulator
}

func day1SecondPart(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	var accumulator = 0
	var firstDigit = -1
	var lastDigit = -1

	for _, b := range dat {
		if unicode.IsDigit(rune(b)) {
			if firstDigit == -1 {
				firstDigit, err = strconv.Atoi(string(b))
				check(err)
				continue
			}

			lastDigit, err = strconv.Atoi(string(b))
			check(err)
		}

		if b == newLine {
			if lastDigit == -1 {
				lastDigit = firstDigit
			}
			accumulator += firstDigit*base + lastDigit
			firstDigit = -1
			lastDigit = -1
		}
	}

	return accumulator
}
