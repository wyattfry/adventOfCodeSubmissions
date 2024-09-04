package day02

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

func TestGetHighestForColorRed(t *testing.T) {
	cases := []testCase{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", output: 4},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", output: 1},
	}
	runTestCases(t, "stringToInt", cases, func(s string) int {
		return getHighestForColor(s, "red")
	})
}

func TestGetHighestForColorBlue(t *testing.T) {
	cases := []testCase{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", output: 6},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", output: 4},
	}
	runTestCases(t, "stringToInt", cases, func(s string) int {
		return getHighestForColor(s, "blue")
	})
}

func TestGetHighestForColorGreen(t *testing.T) {
	cases := []testCase{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", output: 2},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", output: 3},
	}
	runTestCases(t, "stringToInt", cases, func(s string) int {
		return getHighestForColor(s, "green")
	})
}

func TestGetPowerOfSet(t *testing.T) {
	cases := []testCase{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", output: 4 * 2 * 6},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", output: 4 * 3 * 1},
	}
	runTestCases(t, "stringToInt", cases, getPowerOfSet)
}
