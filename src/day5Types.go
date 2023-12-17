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

// MapRange Objective = split the range in multiple sub-ranges without
// changing the total number of numbers covered
func (m Map) MapRange(r *Range) []*Range {
	fmt.Printf("\nMapping range %s\n", r)
	ranges := make([]*Range, 0)

	for i, entry := range m.entries {
		fmt.Printf("Entry:\n%s\n", entry.ToString())

		if !entry.mapRange.ContainsNumber(r.start) {
			continue
		}
		// Any part of the range before matches
		if r.start < entry.mapRange.start {
			orig, newR := r.SplitAt(r.start)
			r = newR
			ranges = append(ranges, orig)
		}

		for j := i; j < len(m.entries); j++ {
			// It's over
			if m.entries[j].mapRange.end >= r.end {
				lastRange := NewRangeEnd(
					max(m.entries[j].mapRange.start, r.start),
					r.end,
				)
				lastRange.Shift(m.entries[j].destination - m.entries[j].source)
				ranges = append(ranges, lastRange)
				// Since we have found the end, we can break
				break
			}

			// The end was not found
			mappedRange := NewRangeEnd(
				m.entries[j].mapRange.start,
				m.entries[j].mapRange.end,
			)
			mappedRange.Shift(entry.destination - entry.source)
			ranges = append(ranges, mappedRange)
		}
	}

	// Adding any reaming part of the range
	highestHigh := m.entries[len(m.entries)-1].mapRange.end
	if highestHigh < r.end {
		ranges = append(ranges, NewRangeEnd(
			highestHigh+1,
			r.end,
		))
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
		e.mapRange,
	)
}

func (e Entry) TryToMap(val int) (bool, int) {
	if e.mapRange.IsOutsideRange(val) {
		return false, -1
	}

	dif := val - e.mapRange.start

	return true, e.destination + dif
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

func (r *Range) String() string {
	if r.NumbersInRange() == 2 {
		return ""
	}
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

func (r *Range) SplitAt(number int) (*Range, *Range) {
	ranges := make([]*Range, 2)
	ranges[0] = NewRangeEnd(r.start, number)
	ranges[1] = NewRangeEnd(number+1, r.end)
	return ranges[0], ranges[1]
}

func (r *Range) ContainsNumber(number int) bool {
	return number >= r.start || number <= r.end
}

func (r *Range) ContainsNumberExclusive(n int) bool {
	return n > r.start && n < r.end
}
