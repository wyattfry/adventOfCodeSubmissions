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
	return -1
}

func calculatePart2(lines []string) int {
	return -1
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
