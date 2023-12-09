package main

import (
	"fmt"
	"os"
)

func part1() {
	data, err := os.ReadFile("./puzzle_part1.txt")
	if err != nil {
		panic("failed to open input file")
	}

	lineStart := 0
	total := 0

	for pos, char := range data {
		if char == '\n' {
			id, valid := checkLinePart1(data[lineStart:pos])
			if valid {
				total += id
			}

			lineStart = pos + 1
		}
	}

	fmt.Println(total)
}

func checkLinePart1(line []byte) (id int, valid bool) {
	var game Game
	var pos = 0

	// Read game number
	game.number, pos = readNextNumber(line[pos:], pos)

	for pos < len(line)-1 {
		if line[pos] == ';' {
			if !game.isValid() {
				return game.number, false
			}

			game.red = 0
			game.green = 0
			game.blue = 0
		}

		var number int
		var color string
		number, pos = readNextNumber(line[pos:], pos)
		color, pos = readNextWord(line[pos:], pos)

		switch color {
		case "red":
			game.red = number
		case "green":
			game.green = number
		case "blue":
			game.blue = number
		}
	}

	return game.number, game.isValid()
}
