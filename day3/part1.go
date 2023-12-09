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

	width, height := readFileDimensions(data)
	_ = width

	symbols := make([][]bool, height)
	for i := range symbols {
		symbols[i] = make([]bool, width)
	}

	for i, char := range data {
		if char == '\n' || char == '\r' {
			continue
		}

		if (char < '0' || char > '9') && char != '.' {
			// TODO simplify this part. I need this calcul because of the \n and \r
			calcIndex := i - (i/(width+2))*2
			symbols[calcIndex/height][calcIndex%width] = true
		}
	}

	total := 0
	for number, newPos, success := readNextNumber(data, 0); success; number, newPos, success = readNextNumber(data[newPos:], newPos) {
		// TODO simplify this part. I need this calcul because of the \n and \r
		if checkIfSymbolPart1(data, symbols, number, newPos, width, height) {
			total += number
		}
	}

	fmt.Println(total)
}

func checkIfSymbolPart1(data []byte, symbols [][]bool, number int, numberIndex, width, height int) bool {
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
				return true
			}
		}
	}

	return false
}
