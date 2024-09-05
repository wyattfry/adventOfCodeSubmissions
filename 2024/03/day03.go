package day03

import (
	"aoc/common"
	"errors"
	"fmt"
	"strings"
)

func Solve(file string) {
	lines := common.Readlines(file)
	part1Result := calculatePartNumberSum(lines)
	fmt.Println("Solution for part 1:", part1Result)

	// part2Result := calculatePowerOfSetsSum(lines)
	// fmt.Println("Solution for part 2:", part2Result)
}

func calculatePartNumberSum(lines []string) int {
	var sum int
	for _, num := range parseNumbers(lines, digitIsNextToSymbol) {
		sum += num
	}

	return sum
}

// 1. advance char x char until a digit is found or end of input
// 2. add digit to a buffer
// 3. check all around for a special symbol
// 4. if one is found mark number to be added to sum
// 5. advance to next char, if digit return to step 2
// 6. if not a digit, and number is marked, add to sum, clear buffer
// 7. return to step 1

func parseNumbers(schematic []string, filter func(schematicString string, digitIndex int, rowLength int) bool) []int {
	var numbers []int
	var inNumber bool
	var currentNumber int
	var numberIsNextToSymbol bool
	schematicString := strings.Join(schematic, "")
	for index, c := range schematicString {
		digit := int(c - '0')
		if digit >= 0 && digit <= 9 {
			// Character is 0-9
			inNumber = true
			currentNumber *= 10
			currentNumber += digit
			if numberIsNextToSymbol || filter(schematicString, index, len(schematic[0])) {
				numberIsNextToSymbol = true
			}
		} else {
			// Character is NOT 0-9
			if inNumber {
				inNumber = false
				if numberIsNextToSymbol {
					numbers = append(numbers, currentNumber)
					numberIsNextToSymbol = false
				}
				currentNumber = 0
			}
		}
	}
	return numbers
}

func isSpecialSymbol(input rune) bool {
	return strings.ContainsRune("#$%&*+-/=@", input)
}

func getRuneAtIndex(s string, index int) (rune, error) {
	if index < 0 || index > len(s)-1 {
		return 0, errors.New("index out of bounds")
	}
	return rune(s[index]), nil
}

func digitIsNextToSymbol(schematicString string, digitIndex int, rowLength int) bool {
	offsets := []int{
		1,
		-1,
		rowLength,
		rowLength + 1,
		rowLength - 1,
		-rowLength,
		-rowLength + 1,
		-rowLength - 1,
	}
	for _, offset := range offsets {
		r, err := getRuneAtIndex(schematicString, digitIndex+offset)
		if err == nil {
			if isSpecialSymbol(r) {
				return true
			}
		}

	}
	return false
}
