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

func (p Point) GenerateArea() *Area {
	x, y := p.GetCoords()

	upperLeft := NewPoint(x-1, y-1)
	lowerLeft := NewPoint(x-1, y+1)
	upperRight := NewPoint(x+1, y-1)
	lowerRight := NewPoint(x+1, y+1)

	return NewArea(upperRight, upperLeft, lowerRight, lowerLeft)
}

func (p Point) GenerateAreaBetween(point *Point) *Area {
	if p.Equals(point) {
		return p.GenerateArea()
	}

	smaller, bigger := p.CompareX(point)

	x, y := smaller.GetCoords()
	upperLeft := NewPoint(x-1, y-1)
	lowerLeft := NewPoint(x-1, y+1)

	x, y = bigger.GetCoords()
	upperRight := NewPoint(x+1, y-1)
	lowerRight := NewPoint(x+1, y+1)

	return NewArea(upperRight, upperLeft, lowerRight, lowerLeft)
}

func (p Point) Equals(point *Point) bool {
	return p.x == point.x &&
		p.y == point.y
}

func (p Point) CompareX(point *Point) (*Point, *Point) {
	if p.x < point.x {
		return &p, point
	}
	return point, &p
}

func (p Point) ToString() string {
	return fmt.Sprintf("{x:%d y:%d}", p.x, p.y)
}

type PartNumber struct {
	area *Area
	val  int
}

func (n PartNumber) Print() {
	fmt.Printf("\nFor %d:", n.val)
	n.area.Print()
}

func NewPartNumber(area *Area, val int) *PartNumber {
	return &PartNumber{area: area, val: val}
}

type Area struct {
	upperRight, upperLeft, lowerRight, lowerLeft *Point
}

func NewArea(upperRight *Point, upperLeft *Point, lowerRight *Point, lowerLeft *Point) *Area {
	return &Area{upperRight: upperRight, upperLeft: upperLeft, lowerRight: lowerRight, lowerLeft: lowerLeft}
}

func (a Area) Contains(point *Point) bool {
	x, y := point.GetCoords()

	// Any left point
	if x < a.upperLeft.x {
		return false
	}
	// Any right point
	if x > a.upperRight.x {
		return false
	}
	// X is in range!

	// Any upper point
	if y < a.upperRight.y {
		return false
	}
	// Any lower point
	if y > a.lowerRight.y {
		return false
	}
	// Y is in range!

	return true
}

func (a Area) Print() {
	fmt.Printf("\n%s - %s\n%s - %s\n",
		a.upperLeft.ToString(), a.upperRight.ToString(),
		a.lowerLeft.ToString(), a.lowerRight.ToString())
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
			if number.area.Contains(symbol) {
				match = true
				break
			}
		}
		if match {
			aggregator += number.val
		}
	}

	return aggregator
}

func parseSchematic(dat []byte) ([]*PartNumber, []*Point) {
	partNumbers := make([]*PartNumber, 0)
	symbols := make([]*Point, 0)

	bytes := make([]byte, 0)
	nBytes := 0

	y := 0
	x := -1

	for _, b := range dat {
		x++
		if unicode.IsDigit(rune(b)) {
			bytes = append(bytes, b)
			continue
		}

		nBytes = len(bytes)
		if nBytes != 0 {
			partNumbers = append(partNumbers,
				NewPartNumber(
					NewPoint(x-nBytes, y).GenerateAreaBetween(NewPoint(x-1, y)),
					convertToNumber(bytes),
				),
			)
			if nBytes == 2 {
				partNumbers[len(partNumbers)-1].Print()
			}
			bytes = make([]byte, 0)
		}

		if b == '\n' {
			x = -1
			y++
			continue
		}

		if b == '.' {
			continue
		}

		symbols = append(symbols, NewPoint(x, y))
	}

	return partNumbers, symbols
}

func day3PartTwo(path string) int {
	return len(path)
}
