package main

import (
	"fmt"
	"sort"
)

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

// MapRange Objective = split the range in multiple arrays without
// changing the total number of numbers covered
func (m Map) MapRange(r *Range) []*Range {
	fmt.Printf("\nMapping range %s\n", r.ToString())
	ranges := make([]*Range, 0)
	match := false
	var rangeStart int

	for _, entry := range m.entries {
		fmt.Printf("Analyzing entry:\n%s\n", entry.ToString())

		// While we haven't found a start
		if entry.mapRange.end > r.start && !match {
			rangeStart = r.start
			match = true
		} else {
			if !match {
				continue
			}
			rangeStart = entry.mapRange.start
		}

		if entry.mapRange.end >= r.end {
			lastRange := NewRangeEnd(rangeStart, r.end)
			lastRange.Shift(entry.destination - entry.source)
			ranges = append(ranges, lastRange)
			// Since we have found the end, we can break
			break
		}

		// The end was not found
		ranges = append(ranges, NewRangeEnd(rangeStart, entry.mapRange.end))
	}

	// No matches found values remain unvaried
	if len(ranges) == 0 {
		// Partial match
		if match {
			highestEnd := m.entries[len(m.entries)-1].mapRange.end
			ranges = append(ranges, NewRangeEnd(rangeStart, highestEnd))
			ranges = append(ranges, NewRangeEnd(highestEnd+1, r.end))
		} else {
			ranges = append(ranges, r)
		}
	}

	return ranges
}

func NewMap(entries []*Entry) *Map {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].mapRange.start < entries[j].mapRange.end
	})
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
	mapRange            *Range
	source, destination int
}

func (e Entry) ToString() string {
	dif := e.destination - e.source
	return fmt.Sprintf(
		"%d for range %s",
		dif,
		e.mapRange.ToString(),
	)
}

func (e Entry) TryToMap(val int) (bool, int) {
	if e.mapRange.IsOutsideRange(val) {
		return false, -1
	}

	dif := val - e.mapRange.start

	return true, e.destination + dif
}

func (e Entry) MapEntry(r *Range) ([]*Range, []*Range) {
	// If we don't share any numbers, then just return the original range
	if !e.mapRange.SharesAnyNumbers(r) {
		ret := make([]*Range, 1)
		ret = append(ret, r)
		return make([]*Range, 0), ret
	}

	// If the mapping range is a superset, then map everything
	if e.mapRange.IsASuperSet(r) {
		ret := make([]*Range, 1)
		r.Shift(e.source - e.destination)
		ret = append(ret, r)
		return make([]*Range, 0), ret
	}

	// If the range to be mapped is a superset of the mapping range,
	// then 3 new ranges are going to be created
	// before mapping -> mapped values -> after mapping
	if r.IsASuperSet(e.mapRange) {
		ret := r.SplitWith(e.mapRange)
		ret[1].Shift(e.source - e.destination)
		return make([]*Range, 0), ret
	}

	panic("Bro...")

	return make([]*Range, 0), make([]*Range, 0)
}

func NewEntry(mapRange int, source int, target int) *Entry {
	return &Entry{mapRange: NewRange(source, mapRange), source: source, destination: target}
}

type Range struct {
	start, end int
}

func NewRange(start int, containedValues int) *Range {
	return &Range{start: start, end: start + containedValues - 1}
}

func NewRangeEnd(start int, end int) *Range {
	return &Range{start: start, end: end}
}

func (r *Range) NumbersInRange() int {
	return r.end - r.start + 1
}

func (r *Range) IsOutsideRange(val int) bool {
	return val < r.start || val > r.end
}

func (r *Range) SharesAnyNumbers(other *Range) bool {
	return (r.start <= other.start && r.end > other.start) ||
		(r.end >= other.end && r.start < other.end)
}

func (r *Range) IsASuperSet(other *Range) bool {
	return r.start <= other.start && r.end >= other.end
}

func (r *Range) ToString() string {
	return fmt.Sprintf("%d-%d", r.start, r.end)
}

func (r *Range) Shift(amount int) {
	r.start += amount
	r.end += amount
}

func (r *Range) SplitWith(other *Range) []*Range {
	ranges := make([]*Range, 3)
	ranges[0] = NewRangeEnd(r.start, other.start)
	ranges[1] = other
	ranges[2] = NewRangeEnd(other.end, r.end)
	return ranges
}

func (r *Range) SplitAt(number int) []*Range {
	ranges := make([]*Range, 3)
	ranges[0] = NewRangeEnd(r.start, number)
	ranges[1] = NewRangeEnd(number+1, r.end)
	return ranges
}

func (r *Range) ContainsNumber(number int) bool {
	return number >= r.start || number <= r.end
}
