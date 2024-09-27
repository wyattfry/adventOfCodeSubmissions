package day11

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
	gc := parseInput(lines)
	gc.expandBy(1)
	var sum int
	for _, dist := range gc.getAllDistances() {
		sum += dist
	}
	return sum
}

func calculatePart2(lines []string) int {
	gc := parseInput(lines)
	gc.expandBy(1000000 - 1)
	var sum int
	for _, dist := range gc.getAllDistances() {
		sum += dist
	}
	return sum
	// 791,134,890,760 is too high
}

func parseInput(lines []string) galaxyCluster {
	gc := galaxyCluster{}
	for rowIdx, row := range lines {
		for colIdx, col := range row {
			if col == GALAXY_RUNE {
				gc.addGalaxy(location{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}
	return gc
}
