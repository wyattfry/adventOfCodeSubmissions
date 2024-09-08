package day03

import (
	"testing"
)

// TestCalculatePartNumberSum tests the calculation of the sum of part numbers next to special symbols.
func TestCalculatePartNumberSum(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{"123*45", "67890", "#2345"},
			expected: 4568013,
		},
		{
			input:    []string{"1#23", "45*6", "7+89"},
			expected: 2413,
		},
		{
			input:    []string{"##", "##"},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := calculatePartNumberSum(test.input)
		if result != test.expected {
			t.Errorf("calculatePartNumberSum(%v) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

// TestDigitIsNextToSymbol tests if digits next to symbols are correctly identified.
func TestDigitIsNextToSymbol(t *testing.T) {
	tests := []struct {
		schematicString string
		digitIndex      int
		rowLength       int
		expected        bool
	}{
		{"123#45", 2, 6, true},
		{"123456", 2, 6, false},
		{"#12345", 1, 6, true},
		{"12#345", 4, 6, false},
	}

	for _, test := range tests {
		result := digitIsNextToSymbol(test.schematicString, test.digitIndex, test.rowLength)
		if result != test.expected {
			t.Errorf("digitIsNextToSymbol(%v, %d, %d) = %v; expected %v", test.schematicString, test.digitIndex, test.rowLength, result, test.expected)
		}
	}
}

// TestCalculateGearRatioSum tests the calculation of gear ratios from part numbers.
func TestCalculateGearRatioSum(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"..2..",
				"..*..",
				"...24"},
			expected: 48, // Example of calculation, adjust based on gear logic
		},
		{
			input: []string{
				"12..",
				".*..",
				".3.2"},
			expected: 36, // Example of calculation, adjust based on gear logic
		},
		{
			input: []string{
				"...2",
				"3.*.",
				"..4."},
			expected: 8,
		},
	}

	for _, test := range tests {
		result := calculateGearRatioSum(test.input)
		if result != test.expected {
			t.Errorf("calculateGearRatioSum(%v) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

// TestMakePartNumbers tests the construction of partNumbers struct.
func TestMakePartNumbers(t *testing.T) {
	schematic := []string{"12*34", "56#78", "90"}
	pn := makePartNumbers(schematic)

	if len(pn.asteriskRowCol) != 1 || pn.asteriskRowCol[0] != "0,2" {
		t.Errorf("Expected asterisk position '0,2', got %v", pn.asteriskRowCol)
	}

	if pn.idNumberMap[0] != 12 || pn.idNumberMap[1] != 34 || pn.idNumberMap[2] != 56 || pn.idNumberMap[3] != 78 || pn.idNumberMap[4] != 90 {
		t.Errorf("Part numbers mapping incorrect: %v", pn.idNumberMap)
	}
}

// TestGetAdjacentNumbers tests the retrieval of numbers adjacent to a given cell.
func TestGetAdjacentNumbers(t *testing.T) {
	schematic := []string{
		"34*.",
		"12..",
		"..5."}
	pn := makePartNumbers(schematic)
	adjacentNumbers := getAdjacentNumbers(pn, 0, 2)

	if len(adjacentNumbers) != 2 || adjacentNumbers[0] != 12 || adjacentNumbers[1] != 34 {
		t.Errorf("Expected adjacent numbers [12, 34], got %v", adjacentNumbers)
	}
}

// TestIsSpecialSymbolAt tests the identification of special symbols in the schematic.
func TestIsSpecialSymbolAt(t *testing.T) {
	schematicString := "123#456"
	index := 3

	if !isSpecialSymbolAt(schematicString, index) {
		t.Errorf("Expected symbol at index %d in %v", index, schematicString)
	}

	index = 2
	if isSpecialSymbolAt(schematicString, index) {
		t.Errorf("Did not expect symbol at index %d in %v", index, schematicString)
	}
}

// TestIsDigit tests the identification of digits.
func TestIsDigit(t *testing.T) {
	if !isDigit('5') {
		t.Errorf("Expected '5' to be identified as a digit.")
	}

	if isDigit('#') {
		t.Errorf("Expected '#' to not be identified as a digit.")
	}
}

// TestGetRuneAtIndex tests safe rune retrieval from a string.
func TestGetRuneAtIndex(t *testing.T) {
	r, err := getRuneAtIndex("12345", 2)
	if err != nil || r != '3' {
		t.Errorf("Expected '3' at index 2, got %c", r)
	}

	_, err = getRuneAtIndex("12345", 5)
	if err == nil {
		t.Errorf("Expected an error for index out of bounds")
	}
}
