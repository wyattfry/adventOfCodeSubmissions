package day06

import (
	"aoc/common"
	"fmt"
	"math"
	"strings"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)
	// 1193500 is too low
	// 1312850

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
}

func calculatePart1(lines []string) int {
	productOfWays := 1
	for _, record := range parseInput(lines) {
		low, high := calculateDistanceRecordButtonTimes(record)
		// fmt.Println("intersections:", low, high)
		waysToWin := calculateWholeNumbersBetween(low, high)
		// fmt.Println("first way", waysToWin[0], " -- last way", waysToWin[len(waysToWin)-1])
		productOfWays *= len(waysToWin)
	}
	return productOfWays
}

func calculatePart2(lines []string) int {
	r := parseInputPart2(lines)
	// fmt.Println("Race Record:", r)
	low, high := calculateDistanceRecordButtonTimes(r)
	// fmt.Println("intersections:", low, high)
	waysToWin := calculateWholeNumbersBetween(low, high)
	// fmt.Println("first way", waysToWin[0], " -- last way", waysToWin[len(waysToWin)-1])
	return len(waysToWin)
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

func parseInputPart2(lines []string) boatRaceRecord {
	time := common.ExtractInts(strings.ReplaceAll(lines[0], " ", ""))[0]
	distance := common.ExtractInts(strings.ReplaceAll(lines[1], " ", ""))[0]

	return boatRaceRecord{
		time:     time,
		distance: distance,
	}
}

func calculateDistanceRecordButtonTimes(r boatRaceRecord) (float64, float64) {
	var output []float64
	for _, i := range []float64{-1, 1} {
		x := (float64(r.time) + i*math.Sqrt(math.Pow(float64(r.time), 2)-4*float64(r.distance))) / 2
		output = append(output, x)
	}
	return output[0], output[1]
}

func calculateWholeNumbersBetween(left, right float64) []int {
	var output []int
	precision := 100
	leftTenths := int(left * float64(precision))
	rightTenths := int(right * float64(precision))
	for i := leftTenths + 1; i < rightTenths; i += 1 {
		if i%precision == 0 {
			output = append(output, i/precision)
		}
	}
	return output
}
