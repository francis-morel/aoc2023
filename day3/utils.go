package main

import (
	"strconv"
)

func readNextNumber(line []byte, pos int) (number, newPos int, success bool) {
	currPos := 0
	startPos := 0
	numberReached := isCharNumber(line[0])

	for pos, char := range line {
		if !numberReached {
			if !isCharNumber(char) {
				continue
			}

			numberReached = true
			startPos = pos
		}

		if !isCharNumber(char) {
			currPos = pos
			break
		}
	}

	number, err := strconv.Atoi(string(line[startPos:currPos]))
	if err != nil {
		return 0, 0, false
	}

	return number, pos + currPos, true
}

func isCharNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func readFileDimensions(data []byte) (width, height int) {
	for i := range data {
		if data[i] == '\n' {
			if width == 0 {
				width = i - 1
			}

			height++
		}
	}

	return
}
