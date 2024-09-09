package day05

import (
	"aoc/common"
	"fmt"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
	// 2254687 is too high
}

func calculatePart1(lines []string) int {
	lowestLocation := 9999999999999

	seedNumbers, srcDstMaps := parseInput(lines)

	for _, seedNum := range seedNumbers {
		convertSeedToLocation(seedNum, srcDstMaps, &lowestLocation)
	}

	return lowestLocation
}

type conversionRule struct {
	destRangeStart, sourceRangeStart, rangeLength int
}

// Parse each line (string) into the starting seed numbers and source-to-destination maps
func parseInput(lines []string) (seedNumbers []int, sourceDestionationMaps [][]conversionRule) {
	var inMap bool
	var currentRules []conversionRule

	lines = append(lines, "")

	for _, line := range lines {
		if inMap {
			if len(line) == 0 {
				inMap = false
				sourceDestionationMaps = append(sourceDestionationMaps, currentRules)
				currentRules = []conversionRule{}
			} else {
				currentRules = append(currentRules, parseConversionRule(line))
			}
		} else if strings.HasSuffix(line, " map:") {
			inMap = true
		} else if strings.HasPrefix(line, "seeds: ") {
			seedNumbers = common.ExtractInts(line)
		}
	}

	return seedNumbers, sourceDestionationMaps
}

func parseConversionRule(rulestring string) conversionRule {
	nums := common.ExtractInts(rulestring)
	return conversionRule{
		destRangeStart:   nums[0],
		sourceRangeStart: nums[1],
		rangeLength:      nums[2],
	}
}

func convertSeedToLocation(seedNumber int, sourceDestionationMaps [][]conversionRule, lowestLocation *int) {
	currentNumber := seedNumber
	for _, srcDstMap := range sourceDestionationMaps {
		for _, rule := range srcDstMap {
			if currentNumber >= rule.sourceRangeStart && currentNumber < rule.sourceRangeStart+rule.rangeLength {
				currentNumber += rule.destRangeStart - rule.sourceRangeStart
				break
			}
		}
	}
	if currentNumber < *lowestLocation {
		*lowestLocation = currentNumber
	}
}

// I thought parallelization would make the solution run faster, but it didn't.
// The sync version took about two minutes, this was estimating over an hour.
// func convertSeedToLocationAsync(wg *sync.WaitGroup, seedNumber int, sourceDestionationMaps [][]conversionRule, lowestLocation *int, bar *progressbar.ProgressBar) {
// 	convertSeedToLocation(seedNumber, sourceDestionationMaps, lowestLocation)
// 	bar.Add(1)
// 	wg.Done()
// }

// Part 2: the `seeds:` line no longer is a list of seedNumbers, but rather
// seed ranges. 12 6 = starts at 12, range is 6, 12 to 17
func calculatePart2(lines []string) int {
	seedNumbers, srcDstMaps := parseInput(lines)
	var totalSeedsToConvert int

	lowestLocation := 9999999999999

	for index, seedNum := range seedNumbers {
		if index%2 != 1 {
			continue
		}
		totalSeedsToConvert += seedNum
	}
	bar := progressbar.Default(int64(totalSeedsToConvert), "Converting Seed numbers to Location numbers")

	// wg := &sync.WaitGroup{}
	for index, seedNum := range seedNumbers {
		// Even index numbers are seed starting numbers
		if index%2 != 1 {
			continue
		}

		// Odd index are seed ranges
		startingSeedNumber := seedNumbers[index-1]
		seedNumberRange := seedNum

		// fmt.Println("Seed -> Location from", startingSeedNumber, "to", startingSeedNumber+seedNumberRange-1)

		for i := startingSeedNumber; i < startingSeedNumber+seedNumberRange; i += 1 {
			// wg.Add(1)
			// convertSeedToLocationAsync(wg, i, srcDstMaps, &lowestLocation, bar)
			convertSeedToLocation(i, srcDstMaps, &lowestLocation)
			bar.Add(1)
		}
	}
	// wg.Wait()

	return lowestLocation
}
