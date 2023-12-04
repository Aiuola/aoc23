package main

import (
	"os"
	"unicode"
)

type Line struct {
	start *Point
	end   *Point
}

func NewHorizontalLine(start *Point, length int) *Line {
	return &Line{start: start, end: NewPoint(start.x+length, start.y)}
}

func (l Line) GetHorizon() int {
	return l.start.y
}

func (l Line) getVertices() (int, int) {
	return l.start.x, l.end.x
}

type SimplePart struct {
	line *Line
	val  int
}

func NewSimplePart(line *Line, val int) *SimplePart {
	return &SimplePart{line: line, val: val}
}

func day3PartTwo(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	simpleParts, gears := fetchGears(dat)

	aggregator := 0
	var match bool
	var gearRatio int

	for _, gear := range gears {
		match, gearRatio = validGear(gear, simpleParts)

		if match {
			aggregator += gearRatio
		}
	}

	return aggregator
}

func validGear(gear *Area, simpleParts []*SimplePart) (bool, int) {
	firstPart := -1
	secondPart := -1
	for _, part := range simpleParts {
		if gear.IntersectsHorizontally(part.line) {
			if firstPart == -1 {
				firstPart = part.val
				continue
			}
			if secondPart == -1 {
				secondPart = part.val
				continue
			}
			return false, -1
		}
	}

	if firstPart == -1 || secondPart == -1 {
		return false, -1
	}

	return true, firstPart * secondPart
}

func fetchGears(dat []byte) ([]*SimplePart, []*Area) {
	simpleParts := make([]*SimplePart, 0)
	gears := make([]*Area, 0)

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
			simpleParts = append(simpleParts,
				NewSimplePart(
					NewHorizontalLine(NewPoint(x-nBytes, y), nBytes-1),
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

		if b != '*' {
			continue
		}

		gears = append(gears, NewPoint(x, y).GenerateArea())
	}

	return simpleParts, gears
}
