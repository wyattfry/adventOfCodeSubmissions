package day04

import (
	"aoc/common"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
}

func calculatePart1(lines []string) int {
	var pointSum int
	for _, line := range lines {
		card := parseLine(line)
		winnerCount := getWinnerCount(card)
		points := int(math.Pow(2, float64(winnerCount-1)))
		pointSum += points
	}
	return pointSum
}

func getWinnerCount(sc scratchcard) int {
	var count int
	for _, num := range sc.winners {
		if slices.Contains(sc.myNumbers, num) {
			count += 1
		}
	}
	return count
}

type scratchcard struct {
	winners   []int
	myNumbers []int
}

// Parse each line (string) into a scratchcard struct containing 2 []ints, for winning numbers and "my numbers"
func parseLine(line string) scratchcard {
	output := scratchcard{
		winners:   []int{},
		myNumbers: []int{},
	}
	numbers := strings.Split(strings.Split(line, ":")[1], "|")
	for _, numstr := range regexp.MustCompile(`\d+`).FindAllString(numbers[0], -1) {
		num, _ := strconv.Atoi(numstr)
		output.winners = append(output.winners, num)
	}
	for _, numstr := range regexp.MustCompile(`\d+`).FindAllString(numbers[1], -1) {
		num, _ := strconv.Atoi(numstr)
		output.myNumbers = append(output.myNumbers, num)
	}

	return output
}

func calculatePart2(lines []string) int {
	var cardSum int
	cardCounts := make([]int, len(lines))
	for i := range cardCounts {
		cardCounts[i] = 1
	}
	for index, line := range lines {
		cardSum += cardCounts[index]
		card := parseLine(line)
		for i := 0; i < getWinnerCount(card); i++ {
			cardCounts[index+1+i] += cardCounts[index]
		}
	}
	return cardSum
}
