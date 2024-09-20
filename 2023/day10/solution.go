package day10

import (
	"aoc/common"
	"fmt"
	"os"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)
	// 135 is too low
	// 706 is too low

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
}

func calculatePart1(lines []string) int {
	ps := pipeSystem{
		system: lines,
	}
	d1 := []byte(ps.toHtml())
	os.WriteFile("day10/day10.html", d1, 0644)
	startRow, startCol := ps.indexOf(START_TILE)
	loopLen := ps.calcLoopLength(startRow, startCol)
	return loopLen/2 + 1
}

func calculatePart2(lines []string) int {
	return -1
}
