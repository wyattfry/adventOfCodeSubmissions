package day04

import (
	"aoc/common"
	"fmt"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
}

func calculatePart1(lines []string) int {
	return -1
}

func calculatePart2(lines []string) int {
	return -1
}
