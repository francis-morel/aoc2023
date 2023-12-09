package main

import "strconv"

func readNextNumber(line []byte, pos int) (number, newPos int) {
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
		panic("failed to convert number")
	}

	return number, pos + currPos
}

func readNextWord(line []byte, pos int) (color string, newPos int) {
	currPos := 0
	startPos := 0
	letterReached := isLetter(line[0])

	for pos, char := range line {
		if !letterReached {
			if !isLetter(char) {
				continue
			}

			letterReached = true
			startPos = pos
		}

		if !isLetter(char) {
			currPos = pos
			break
		}
	}

	return string(line[startPos:currPos]), pos + currPos
}

func isCharNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z'
}
