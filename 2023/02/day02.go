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
	lines := common.Readlines(file)
	part1Result := calculateGameIDSum(lines)
	fmt.Println("Solution for part 1:", part1Result)

	part2Result := calculatePowerOfSetsSum(lines)
	fmt.Println("Solution for part 2:", part2Result)
}

func calculateGameIDSum(lines []string) int {
	var gameIdSum int
	for index, line := range lines {
		if isValidGameID(line) {
			gameIdSum += index + 1
		}
	}
	return gameIdSum
}

func isValidGameID(line string) bool {
	for color, count := range colorsCount {
		if getHighestForColor(line, color) > count {
			return false
		}
	}
	return true
}

func calculatePowerOfSetsSum(lines []string) int {
	var powerOfSetsSum int
	for _, line := range lines {
		powerOfSetsSum += getPowerOfSet(line)
	}
	return powerOfSetsSum
}

func getHighestForColor(gameString, color string) int {
	re := regexp.MustCompile(fmt.Sprintf(`(\d+)\s+%s`, color))
	matches := re.FindAllStringSubmatch(gameString, -1)

	highest := 0
	for _, match := range matches {
		if num, err := strconv.Atoi(match[1]); err == nil && num > highest {
			highest = num
		}
	}
	return highest
}

func getPowerOfSet(gameString string) int {
	powerOfSet := 1
	for color := range colorsCount {
		powerOfSet *= getHighestForColor(gameString, color)
	}
	return powerOfSet
}
