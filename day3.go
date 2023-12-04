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
	origX, origY := p.GetCoords()
	var x, y int
	bounds := make([]*Point, 0, 2*2)

	for i := -1; i < 2; i += 2 {
		y = origY + i
		for j := -1; j < 2; j += 2 {
			x = origX + i
			bounds = append(bounds, NewPoint(x, y))
		}
	}

	return NewArea(bounds[0], bounds[1], bounds[2], bounds[3])
}

func (p Point) GenerateAreaBetween(point *Point) *Area {
	if p.Equals(point) {
		return p.GenerateArea()
	}

	bounds := make([]*Point, 0, 2*2)
	smaller, bigger := p.CompareX(point)

	x, y := smaller.GetCoords()
	bounds = append(bounds, NewPoint(x-1, y-1))
	bounds = append(bounds, NewPoint(x-1, y+1))

	x, y = bigger.GetCoords()
	bounds = append(bounds, NewPoint(x+1, y-1))
	bounds = append(bounds, NewPoint(x+1, y+1))

	return NewArea(bounds[0], bounds[1], bounds[2], bounds[3])
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
	return fmt.Sprintf("{%d - %d}", p.x, p.y)
}

type PartNumber struct {
	area *Area
	val  int
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

// TODO: bro idk it's late
func (a Area) Intersects(other *Area) bool {
	if a.upperRight.x < other.lowerLeft.x || other.upperRight.x < a.lowerRight.x {
		return false
	}

	if a.upperRight.y < other.lowerLeft.y || other.upperRight.y < a.lowerLeft.y {
		return false
	}

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

	partNumbers, symbolAreas := parseSchematic(dat)

	for _, area := range symbolAreas {
		area.Print()
	}

	var match bool
	aggregator := 0

	for _, number := range partNumbers {
		match = false
		for _, symbolArea := range symbolAreas {
			if number.area.Intersects(symbolArea) {
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

func parseSchematic(dat []byte) ([]*PartNumber, []*Area) {
	partNumbers := make([]*PartNumber, 0)
	symbolAreas := make([]*Area, 0)

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
					NewPoint(x-nBytes, y).GenerateAreaBetween(NewPoint(x, y)),
					convertToNumber(bytes),
				),
			)
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
		symbolAreas = append(symbolAreas, NewPoint(x, y).GenerateArea())
		continue
	}

	return partNumbers, symbolAreas
}

func day3PartTwo(path string) int {
	return len(path)
}
