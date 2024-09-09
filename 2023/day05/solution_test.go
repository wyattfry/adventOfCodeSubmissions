package day05

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

var (
	example = strings.Split(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`, "\n")
)

func Test_parseInput(t *testing.T) {
	seedNumbers, srcDstMaps := parseInput(example)
	wantSeedNumbers := []int{13, 14, 55, 79}
	slices.Sort(seedNumbers)
	slices.Sort(wantSeedNumbers)

	if !reflect.DeepEqual(seedNumbers, wantSeedNumbers) {
		t.Error("Wrong seedNumbers")
	}

	if srcDstMaps[0][0].destRangeStart != 50 {
		t.Error("Wrong srcDstMaps, expected 50 at [0][0].destRangeStart, got", srcDstMaps[0][0])
	}

	if len(srcDstMaps) != 7 {
		t.Error("parseInput(example) expected len(srcDstMaps) = 7, but got", len(srcDstMaps))
	}
}

func Test_convertSeedToLocation(t *testing.T) {
	testCases := map[int]int{
		79: 82,
		// 14: 43,
		// 55: 86,
		// 13: 35,
	}
	_, srcDstMaps := parseInput(example)
	for input, expect := range testCases {
		result := 9999999999999999
		convertSeedToLocation(input, srcDstMaps, &result)
		if result != expect {
			t.Errorf("convertSeedToLocation(%d) = %d, but wanted %d", input, result, expect)
		}
	}
}

func Test_calculatePart1(t *testing.T) {
	result := calculatePart1(example)
	expect := 35
	if result != expect {
		t.Error("calculatePart1(example) = ", result, ", but wanted", expect)
	}
}

func Test_calculatePart2(t *testing.T) {
	result := calculatePart2(example)
	expect := 46
	if result != expect {
		t.Error("calculatePart2(example) = ", result, ", but wanted", expect)
	}
}
