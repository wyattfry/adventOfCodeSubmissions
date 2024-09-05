package day03

import (
	"reflect"
	"testing"
)

type parseNumbersTestCase struct {
	input  []string
	output []int
}

func TestParseNumbers(t *testing.T) {
	cases := []parseNumbersTestCase{
		{input: []string{".1."}, output: []int{1}},
		{input: []string{"#1@"}, output: []int{1}},
		{input: []string{"..123...234..."}, output: []int{123, 234}},
	}

	for _, tc := range cases {
		result := parseNumbers(tc.input, func(string, int, int) bool {
			return true
		})
		if !reflect.DeepEqual(result, tc.output) {
			t.Errorf(`%s("%s") = %d, wanted %d`, "parseNumbers", tc.input, result, tc.output)
		}
	}
}

type digitIsNextToSymbolTestCaseInput struct {
	schematicString string
	digitIndex      int
	rowLength       int
}

type digitIsNextToSymbolTestCase struct {
	input  digitIsNextToSymbolTestCaseInput
	output bool
}

func TestDigitIsNextToSymbol(t *testing.T) {
	cases := []digitIsNextToSymbolTestCase{
		{input: digitIsNextToSymbolTestCaseInput{
			schematicString: "*1",
			digitIndex:      1,
			rowLength:       2,
		}, output: true},
		{input: digitIsNextToSymbolTestCaseInput{
			schematicString: ".1",
			digitIndex:      1,
			rowLength:       2,
		}, output: false},
		{input: digitIsNextToSymbolTestCaseInput{
			schematicString: ".*..1....",
			digitIndex:      4,
			rowLength:       3,
		}, output: true},
		{input: digitIsNextToSymbolTestCaseInput{
			schematicString: "....1....",
			digitIndex:      4,
			rowLength:       3,
		}, output: false},
		{input: digitIsNextToSymbolTestCaseInput{
			schematicString: ".....123......&",
			digitIndex:      7,
			rowLength:       5,
		}, output: false},
	}
	for _, tc := range cases {
		result := digitIsNextToSymbol(tc.input.schematicString, tc.input.digitIndex, tc.input.rowLength)
		if result != tc.output {
			t.Errorf(`%s("%s") = %t, wanted %t`, "digitIsNextToSymbol", tc.input.schematicString, result, tc.output)
		}
	}
}

type isSpecialSymbolTestCase struct {
	input  rune
	output bool
}

func TestIsSpecialSymbol(t *testing.T) {
	cases := []isSpecialSymbolTestCase{
		{input: []rune(".")[0], output: false},
		{input: []rune("4")[0], output: false},
		{input: []rune("$")[0], output: true},
	}
	for _, tc := range cases {
		result := isSpecialSymbol(tc.input)
		if result != tc.output {
			t.Errorf(`%s("%c") = %t, wanted %t`, "isSpecialSymbol", tc.input, result, tc.output)
		}
	}
}

type calculatePartNumberSumTestCase struct {
	input  []string
	output int
}

func TestCaclulatePartNumberSum(t *testing.T) {
	cases := []calculatePartNumberSumTestCase{
		{input: []string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		},
			output: 4361},
	}
	for _, tc := range cases {
		result := calculatePartNumberSum(tc.input)
		if result != tc.output {
			t.Errorf("calculatePartNumberSum(%s) = %d, wanted %d", tc.input, result, tc.output)
		}
	}
}
