package main

import (
	"fmt"
	"os"
	"strings"
)

func part1() {
	data, err := os.ReadFile("puzzle_part1.txt")
	if err != nil {
		panic("Failed to read the file")
	}

	file := string(data)
	fileLines := strings.Split(file, "\r\n")

	total := 0
	for _, line := range fileLines {
		gameNumbers := strings.Split(line, ":")
		if len(gameNumbers) != 2 {
			break
		}

		cardNumbers := strings.Split(gameNumbers[1], " | ")
		winningNumbers := strings.Fields(cardNumbers[0])
		myNumbers := strings.Fields(cardNumbers[1])

		gameScore := 0
		first := true
		for _, winningNumber := range winningNumbers {
			for _, myNumber := range myNumbers {
				if myNumber == winningNumber {
					if first {
						gameScore = 1
						first = false
					} else {
						gameScore *= 2
					}

					break
				}
			}
		}

		total += gameScore
	}

	fmt.Println(total)
}
