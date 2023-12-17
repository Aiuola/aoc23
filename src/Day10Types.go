package main

import (
	"fmt"
	"sort"
)

type Pipe struct {
	connNorth, connSouth, connLeft, connRight bool
	symbol                                    string
}

func (p Pipe) String() string {
	return p.symbol
}

func NewPipe(connNorth bool, connSouth bool, connLeft bool, connRight bool, symbol string) *Pipe {
	return &Pipe{
		connNorth: connNorth,
		connSouth: connSouth,
		connLeft:  connLeft,
		connRight: connRight,
		symbol:    symbol,
	}
}

var charToPipeMap map[uint8]*Pipe

func initPipeMap() {
	charToPipeMap = make(map[uint8]*Pipe)

	charToPipeMap['|'] = &Pipe{connSouth: true, connNorth: true, symbol: "|"}
	charToPipeMap['-'] = &Pipe{connLeft: true, connRight: true, symbol: "-"}
	charToPipeMap['L'] = &Pipe{connNorth: true, connRight: true, symbol: "L"}
	charToPipeMap['J'] = &Pipe{connNorth: true, connLeft: true, symbol: "J"}
	charToPipeMap['7'] = &Pipe{connSouth: true, connLeft: true, symbol: "7"}
	charToPipeMap['F'] = &Pipe{connSouth: true, connRight: true, symbol: "F"}
	charToPipeMap['.'] = &Pipe{symbol: "."}
	charToPipeMap['S'] = &Pipe{symbol: "S"}
}

func determineStartingSymbol(north uint8, south uint8, left uint8, right uint8) *Pipe {
	n, s, l, r :=
		charToPipeMap[north].connSouth,
		charToPipeMap[south].connNorth,
		charToPipeMap[left].connRight,
		charToPipeMap[right].connLeft

	matches := 0
	directions := [...]bool{n, s, r, l}
	for _, direction := range directions {
		if direction {
			matches++
		}
	}
	if matches != 2 {
		panic(fmt.Sprintf("There should only be two matches but I found %d", matches))
	}

	return NewPipe(n, s, l, r, "S")
}

func (p Pipe) NextIndexes(indexes *Indexes) *Indexes {
	if p.connRight && indexes.j+1 != indexes.prevJ {
		return indexes.Move(0, 1)
	}
	if p.connLeft && indexes.j-1 != indexes.prevJ {
		return indexes.Move(0, -1)
	}
	if p.connSouth && indexes.i+1 != indexes.prevI {
		return indexes.Move(1, 0)
	}
	if p.connNorth && indexes.i-1 != indexes.prevI {
		return indexes.Move(-1, 0)
	}

	panic("Wait what...")
}

func (p Pipe) LoopSize(lines []string, indexes *Indexes) int {
	nextPipe := &p
	loop := make(map[int][]int)
	cont := 0

	for {
		indexes = nextPipe.NextIndexes(indexes)
		nextPipe = charToPipeMap[lines[indexes.i][indexes.j]]
		cont++

		if nextPipe.symbol == "S" {
			if len(loop)%2 == 1 {
				panic("Loop length was not dividable by two")
			}
			return cont / 2
		}
	}
}

func (p Pipe) ExploreLoop(lines []string, indexes *Indexes) Loop {
	nextPipe := &p
	loop := make(map[int][]int)

	for {
		indexes = nextPipe.NextIndexes(indexes)
		nextPipe = charToPipeMap[lines[indexes.i][indexes.j]]
		loop[indexes.i] = append(loop[indexes.i], indexes.j)

		if nextPipe.symbol == "S" {
			return loop
		}
	}
}

type Indexes struct {
	i, j         int
	prevI, prevJ int
}

func NewIndexes(i int, j int, prevI int, prevJ int) *Indexes {
	return &Indexes{i: i, j: j, prevI: prevI, prevJ: prevJ}
}

func (i Indexes) Move(iAmount int, jAmount int) *Indexes {
	return NewIndexes(
		i.i+iAmount, i.j+jAmount,
		i.i, i.j,
	)
}

func (i Indexes) GetCurrIndex() *Index {
	return NewIndex(i.i, i.j)
}

type Index struct {
	i, j int
}

func (i Index) String() string {
	return fmt.Sprintf("%d, %d", i.i, i.j)
}

func NewIndex(i int, j int) *Index {
	return &Index{i: i, j: j}
}

type Loop map[int][]int

func (l Loop) String() string {
	keys := make([]int, len(l))
	for k := range l {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	res := "Loop {"
	for key := range keys {
		if len(l[key]) != 0 {
			res = fmt.Sprintf("%s\n\t%d -> %v", res, key, l[key])
		}
	}
	return res + "\n}"
}

func (l Loop) Sort() {
	for i, _ := range l {
		sort.Ints(l[i])
	}
}

func (l Loop) GenerateRanges() LoopRanges {
	l.Sort()
	loopRanges := make(map[int][]*Range)

	for key, values := range l {
		loopRanges[key] = GenerateRanges(values)
	}
	return loopRanges
}

type LoopRanges map[int][]*Range

func GenerateRanges(values []int) []*Range {
	items := len(values)
	if items%2 == 1 {
		items--
	}
	ranges := make([]*Range, items/2)

	for i := 0; i < items; i += 2 {
		ranges[i/2] = NewRangeEnd(values[i], values[i+1])
	}

	return ranges
}

func (lr LoopRanges) TilesEnclosed(lines []string) int {
	fmt.Println(lr)
	matches := 0

	for i, line := range lines {
		for j, char := range line {
			if char != '.' {
				continue
			}
			for _, loopRange := range lr[i] {
				if loopRange.ContainsNumberExclusive(j) {
					fmt.Printf("Match at %d - %d\n", i, j)
					matches++
					break
				}
			}
		}
	}

	return matches
}

func (lr LoopRanges) String() string {
	keys := make([]int, len(lr))
	for k := range lr {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	res := "Loop-ranges {"
	for key := range keys {
		if len(lr[key]) != 0 {
			res = fmt.Sprintf("%s\n\t%d -> %s", res, key, lr[key])
		}
	}
	return res + "\n}"
}
