package main

type SortedArray struct {
	elements []int
}

func NewSortedArray() *SortedArray {
	return &SortedArray{elements: make([]int, 0)}
}

func (s *SortedArray) Insert(val int) {
	size := len(s.elements)
	if size == 0 {
		s.elements = append(s.elements, val)
		return
	}
	for i, elem := range s.elements {
		if val > elem {
			continue
		}

		s.insert(val, i)
		return
	}
	s.elements = append(s.elements, val)
	return
}

func (s *SortedArray) insert(val int, index int) {
	s.elements = append(s.elements, 0)
	copy(s.elements[index+1:], s.elements[index:])
	s.elements[index] = val
}

func (s *SortedArray) Contains(val int) bool {
	low, high := 0, len(s.elements)-1

	if high == 0 {
		return false
	}
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if val == s.elements[mid] {
			return true
		}

		if val > s.elements[mid] {
			low = mid + 1
			continue
		}

		high = mid - 1
	}

	return false
}

func (s *SortedArray) ToString() any {
	return arrToString(s.elements)
}
