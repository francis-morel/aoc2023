package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	data, err := os.ReadFile("puzzle_part1.txt")
	if err != nil {
		panic("Failed to open file")
	}

	dataStr := string(data)
	fileLines := strings.Split(dataStr, "\n")

	timeSections := strings.Split(fileLines[0], ":")
	distanceSections := strings.Split(fileLines[1], ":")

	times := strings.Fields(timeSections[1])
	distances := strings.Fields(distanceSections[1])

	product := 1
	for i := range times {
		time, err1 := strconv.Atoi(times[i])
		distance, err2 := strconv.Atoi(distances[i])

		if err1 != nil || err2 != nil {
			panic("Failed to convert data of race#" + fmt.Sprint(i))
		}

		product *= waysToBeat(time, distance)
	}

	fmt.Println(product)

}

func part2() {
	fmt.Print(waysToBeat(54946592, 302147610291404))
}

func waysToBeat(timeAvailable, bestDistance int) int {
	canBeatCount := 0

	for i := 1; i < timeAvailable; i++ {
		if testRace(timeAvailable, i) > bestDistance {
			canBeatCount++
		}
	}
	return canBeatCount
}

func testRace(totalTime, holdTime int) int {
	runTime := totalTime - holdTime
	return runTime * holdTime
}
