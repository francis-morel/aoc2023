package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	part2()
}

func part2() {
	data, err := os.ReadFile("./part1_input.txt")
	if err != nil {
		panic("problem reading input file")
	}

	result := []byte{' ', ' ', '\n'}
	var total int = 0

	// spelled := map[string]byte{
	// 	"one":   '1',
	// 	"two":   '2',
	// 	"three": '3',
	// 	"four":  '4',
	// 	"five":  '5',
	// 	"six":   '6',
	// 	"seven": '7',
	// 	"eight": '8',
	// 	"nine":  '9',
	// }

	for _, char := range data {
		if char >= '0' && char <= '9' {
			if result[0] == ' ' {
				result[0] = char
			} else {
				result[1] = char
			}
		} else if char >= 'a' && char <= 'z' {
		} else if char == '\n' {
			// Only character left is '\n'
			if result[1] == ' ' {
				result[1] = result[0]
			}

			number := string(result[0]) + string(result[1])
			num, err := strconv.Atoi(number)
			if err != nil {
				panic("invalid number conversion")
			}

			total += num
		}
	}

	fmt.Println(total)
}

func part1() {
	data, err := os.ReadFile("./part1_input.txt")
	if err != nil {
		panic("problem reading input file")
	}

	out, err := os.Create("./part1_output.txt")
	if err != nil {
		panic("problem creating output file")
	}
	defer out.Close()

	result := []byte{' ', ' ', '\n'}
	var total int = 0

	for _, char := range data {
		if char >= '0' && char <= '9' {
			if result[0] == ' ' {
				result[0] = char
			} else {
				result[1] = char
			}
		}

		if char == '\n' {
			if result[1] == ' ' {
				result[1] = result[0]
			}

			number := string(result[0]) + string(result[1])
			num, err := strconv.Atoi(number)
			if err != nil {
				panic("invalid number conversion")
			}

			total += num

			out.Write(result)
			result[0] = ' '
			result[1] = ' '
		}
	}

	fmt.Println(total)
}
