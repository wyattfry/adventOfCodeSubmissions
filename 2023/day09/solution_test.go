package day09

import (
	"aoc/common"
	"reflect"
	"testing"
)

type testCase struct {
	line                  string
	extrapolatedNextValue int
	sequence              []int
}

var (
	exampleCases = []testCase{
		{
			line:                  "0   3   6   9  12  15",
			extrapolatedNextValue: 18,
			sequence:              []int{0, 3, 6, 9, 12, 15},
		},
		{
			line:                  "1   3   6  10  15  21",
			extrapolatedNextValue: 28,
			sequence:              []int{1, 3, 6, 10, 15, 21},
		},
		{
			line:                  "10  13  16  21  30  45",
			extrapolatedNextValue: 68,
			sequence:              []int{10, 13, 16, 21, 30, 45},
		},
	}
)

func Test_parseInput(t *testing.T) {
	lines := []string{}
	for _, ec := range exampleCases {
		lines = append(lines, ec.line)
	}
	var i int
	for sequence := range parseInput(lines) {
		if !reflect.DeepEqual(sequence, exampleCases[i].sequence) {
			t.Errorf(`=%v but wanted %v`, sequence, exampleCases[i].sequence)
		}
		i++
	}

}

func Test_calculatePart1(t *testing.T) {
	lines := []string{}
	var sum int
	for _, ec := range exampleCases {
		lines = append(lines, ec.line)
		sum += ec.extrapolatedNextValue
	}

	result := calculatePart1(lines)
	if result != sum {
		t.Errorf(`=%d but wanted %d`, result, sum)
	}

}

type calcDifferencesCase struct {
	input []int
	want  []int
}

func Test_calcDifferences(t *testing.T) {
	testCases := []calcDifferencesCase{
		{
			input: []int{0, 3, 6, 9, 12, 15},
			want:  []int{3, 3, 3, 3, 3},
		},
		{
			input: []int{1, 3, 6, 10, 15, 21},
			want:  []int{2, 3, 4, 5, 6},
		},
	}
	for _, tc := range testCases {
		got := calcDifferences(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(`got %v, but wanted %v`, got, tc.want)
		}
	}
}

func Test_calcDifferencesRecurse(t *testing.T) {
	want := [][]int{{3, 3, 3, 3, 3}, {0, 0, 0, 0}}
	got := calcDifferencesRecurse(exampleCases[0].sequence)
	common.AssertEqual(want, got, t)
}

func Test_extrapolate(t *testing.T) {
	for _, tc := range exampleCases {
		got := extrapolate(tc.sequence)

		common.AssertEqual(tc.extrapolatedNextValue, got, t)
	}
}

// func Test_calculatePart2(t *testing.T) {
// 	for _, tc := range exampleCasesPart2 {

// 		result := calculatePart2(tc.lines)
// 		if result != tc.solution {
// 			t.Error("f() =", result, "but wanted", tc.solution)
// 		}
// 	}
// }
