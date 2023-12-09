package main

import (
	"fmt"
)

type Hand struct {
	cards []int
	bid   int
	hType int
}

const (
	Joker     = 11
	FiveKind  = 7
	FourKind  = 6
	FullHouse = 5
	ThreeKind = 4
	TwoPair   = 3
	OnePair   = 2
	HighCard  = 1
)

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
		val = FiveKind
		break
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
			val = FourKind
		} else {
			val = FullHouse
		}
		break
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
			val = ThreeKind
		} else {
			// Two pairs
			val = TwoPair
		}
		break
	case 4:
		// One pair
		val = OnePair
		break
	case 5:
		// Highest card
		val = HighCard
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

	for k := 0; k < len(t[i].cards); k++ {
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

func numberOfJokers(cards []int) int {
	count := 0
	for _, card := range cards {
		if card == Joker {
			count++
		}
	}
	return count
}

func NewJokerHand(cards []int, bid int) *Hand {
	nJokers := numberOfJokers(cards)
	if nJokers == 0 {
		return NewHand(cards, bid)
	}

	handType := determineHandTypeJoker(cards)
	weakenJokers(cards)
	return &Hand{cards: cards, bid: bid, hType: handType}
}

func weakenJokers(cards []int) {
	for i, card := range cards {
		if card == Joker {
			cards[i] = 1
		}
	}
}

func (h Hand) ToStringComparison() string {
	return fmt.Sprintf("%d %s\n", h.hType, arrToString(h.cards))
}

func determineHandTypeJoker(cards []int) int {
	jokerIndexes := make([]int, 0)
	buckets := make(map[int]int)

	for i, card := range cards {
		if card == Joker {
			jokerIndexes = append(jokerIndexes, i)
			continue
		}
		buckets[card]++
	}

	bestType := -1
	var tempType int
	for k, _ := range buckets {
		for _, index := range jokerIndexes {
			cards[index] = k
		}
		tempType = determineHandType(cards)
		if tempType > bestType {
			bestType = tempType
		}
	}

	for _, index := range jokerIndexes {
		cards[index] = Joker
	}

	// This means that there are no buckets, hence a full hand of jokers
	if bestType == -1 {
		bestType = FiveKind
	}

	return bestType
}
