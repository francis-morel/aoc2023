package main

import (
	"fmt"
	"os"
	"strings"
)

type Scratchcard struct {
	quantityWinningNumbers int
	occurences             int
}

func (s *Scratchcard) fill(line string) {
	gameNumbers := strings.Split(line, ":")
	if len(gameNumbers) != 2 {
		return
	}

	cardNumbers := strings.Split(gameNumbers[1], " | ")
	winningNumbers := strings.Fields(cardNumbers[0])
	myNumbers := strings.Fields(cardNumbers[1])

	gameScore := 0
	quantityWins := 0
	for _, winningNumber := range winningNumbers {
		for _, myNumber := range myNumbers {
			if myNumber == winningNumber {
				if quantityWins == 0 {
					gameScore = 1
				} else {
					gameScore *= 2
				}

				quantityWins++
				break
			}
		}
	}

	s.quantityWinningNumbers = quantityWins
}

func part2() {
	data, err := os.ReadFile("puzzle_part1.txt")
	if err != nil {
		panic("Failed to read the file")
	}

	file := string(data)
	fileLines := strings.Split(file, "\r\n")

	deck := make([]Scratchcard, len(fileLines)-1)

	handleScratchcards(fileLines, deck)
	fmt.Println(countCards(deck))
}

func handleScratchcards(fileLines []string, deck []Scratchcard) {
	for i, line := range fileLines {
		if line == "" {
			break
		}

		deck[i].fill(line)
		deck[i].occurences++

		for j := 0; j < deck[i].quantityWinningNumbers; j++ {
			cardNumber := i + j + 1

			if cardNumber > len(deck)-1 {
				break
			}

			if deck[i].occurences > 0 {
				deck[cardNumber].occurences += deck[i].occurences
			} else {
				deck[cardNumber].occurences += 1
			}
		}
	}
}

func countCards(deck []Scratchcard) int {
	quantityCards := 0

	for _, card := range deck {
		quantityCards += card.occurences
	}

	return quantityCards
}
