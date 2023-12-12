package main

import "fmt"

type Sequences []Sequence

func (s Sequences) String() string {
	ret := "Sequences:"
	for _, seq := range s {
		ret = fmt.Sprintf("%s\n%v", ret, seq)
	}
	return ret
}

func (s Sequences) ExtrapolateSequences(day2 bool) int {
	sum := 0
	var prediction int
	for _, sequence := range s {
		if day2 {
			prediction = sequence.DiscoverHistory()
		} else {
			prediction = sequence.ExtrapolatePrediction()
		}
		sum += prediction
	}

	return sum
}

type Sequence []int

func (s Sequence) String() string {
	return fmt.Sprintf("%v", arrToString(s))
}

// GenerateSequences Only for visualization/debug
func (s Sequence) GenerateSequences() Sequences {
	a, b := s.IsMadeOfAllZeroes()
	if a {
		return b
	}
	childSequence := make(Sequence, len(s)-1)
	var prev int

	for i, item := range s {
		if i == 0 {
			prev = item
			continue
		}
		childSequence[i-1] = item - prev
		prev = item
	}

	sequences := make(Sequences, 1)
	sequences[0] = s
	sequences = append(sequences, childSequence.GenerateSequences()...)

	return sequences
}

func (s Sequence) ExtrapolatePrediction() int {
	a, _ := s.IsMadeOfAllZeroes()
	if a {
		return 0
	}

	childSequence := make(Sequence, len(s)-1)
	var prev int

	for i, item := range s {
		if i == 0 {
			prev = item
			continue
		}
		childSequence[i-1] = item - prev
		prev = item
	}

	return s[len(s)-1] + childSequence.ExtrapolatePrediction()
}

func (s Sequence) IsMadeOfAllZeroes() (bool, Sequences) {
	match := false
	for _, item := range s {
		if item != 0 {
			match = true
			break
		}
	}
	if !match {
		ret := make(Sequences, 1)
		ret[0] = s
		return true, ret
	}

	return false, nil
}

func (s Sequence) DiscoverHistory() int {
	a, _ := s.IsMadeOfAllZeroes()
	if a {
		return 0
	}

	childSequence := make(Sequence, len(s)-1)
	var prev int

	for i, item := range s {
		if i == 0 {
			prev = item
			continue
		}
		childSequence[i-1] = item - prev
		prev = item
	}

	return s[0] - childSequence.DiscoverHistory()
}
