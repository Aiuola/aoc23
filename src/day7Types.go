package main

import "fmt"

type Hand struct {
	cards []int
	bid   int
	hType int
}

func NewHand(cards []int, bid int) *Hand {
	return &Hand{cards: cards, bid: bid, hType: determineHandType(cards)}
}

func determineHandType(cards []int) int {
	buckets := make(map[int]int)

	for _, card := range cards {
		buckets[card]++
	}

	size := len(buckets)

	var val int
	switch size {
	case 1:
		// Five of a kind
		val = 7
	case 2:
		var bucket int
		for _, value := range buckets {
			if value == 1 {
				continue
			}
			bucket = value
			break
		}

		if bucket == 4 {
			// Four of a kind
			val = 6
		} else {
			// Full-house
			val = 5
		}
	case 3:
		var bucket int
		for _, value := range buckets {
			if value == 1 {
				continue
			}
			bucket = value
			break
		}

		if bucket == 3 {
			// Three of a kind
			val = 4
		} else {
			// Two pairs
			val = 3
		}
	case 4:
		// One pair
		val = 2
	case 5:
		// Highest card
		val = 1
	}
	return val
}

func (h Hand) ToString() string {
	return fmt.Sprintf("%s of type %d bid %d\n", arrToString(h.cards), h.hType, h.bid)
}

type ByType []*Hand

func (t ByType) Len() int {
	return len(t)
}

func (t ByType) Less(i, j int) bool {
	if t[i].hType != t[j].hType {
		return t[i].hType > t[j].hType
	}

	for k := 0; k < 5; k++ {
		if t[i].cards[k] == t[j].cards[k] {
			continue
		}
		return t[i].cards[k] > t[j].cards[k]
	}

	panic("Comparing same hand...")
	return false
}

func (t ByType) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
