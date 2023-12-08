package main

import (
	"fmt"
	"os"
)

type game struct {
	number int
	red    int
	green  int
	blue   int
}

func (game game) isValid() bool {
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	return game.red <= redLimit && game.green <= greenLimit && game.blue <= blueLimit
}

func main() {
	data, err := os.ReadFile("./puzzle_part1.txt")
	if err != nil {
		panic("failed to open input file")
	}

	lineStart := 0
	total := 0

	for pos, char := range data {
		if char == '\n' {
			id, valid := checkLine(data[lineStart:pos])
			if valid {
				total += id
			}

			lineStart = pos + 1
		}
	}

	fmt.Println(total)
}

func checkLine(line []byte) (id int, valid bool) {
	fmt.Println(string(line))
	var game game
	// gameNumber := make([]byte, 0)
	//
	// for _, char := range line {
	//
	// }

	return game.number, game.isValid()
}
