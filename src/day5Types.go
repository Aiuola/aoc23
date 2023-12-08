package main

import "fmt"

type Map struct {
	entries []*Entry
}

func (m Map) MapValue(val int) int {
	var possible bool
	var mappedVal int

	for _, entry := range m.entries {
		possible, mappedVal = entry.TryToMap(val)
		if possible {
			break
		}
	}

	if !possible {
		return val
	}

	return mappedVal
}

func NewMap(entries []*Entry) *Map {
	return &Map{entries: entries}
}

func (m Map) ToString() string {
	result := "Map containing entries:\n"
	for _, entry := range m.entries {
		result = fmt.Sprintf("%s%s\n", result, entry.ToString())
	}
	return result
}

type Entry struct {
	mapRange       *Range
	source, target int
}

func (e Entry) ToString() string {
	return fmt.Sprintf(
		"%d %d with range %d-%d",
		e.source,
		e.target,
		e.mapRange.start,
		e.mapRange.end,
	)
}

func (e Entry) TryToMap(val int) (bool, int) {
	if e.mapRange.IsOutsideRange(val) {
		return false, -1
	}

	dif := val - e.mapRange.start

	return true, e.target + dif
}

func NewEntry(mapRange int, source int, target int) *Entry {
	return &Entry{mapRange: NewRange(source, mapRange), source: source, target: target}
}

type Range struct {
	start, end int
}

func NewRange(start int, containedValues int) *Range {
	return &Range{start: start, end: start + containedValues - 1}
}

func (r Range) IsOutsideRange(val int) bool {
	return val < r.start || val > r.end
}
