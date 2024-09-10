package day06

import (
	"aoc/common"
	"fmt"
	"math"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	// part2Result := calculatePart2(lines)
	// fmt.Println("Solution for part 2:", part2Result, file)
}

func calculatePart1(lines []string) int {
	// for _, record := range parseInput(lines) {

	// }
	return -1
}

type boatRaceRecord struct {
	time, distance int
}

// Parse each line (string) into the starting seed numbers and source-to-destination maps
func parseInput(lines []string) []boatRaceRecord {
	records := []boatRaceRecord{}
	times := common.ExtractInts(lines[0])
	distances := common.ExtractInts(lines[1])

	for index, time := range times {
		records = append(records, boatRaceRecord{
			time:     time,
			distance: distances[index],
		})
	}

	return records
}

func calculateDistanceBetweenIntersections(r boatRaceRecord) float64 {
	left := (float64(r.time) - math.Sqrt(math.Pow(float64(r.time), 2)-4*float64(r.distance))) / 2
	right := (float64(r.time) + math.Sqrt(math.Pow(float64(r.time), 2)-4*float64(r.distance))) / 2
	fmt.Println(r, left, right)
	return right - left
}
