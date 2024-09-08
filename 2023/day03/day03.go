package day03

import (
	"aoc/common"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	digitPattern = regexp.MustCompile(`\d{1}`)
)

// Solve reads the input file and prints the solutions for both parts.
func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePartNumberSum(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculateGearRatioSum(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
}

// calculatePartNumberSum calculates the sum of numbers adjacent to special symbols.
func calculatePartNumberSum(lines []string) int {
	numbers := parseNumbers(lines, digitIsNextToSymbol)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// parseNumbers extracts numbers from a 2D schematic based on a filter function.
func parseNumbers(schematic []string, filter func(schematicString string, digitIndex int, rowLength int) bool) []int {
	var numbers []int
	schematicString := strings.Join(schematic, "")
	inNumber := false
	currentNumber := 0
	numberIsNextToSymbol := false

	for index, c := range schematicString {
		if isDigit(c) {
			// Start or continue a number
			inNumber = true
			currentNumber = currentNumber*10 + int(c-'0')

			// Check for symbols around the digit
			if numberIsNextToSymbol || filter(schematicString, index, len(schematic[0])) {
				numberIsNextToSymbol = true
			}
		} else {
			// Handle end of a number
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

// digitIsNextToSymbol checks if a digit is adjacent to any special symbol.
func digitIsNextToSymbol(schematicString string, digitIndex int, rowLength int) bool {
	for _, offset := range generateOffsets(rowLength) {
		if isSpecialSymbolAt(schematicString, digitIndex+offset) {
			return true
		}
	}
	return false
}

// generateOffsets returns offsets for checking surrounding characters.
func generateOffsets(rowLength int) []int {
	return []int{
		1, -1,
		rowLength, rowLength + 1, rowLength - 1,
		-rowLength, -rowLength + 1, -rowLength - 1,
	}
}

// isSpecialSymbolAt checks if there is a special symbol at the specified index.
func isSpecialSymbolAt(s string, index int) bool {
	r, err := getRuneAtIndex(s, index)
	if err != nil {
		return false
	}
	return isSpecialSymbol(r)
}

// isSpecialSymbol checks if a character is a special symbol.
func isSpecialSymbol(input rune) bool {
	return strings.ContainsRune("#$%&*+-/=@", input)
}

// isDigit checks if a character is a digit.
func isDigit(input rune) bool {
	return digitPattern.MatchString(string(input))
}

// getRuneAtIndex safely gets a rune at a specific index.
func getRuneAtIndex(s string, index int) (rune, error) {
	if index < 0 || index >= len(s) {
		return 0, errors.New("index out of bounds")
	}
	return rune(s[index]), nil
}

// Part 2 logic for calculating the gear ratio sum.
func calculateGearRatioSum(schematic []string) int {
	pn := makePartNumbers(schematic)
	sum := 0
	for _, asteriskLocation := range pn.asteriskRowCol {
		row, col := parseRowCol(asteriskLocation)
		adjacentNumbers := getAdjacentNumbers(pn, row, col)
		if len(adjacentNumbers) == 2 {
			sum += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}
	return sum
}

// parseRowCol parses a row and column from a string representation.
func parseRowCol(location string) (int, int) {
	s := strings.Split(location, ",")
	row, _ := strconv.Atoi(s[0])
	col, _ := strconv.Atoi(s[1])
	return row, col
}

// makePartNumbers constructs a struct mapping indices to numbers and identifying asterisk positions.
func makePartNumbers(schematic []string) partNumbers {
	pn := partNumbers{
		schematic:      schematic,
		indexIdMap:     make(map[string]int),
		idNumberMap:    make(map[int]int),
		asteriskRowCol: []string{},
	}

	currentNumberId := 0
	inNumber := false
	currentNumber := 0

	for rowNumber, row := range schematic {
		for colNumber, col := range row {
			if isDigit(col) {
				inNumber = true
				currentNumber = currentNumber*10 + int(col-'0')
				pn.indexIdMap[fmt.Sprintf("%d,%d", rowNumber, colNumber)] = currentNumberId
			} else {
				if inNumber {
					pn.idNumberMap[currentNumberId] = currentNumber
					currentNumberId++
					inNumber = false
					currentNumber = 0
				}
				if col == '*' {
					pn.asteriskRowCol = append(pn.asteriskRowCol, fmt.Sprintf("%d,%d", rowNumber, colNumber))
				}
			}
		}
		if inNumber {
			pn.idNumberMap[currentNumberId] = currentNumber
			currentNumberId++
			inNumber = false
			currentNumber = 0
		}
	}
	return pn
}

// getAdjacentNumbers retrieves numbers adjacent to a specific cell.
func getAdjacentNumbers(pn partNumbers, row, col int) []int {
	idNumbers := make(map[int]int)

	for _, rc := range getSurroundingCells(row, col) {
		if rc[0] < 0 || rc[0] >= len(pn.schematic) || rc[1] < 0 || rc[1] >= len(pn.schematic[0]) {
			continue
		}
		id, exists := pn.indexIdMap[fmt.Sprintf("%d,%d", rc[0], rc[1])]
		if exists {
			idNumbers[id] = pn.idNumberMap[id]
		}
	}

	result := make([]int, 0, len(idNumbers))
	for _, v := range idNumbers {
		result = append(result, v)
	}
	return result
}

// getSurroundingCells returns coordinates for surrounding cells.
func getSurroundingCells(row, col int) [][]int {
	return [][]int{
		{row - 1, col}, {row - 1, col + 1},
		{row, col + 1}, {row + 1, col + 1},
		{row + 1, col}, {row + 1, col - 1},
		{row, col - 1}, {row - 1, col - 1},
	}
}

// Struct to represent the part numbers and their relations in the grid.
type partNumbers struct {
	schematic      []string
	indexIdMap     map[string]int
	idNumberMap    map[int]int
	asteriskRowCol []string
}
