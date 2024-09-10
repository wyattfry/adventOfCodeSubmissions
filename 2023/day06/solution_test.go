package day06

import (
	"math"
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

func Test_calculateDistanceBetweenIntersections(t *testing.T) {
	expects := []int{4, 8, 9}
	for index, r := range parseInput(example) {
		resultf64 := calculateDistanceBetweenIntersections(r)
		result := int(math.Ceil(resultf64))
		if result != expects[index] {
			t.Error("calculateDistanceBetweenIntersections", r, "=", result, " but wanted", expects[index])
		}
	}

}
