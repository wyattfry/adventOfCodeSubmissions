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

func Solve(file string) {
	lines := common.Readlines(file)
	part1Result := calculatePartNumberSum(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculateGearRatioSum(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
}

func calculatePartNumberSum(lines []string) int {
	var sum int
	for _, num := range parseNumbers(lines, digitIsNextToSymbol) {
		sum += num
	}

	return sum
}

func parseNumbers(schematic []string, filter func(schematicString string, digitIndex int, rowLength int) bool) []int {
	var numbers []int
	var inNumber bool
	var currentNumber int
	var numberIsNextToSymbol bool

	// convert string array schematic to single string to ease traversal
	schematicString := strings.Join(schematic, "")

	// 1. advance char x char until a digit is found or end of input
	for index, c := range schematicString {
		digit := int(c - '0') // convert rune into int
		if digit >= 0 && digit <= 9 {
			// Character is 0-9
			inNumber = true
			currentNumber *= 10

			// 2. add digit to a buffer
			currentNumber += digit

			// 3. check all around for a special symbol
			if numberIsNextToSymbol || filter(schematicString, index, len(schematic[0])) {

				// 4. if one is found mark number to be added to sum
				numberIsNextToSymbol = true
			}

			// Part 2
			// asteriskIndicies = getAdjacentAsterisks(schematicString, index, len(schematic[0])) {
			//
			// }

			// 5. advance to next char, if digit return to step 2
		} else {
			// Character is NOT 0-9
			if inNumber {
				inNumber = false
				if numberIsNextToSymbol {

					// 6. if not a digit, and number is marked, add to sum
					numbers = append(numbers, currentNumber)
					numberIsNextToSymbol = false
				}
				// clear buffer
				currentNumber = 0
			}
		}
	}
	return numbers
}

func isSpecialSymbol(input rune) bool {
	return strings.ContainsRune("#$%&*+-/=@", input)
}

func isDigit(input rune) bool {
	// return strings.ContainsRune("0123456789", input)
	return digitPattern.Match([]byte(string(input)))
}

func getRuneAtIndex(s string, index int) (rune, error) {
	if index < 0 || index > len(s)-1 {
		return 0, errors.New("index out of bounds")
	}
	return rune(s[index]), nil
}

// What if a number is next to two symbols?
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

// Part 2: find '*'s with exactly two numbers around it ( signifying a gear), multiply and get sum of all
// 1. iterate through schematic ([]string), build two maps and array:
//   1.1 [row,col] = numberId -- for every "cell" that is a digit
//   1.2 [numberId] = number -- for every number (1 or more digits)
//   1.3 ["asterisk row,col"]
// 2. make a function that for a given row,col returns an array of adjacent numbers
//    using the above maps
// 3. iterate through asterisk location array, getting all adjacent numbers for each
// 4. for those with exactly two adjacent numbers, multiply them, add to a running sum

type partNumbers struct {
	schematic      []string
	indexIdMap     map[string]int // ["row,col"] = numberId
	idNumberMap    map[int]int    // [numberId] = number
	asteriskRowCol []string
}

func makePartNumbers(schematic []string) partNumbers {
	var output partNumbers
	output.schematic = append(output.schematic, schematic...)
	output.idNumberMap = map[int]int{
		-1: -1,
	}
	output.indexIdMap = map[string]int{
		"foo": -1,
	}

	var inNumber bool
	var currentNumber int
	var currentNumberId int

	for rowNumber, row := range schematic {
		for colNumber, col := range row {
			if isDigit(col) {
				inNumber = true
				currentNumber *= 10
				currentNumber += int(col - '0')
				output.indexIdMap[strconv.Itoa(rowNumber)+","+strconv.Itoa(colNumber)] = currentNumberId
			} else {
				if inNumber {
					inNumber = false
					output.idNumberMap[currentNumberId] = currentNumber
					currentNumberId += 1
					currentNumber = 0
				}
				if col == '*' {
					output.asteriskRowCol = append(output.asteriskRowCol, strconv.Itoa(rowNumber)+","+strconv.Itoa(colNumber))
				}
			}
		}
		// End of Row
		if inNumber {
			inNumber = false
			output.idNumberMap[currentNumberId] = currentNumber
			currentNumberId += 1
			currentNumber = 0
		}
	}

	return output
}

func getAdjacentNumbers(pn partNumbers, row int, col int) []int {
	idNumbers := map[int]int{}

	rcs := [][]int{
		{row - 1, col + 0}, // N
		{row - 1, col + 1}, // NE
		{row - 0, col + 1}, // E
		{row + 1, col + 1}, // SE
		{row + 1, col + 0}, // S
		{row + 1, col - 1}, // SW
		{row + 0, col - 1}, // W
		{row - 1, col - 1}, // NW
	}

	for _, rc := range rcs {
		// Continue if out of bounds
		if rc[0] < 0 || rc[0] > len(pn.schematic)-1 {
			continue
		}
		if rc[1] < 0 || rc[1] > len(pn.schematic[0])-1 {
			continue
		}

		id, ok := pn.indexIdMap[strconv.Itoa(rc[0])+","+strconv.Itoa(rc[1])]
		if ok {
			idNumbers[id] = pn.idNumberMap[id]
		}
	}

	// map[string]int -> []int (of map values)
	r := make([]int, 0, len(idNumbers))
	for _, v := range idNumbers {
		r = append(r, v)
	}
	return r
}

func calculateGearRatioSum(schematic []string) int {
	var sum int
	pn := makePartNumbers(schematic)
	for _, asteriskLocation := range pn.asteriskRowCol {
		s := strings.Split(asteriskLocation, ",")
		row, _ := strconv.Atoi(s[0])
		col, _ := strconv.Atoi(s[1])
		adjacentNumbers := getAdjacentNumbers(pn, row, col)
		if len(adjacentNumbers) == 2 {
			sum += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}

	return sum
}
