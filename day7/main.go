package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const PART2 bool = true

type hand struct {
	cards         []rune
	bid           int
	strength      int
	cardsStrength []int
}

func (h *hand) fill(cards, bid string) {
	h.cards = []rune(cards)
	var err error
	h.bid, err = strconv.Atoi(bid)

	if err != nil {
		panic("Failed to convert bid to int value")
	}

	h.fillStrength()
}

func (h *hand) fillStrength() {
	cards := make(map[rune]int)

	for _, card := range h.cards {
		_, ok := cards[card]
		if ok {
			cards[card]++
		} else {
			cards[card] = 1
		}
	}

	if PART2 {
		jokerCount, jokerPresent := cards['J']

		if jokerPresent {
			keyOfBiggest := ' '
			for k, v := range cards {
				if keyOfBiggest == ' ' && k != 'J' {
					keyOfBiggest = k
					continue
				}

				if v > cards[keyOfBiggest] && k != 'J' {
					keyOfBiggest = k
				}
			}

			cards[keyOfBiggest] += jokerCount
			delete(cards, 'J')
		}

	}

	switch len(cards) {
	case 5:
		h.strength = 1
	case 4:
		h.strength = 2
	case 1:
		h.strength = 7
	}

	if len(cards) == 2 {
		fourOfAKind := false
		for _, v := range cards {
			if v == 4 {
				fourOfAKind = true
			}
		}

		if fourOfAKind {
			h.strength = 6
		} else {
			h.strength = 5
		}
	}

	if len(cards) == 3 {
		threeOfAKind := false
		for _, v := range cards {
			if v == 3 {
				threeOfAKind = true
			}
		}

		if threeOfAKind {
			h.strength = 4
		} else {
			h.strength = 3
		}
	}

}

func (h *hand) fillCardsStrength() {
	if len(h.cardsStrength) > 0 {
		return
	}

	h.cardsStrength = []int{0, 0, 0, 0, 0}
	for i, card := range h.cards {
		h.cardsStrength[i] = cardStrength(card)
	}
}

func (h1 *hand) isWeaker(h2 *hand) bool {
	if h1.strength != h2.strength {
		return h1.strength < h2.strength
	}

	h1.fillCardsStrength()
	h2.fillCardsStrength()

	for i := 0; i < 5; i++ {
		if h1.cardsStrength[i] != h2.cardsStrength[i] {
			return h1.cardsStrength[i] < h2.cardsStrength[i]
		}
	}

	return false
}

func main() {
	data, err := os.ReadFile("puzzle_part1.txt")
	if err != nil {
		panic("Failed to open file")
	}

	dataStr := string(data)
	dataSplit := strings.Fields(dataStr)

	hands := make([]hand, len(dataSplit)/2)
	for i := 0; i < len(dataSplit); i += 2 {
		hands[i/2].fill(dataSplit[i], dataSplit[i+1])
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].isWeaker(&hands[j])
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
		fmt.Println(string(hand.cards), hand.bid, hand.strength)
	}

	fmt.Println(total)

}

func cardStrength(card rune) int {
	// in ascii '2' = 50 and '9' = 57
	if card >= '2' && card <= '9' {
		return int(card) - 49
	}

	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if !PART2 {
			return 11
		} else {
			return 0
		}
	case 'T':
		return 10
	}

	return 0
}
