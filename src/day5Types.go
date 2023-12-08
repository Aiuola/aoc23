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
	var mappedRanges, stillUnMappedRanges, unMappedRanges, ranges []*Range

	for _, entry := range m.entries {
		// For each entry, we need to know if it's no longer possible
		// to map values
		if unMappedRanges == nil {
			unMappedRanges, mappedRanges = entry.MapEntry(r)
		} else {
			for _, unMappedRange := range unMappedRanges {
				mappedRanges, stillUnMappedRanges = entry.MapEntry(unMappedRange)
				ranges = append(ranges, mappedRanges...)
			}
		}
		ranges = append(ranges, mappedRanges...)
	}

	if stillUnMappedRanges != nil {
		return ranges
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
		"shifting by %d for range %s",
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
		ret := r.SplitInSubRanges(e.mapRange)
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

func (r *Range) SplitInSubRanges(other *Range) []*Range {
	ranges := make([]*Range, 3)
	ranges[0] = NewRangeEnd(r.start, other.start)
	ranges[1] = other
	ranges[2] = NewRangeEnd(other.end, r.end)
	return ranges
}
