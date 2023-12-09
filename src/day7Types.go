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

func determineHandTypeJoker(cards []int, nJokers int) int {
	oldType := determineHandType(cards)

	switch oldType {
	// There are two (or 1) buckets, and one of them contains jokers.
	case FiveKind, FourKind, FullHouse:
		return FiveKind
	// I either have 3 or one jokers,
	// Either way I am getting a four kind
	case ThreeKind:
		return FourKind
	// If I have two jokers then I can go to a four kind
	// Otherwise just a full house
	case TwoPair:
		if nJokers == 1 {
			return FullHouse
		}
		return FourKind
	// It doesn't matter if I have two or one jokers
	// Either way the best I can get is a ThreeKind
	case OnePair:
		return ThreeKind
	// :(
	case HighCard:
		return TwoPair
	}

	panic("Wait what?")
	return -1
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

	handType := determineHandTypeJoker(cards, nJokers)
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
