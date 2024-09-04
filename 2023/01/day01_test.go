package day01

import (
	"testing"
)

type testCase struct {
	input  string
	output int
}

// Helper function to run test cases
func runTestCases(t *testing.T, testName string, cases []testCase, testFunc func(string) int) {
	for _, tc := range cases {
		t.Run(testName+"_"+tc.input, func(t *testing.T) {
			result := testFunc(tc.input)
			if tc.output != result {
				t.Errorf(`%s("%s") = %d, wanted %d`, testName, tc.input, result, tc.output)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	cases := []testCase{
		{input: "one", output: 1},
		{input: "1", output: 1},
		{input: "two", output: 2},
		{input: "2", output: 2},
		{input: "three", output: 3},
		{input: "3", output: 3},
		{input: "four", output: 4},
		{input: "asdfasdf", output: -1},
	}
	runTestCases(t, "stringToInt", cases, stringToInt)
}

func TestGetLastMatchDigit(t *testing.T) {
	cases := []testCase{
		{input: "abc1", output: 1},
		{input: "abcone", output: 1},
		{input: "oneight", output: 8},
		{input: "asdf3asdf", output: 3},
	}
	runTestCases(t, "getLastMatchDigit", cases, getLastMatchDigit)
}

func TestGetFirstAndLastMatchDigit(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{input: "two1nine", output: 29},
		{input: "eightwothree", output: 83},
		{input: "abcone2threexyz", output: 13},
		{input: "xtwone3four", output: 24},
		{input: "4nineeightseven2", output: 42},
		{input: "zoneight234", output: 14},
		{input: "7pqrstsixteen", output: 76},
		{input: "2oneight", output: 28},
		{input: "oneight", output: 18},
		{input: "eighthree", output: 83},
	}

	for _, tc := range cases {
		t.Run("getFirstAndLastMatchDigit_"+tc.input, func(t *testing.T) {
			firstDigit := tc.output / 10
			lastDigit := tc.output % 10

			resultFirst := getFirstMatchDigit(tc.input)
			resultLast := getLastMatchDigit(tc.input)

			if resultFirst != firstDigit {
				t.Errorf(`getFirstMatchDigit("%s") = %d, wanted %d`, tc.input, resultFirst, firstDigit)
			}
			if resultLast != lastDigit {
				t.Errorf(`getLastMatchDigit("%s") = %d, wanted %d`, tc.input, resultLast, lastDigit)
			}
		})
	}
}
