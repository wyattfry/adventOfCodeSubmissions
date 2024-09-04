package day02

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
)

var (
	colorsCount = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func Solve(file string) {
	var gameIdSum int
	for index, line := range common.Readlines(file) {
		for color, count := range colorsCount {
			if getHighestForColor(line, color) > count {
				continue
			}
		}
		gameId := index + 1
		gameIdSum += gameId
	}
	fmt.Println("Solution for part 1:", gameIdSum)

	var powerOfSetsSum int
	for _, line := range common.Readlines(file) {
		powerOfSetsSum += getPowerOfSet(line)
	}
	fmt.Println("Solution for part 2:", powerOfSetsSum)
}

func getHighestForColor(gamestring string, color string) int {
	matches := regexp.MustCompile(fmt.Sprintf("(\\d+)( %s)", color)).FindAllStringSubmatch(gamestring, -1)
	var highest int
	for _, str := range matches {
		if num, err := strconv.Atoi(str[1]); err == nil {
			if num > highest {
				highest = num
			}
		}
	}

	return highest
}

func getPowerOfSet(gamestring string) int {
	powerOfSet := 1
	for color := range colorsCount {
		powerOfSet *= getHighestForColor(gamestring, color)
	}
	return powerOfSet
}
