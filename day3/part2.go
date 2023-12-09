package main

import (
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func part2() {
	data, err := os.ReadFile("./puzzle_part1.txt")
	if err != nil {
		panic("failed to open input file")
	}

	width, height := readFileDimensions(data)
	_ = width

	symbols := make([][]bool, height)
	for i := range symbols {
		symbols[i] = make([]bool, width)
	}

	for i, char := range data {
		if char == '*' {
			// TODO simplify this part. I need this calcul because of the \n and \r
			calcIndex := i - (i/(width+2))*2
			symbols[calcIndex/height][calcIndex%width] = true
		}
	}

	total := 0
	symbolNumbers := make(map[string][]int)
	for number, newPos, success := readNextNumber(data, 0); success; number, newPos, success = readNextNumber(data[newPos:], newPos) {
		// TODO simplify this part. I need this calcul because of the \n and \r
		if coord, ok := checkIfSymbolPart2(data, symbols, number, newPos, width, height); ok {
			symbolNumbers[coord.String()] = append(symbolNumbers[coord.String()], number)
		}
	}

	for _, numbers := range symbolNumbers {
		if len(numbers) == 2 {
			gearRatio := 1
			for _, number := range numbers {
				gearRatio *= number
			}
			total += gearRatio
		}
	}

	fmt.Println(total)
}

func checkIfSymbolPart2(data []byte, symbols [][]bool, number int, numberIndex, width, height int) (Coord, bool) {
	// TODO simplify this part. I need this calcul because of the \n and \r
	numberLen := len(fmt.Sprint(number))
	numberStartPos := numberIndex - numberLen
	numberEndPos := numberIndex - 1

	startWidthIndex := ((numberStartPos - (numberStartPos/(width+2))*2) % width) - 1
	endWidthIndex := ((numberEndPos - (numberEndPos/(width+2))*2) % width) + 1
	heightIndex := (numberStartPos - (numberStartPos/(width+2))*2) / height

	for i := heightIndex - 1; i <= heightIndex+1; i++ {
		for j := startWidthIndex; j <= endWidthIndex; j++ {
			if i < 0 || i >= height || j < 0 || j >= width {
				continue
			}

			if symbols[i][j] {
				return Coord{
						x: j,
						y: i,
					},
					true
			}
		}
	}

	return Coord{},
		false
}
