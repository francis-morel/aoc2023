package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	source int
	dest   int
	length int
}

type Seed struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

func main() {
	data, err := os.ReadFile("puzzle_part1.txt")
	if err != nil {
		panic("Failed to open the file")
	}

	dataString := string(data)
	sections := strings.Split(dataString, "\r\n\r\n")

	var seeds []Seed
	var rules = make(map[string][]Rule)

	for _, section := range sections {
		sectionSplit := strings.Split(section, ":")
		title := sectionSplit[0]
		content := sectionSplit[1]

		if title == "seeds" {
			seeds = makeSeeds(content)
			continue
		}
		subTitle := strings.Fields(title)
		rules[subTitle[0]] = makeRules(content)
	}

	seeds = mapSeeds(seeds, rules)

	smallest := math.MaxInt
	for _, seed := range seeds {
		if seed.location < smallest {
			smallest = seed.location
		}
	}

	fmt.Println("Smallest seed:", smallest)
}

func mapSeeds(seeds []Seed, rules map[string][]Rule) []Seed {
	for i := range seeds {
		seeds[i].soil = mapSeedToRule(seeds[i].seed, rules["seed-to-soil"])
		seeds[i].fertilizer = mapSeedToRule(seeds[i].soil, rules["soil-to-fertilizer"])
		seeds[i].water = mapSeedToRule(seeds[i].fertilizer, rules["fertilizer-to-water"])
		seeds[i].light = mapSeedToRule(seeds[i].water, rules["water-to-light"])
		seeds[i].temperature = mapSeedToRule(seeds[i].light, rules["light-to-temperature"])
		seeds[i].humidity = mapSeedToRule(seeds[i].temperature, rules["temperature-to-humidity"])
		seeds[i].location = mapSeedToRule(seeds[i].humidity, rules["humidity-to-location"])
	}

	return seeds
}

func mapSeedToRule(seed int, rules []Rule) int {
	for _, rule := range rules {
		if seed >= rule.source && seed < rule.source+rule.length {
			if rule.dest > rule.source {
				return seed + Abs(rule.source-rule.dest)
			} else {
				return seed - Abs(rule.source-rule.dest)
			}
		}
	}

	// No rule matched, use same value
	return seed
}

func makeSeeds(string string) []Seed {
	stringNumbers := strings.Fields(string)
	seeds := make([]Seed, len(stringNumbers))

	for i, number := range stringNumbers {
		conv, err := strconv.Atoi(number)
		if err != nil {
			panic("Failed to parse seeds data")
		}

		seeds[i] = Seed{
			seed: conv,
		}
	}

	return seeds
}

func makeRules(string string) []Rule {
	stringTrimmed := strings.Trim(string, "\r\n ")
	lines := strings.Split(stringTrimmed, "\r\n")

	rules := make([]Rule, len(lines))

	for i, line := range lines {
		if line == "" {
			continue
		}

		lineNumbers := strings.Fields(line)
		dest, err0 := strconv.Atoi(lineNumbers[0])
		source, err1 := strconv.Atoi(lineNumbers[1])
		length, err2 := strconv.Atoi(lineNumbers[2])

		if err0 != nil || err1 != nil || err2 != nil {
			panic("Failed to parse rules")
		}

		rules[i] = Rule{
			source,
			dest,
			length,
		}
	}

	return rules
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
