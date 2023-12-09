package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func day7PartOne(path string) int {
	dat, err := os.ReadFile(path)
	check(err)

	hands := parseCards(dat)
	sort.Sort(ByType(hands))

	aggregator := 0
	var modifier int
	for i, hand := range hands {
		modifier = len(hands) - i
		aggregator += hand.bid * modifier
	}

	return aggregator
}

func parseCards(dat []byte) []*Hand {
	var err error
	hands := make([]*Hand, 0)

	cards := make([]int, 0)
	var card int
	currNumber := make([]byte, 0)
	var bid int

	parsingCards := true
	for _, b := range dat {
		if b == ' ' || b == '\n' {
			if b == '\n' {
				bid, err = strconv.Atoi(string(currNumber))
				check(err)
				currNumber = make([]byte, 0)

				hands = append(hands, NewHand(cards, bid))
				cards = make([]int, 0)
			}
			parsingCards = !parsingCards
			continue
		}

		if parsingCards {
			if !unicode.IsDigit(rune(b)) {
				card = mapToInt(b)
			} else {
				card, err = strconv.Atoi(string(b))
				check(err)
			}
			cards = append(cards, card)
			continue
		}

		currNumber = append(currNumber, b)
	}
	return hands
}

func mapToInt(b byte) int {
	var val int
	switch b {
	case 'A':
		val = 14
		break
	case 'K':
		val = 13
		break
	case 'Q':
		val = 12
		break
	case 'J':
		val = 11
		break
	case 'T':
		val = 10
		break
	default:
		panic(fmt.Sprintf("Unknown char %s", string(b)))
	}
	return val
}

func day7PartTwo(path string) int {
	return len(path)
}
