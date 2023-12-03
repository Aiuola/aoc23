package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Point struct {
	x, y int
}

func NewPoint(x int, y int) *Point {
	return &Point{x: x, y: y}
}

func (p Point) GetCoords() (int, int) {
	return p.x, p.y
}

func (p Point) GenerateNeighbours() [3 * 3]*Point {
	origX, origY := p.GetCoords()
	var x, y int
	neighbours := make([]*Point, 0, 3*3)

	for i := -1; i < 2; i++ {
		y = origY + i
		for j := -1; j < 2; j++ {
			x = origX + i
			neighbours = append(neighbours, NewPoint(x, y))
		}
	}
	return [9]*Point(neighbours)
}

func (p Point) Equals(point *Point) bool {
	return p.x == point.x &&
		p.y == point.y
}

type PartNumber struct {
	start  *Point
	length int
	val    int
}

func NewPartNumber(start *Point, length int, val int) *PartNumber {
	return &PartNumber{start: start, length: length, val: val}
}

type Symbol struct {
	affectedPoints [9]*Point
}

func NewSymbol(point *Point) *Symbol {
	return &Symbol{affectedPoints: point.GenerateNeighbours()}
}

func (s Symbol) AdjacentTo(number *PartNumber) bool {
	for _, point := range s.affectedPoints {
		if point.Equals(number.start) {
			return true
		}
	}
	return false
}

func convertToNumber(bytes []byte) int {
	val, err := strconv.Atoi(string(bytes[:]))
	check(err)
	return val
}

func day3PartOne(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	partNumbers, symbols := parseSchematic(dat)

	var match bool
	aggregator := 0

	for _, number := range partNumbers {
		match = false
		for _, symbol := range symbols {
			if symbol.AdjacentTo(number) {
				match = true
				break
			}
		}
		if match {
			aggregator += number.val
		}
	}

	return len(partNumbers) + len(symbols)
}

func parseSchematic(dat []byte) ([]*PartNumber, []*Symbol) {
	partNumbers := make([]*PartNumber, 0)
	symbols := make([]*Symbol, 0)

	bytes := make([]byte, 0)
	nBytes := 0

	y := 0

	for x, b := range dat {
		if unicode.IsDigit(rune(b)) {
			bytes = append(bytes, b)
			continue
		}

		nBytes = len(bytes)
		if nBytes != 0 {
			partNumbers = append(partNumbers,
				NewPartNumber(
					NewPoint(x, y),
					nBytes,
					convertToNumber(bytes),
				),
			)
			bytes = make([]byte, 0)
		}

		if b == '\n' {
			y++
			continue
		}

		if b == '.' {
			continue
		}
		fmt.Printf("Found symbol %s\n", string(b))
		symbols = append(symbols, NewSymbol(NewPoint(x, y)))
		continue
	}

	return partNumbers, symbols
}

func day3PartTwo(path string) int {
	return len(path)
}
