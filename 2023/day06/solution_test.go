package day06

import (
	"strings"
	"testing"
)

var (
	example = strings.Split(`Time:      7  15   30
	Distance:  9  40  200`, "\n")
)

func Test_parseInput(t *testing.T) {
	records := parseInput(example)
	expects := [][]int{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	for index, record := range records {
		if record.time != expects[index][0] {
			t.Error("Test_parseInput, expected time", expects[index][0], "got", record.time)
		}
		if record.distance != expects[index][1] {
			t.Error("Test_parseInput, expected distance", expects[index][0], "got", record.distance)
		}
	}

}

func Test_calculateDistanceRecordButtonTimes(t *testing.T) {
	expects := [][]int{
		{1, 5},
		{3, 11},
		{10, 20},
	}
	for index, r := range parseInput(example) {
		left, right := calculateDistanceRecordButtonTimes(r)
		resultAsInts := []int{int(left), int(right)}
		if resultAsInts[0] != expects[index][0] || resultAsInts[1] != expects[index][1] {
			t.Error("calculateDistanceRecordButtonTimes", r, "=", resultAsInts, " but wanted", expects[index])
		}
	}

}

func Test_calculateWholeNumbersBetween(t *testing.T) {
	cases := map[int][]float64{
		4: {1.6, 5.3},
		8: {3.4, 11.53},
		9: {10, 20},
	}
	for expect, input := range cases {
		result := calculateWholeNumbersBetween(input[0], input[1])
		if len(result) != expect {
			t.Error("calculateWholeNumbersBetween", input, "=", result, "length", len(result), "but expected len", expect)
		}
	}
}

func Test_calculatePart1(t *testing.T) {
	expect := 288
	result := calculatePart1(example)
	if result != expect {
		t.Error("calculatePart1 =", result, ", but wanted", expect)
	}
}

func Test_parseInputPart2(t *testing.T) {
	expect := []int{71530, 940200}
	result := parseInputPart2(example)
	if result.time != expect[0] {
		t.Error("parseInputPart2 =", result.time, "but wanted", expect[0])
	}
	if result.distance != expect[1] {
		t.Error("parseInputPart2 =", result.distance, "but wanted", expect[1])
	}
}

func Test_calculatePart2(t *testing.T) {
	expect := 71503
	result := calculatePart2(example)
	if result != expect {
		t.Error("calculatePart2 =", result, "but wanted", expect)
	}
}
