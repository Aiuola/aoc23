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

var stringToInt = map[string]int{}

func initMap() {
	stringToInt = make(map[string]int)

	stringToInt["one"] = 1
	stringToInt["two"] = 2
	stringToInt["three"] = 3
	stringToInt["four"] = 4
	stringToInt["five"] = 5
	stringToInt["six"] = 6
	stringToInt["seven"] = 7
	stringToInt["eight"] = 8
	stringToInt["nine"] = 9
}

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

func anyValidConversion(byteMatrix [][]byte) (bool, int) {
	var word string
	var val int
	var isPresent bool
	for _, bytes := range byteMatrix {
		if len(bytes) > 5 {
			continue
		}
		word = string(bytes)
		val, isPresent = stringToInt[word]
		if isPresent {
			return isPresent, val
		}
	}
	return false, -1
}

func updateDigits(words [][]byte, firstDigit int, lastDigit int, newDigit int) (int, int) {
	words = make([][]byte, 0)
	if firstDigit == -1 {
		return newDigit, lastDigit
	}

	return firstDigit, newDigit
}

func day1SecondPart(path string) int {
	dat, err := os.ReadFile(path)
	check(err)
	initMap()

	var accumulator = 0
	var firstDigit = -1
	var lastDigit = -1
	var val int
	var words = make([][]byte, 0)

	var validConversion bool
	var convertedValue int

	for _, b := range dat {
		if unicode.IsDigit(rune(b)) {
			val, err = strconv.Atoi(string(b))
			check(err)
			firstDigit, lastDigit = updateDigits(words, firstDigit, lastDigit, val)
			continue
		}

		words = append(words, make([]byte, 0))
		for i := 0; i < len(words); i++ {
			words[i] = append(words[i], b)
		}
		validConversion, convertedValue = anyValidConversion(words)
		if validConversion {
			firstDigit, lastDigit = updateDigits(words, firstDigit, lastDigit, convertedValue)
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
