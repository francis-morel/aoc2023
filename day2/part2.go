package main

import (
	"fmt"
	"os"
)

func part2() {
	data, err := os.ReadFile("./puzzle_part1.txt")
	if err != nil {
		panic("failed to open input file")
	}

	lineStart := 0
	total := 0

	for pos, char := range data {
		if char == '\n' {
			total += checkLinePart2(data[lineStart:pos])
			lineStart = pos + 1
		}
	}

	fmt.Println(total)
}

func checkLinePart2(line []byte) (power int) {
	var game Game
	var pos = 0

	// Read game number
	game.number, pos = readNextNumber(line[pos:], pos)

	for pos < len(line)-1 {
		var number int
		var color string
		number, pos = readNextNumber(line[pos:], pos)
		color, pos = readNextWord(line[pos:], pos)

		switch color {
		case "red":
			if number > game.red {
				game.red = number
			}
		case "green":
			if number > game.green {
				game.green = number
			}
		case "blue":
			if number > game.blue {
				game.blue = number
			}
		}
	}

	return game.red * game.green * game.blue
}
