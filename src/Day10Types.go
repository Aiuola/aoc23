package main

import "fmt"

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

func (p Pipe) ExploreLoop(lines []string, indexes *Indexes) Loop {
	nextPipe := &p
	loop := make([]*Index, 0)

	for {
		indexes = nextPipe.NextIndexes(indexes)
		nextPipe = charToPipeMap[lines[indexes.i][indexes.j]]
		loop = append(loop, indexes.GetCurrIndex())

		if nextPipe.symbol == "S" {
			if len(loop)%2 == 1 {
				panic("Loop length was not dividable by two")
			}
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

func NewIndex(i int, j int) *Index {
	return &Index{i: i, j: j}
}

type Loop []*Index
