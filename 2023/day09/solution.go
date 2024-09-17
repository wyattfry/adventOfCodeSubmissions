package day09

import (
	"aoc/common"
	"fmt"
	"iter"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
}

func calculatePart1(lines []string) int {
	var extrapolatedNextValueSum int

	for sequence := range parseInput(lines) {
		extrapolatedNextValueSum += extrapolate(sequence)
	}

	return extrapolatedNextValueSum
}

func calculatePart2(lines []string) int {
	var extrapolatedPrevValueSum int

	for sequence := range parseInput(lines) {
		extrapolatedPrevValueSum += extrapolatePrev(sequence)
	}

	return extrapolatedPrevValueSum
}

func parseInput(lines []string) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		var sequence []int
		for _, line := range lines {
			sequence = common.ExtractInts(line)
			if !yield(sequence) {
				return
			}
		}
	}
}
