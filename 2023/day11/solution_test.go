package day11

import (
	"aoc/common"
	"strings"
	"testing"
)

func Test_addGalaxy(t *testing.T) {
	sut := galaxyCluster{}
	sut.addGalaxy(location{
		row: 0,
		col: 0,
	})
	result := sut.idToLocation
	want := map[int]location{
		0: {0, 0},
	}
	common.AssertEqual(want, result, t)
}

func Test_parseInput(t *testing.T) {
	sut := parseInput(strings.Split(`#...
...
.#.`, "\n"))

	common.AssertEqual(map[int]location{
		0: {row: 0, col: 0},
		1: {row: 2, col: 1},
	}, sut.idToLocation, t)

	common.AssertEqual(map[int][]int{
		0: {0},
		1: {1},
	}, sut.colGalaxyIDs, t)

	common.AssertEqual(map[int][]int{
		0: {0},
		2: {1},
	}, sut.rowGalaxyIDs, t)

	common.AssertEqual(2, sut.highestRow, t)
	common.AssertEqual(1, sut.highestCol, t)
}

func Test_expand(t *testing.T) {
	sut := parseInput(strings.Split(`#...
...
#..
..#`, "\n"))
	sut.expand()
	result := sut.idToLocation
	want := map[int]location{
		0: {0, 0},
		1: {3, 0},
		2: {4, 3},
	}
	common.AssertEqual(want, result, t)
}

func Test_getDistanceBetweenGalaxies(t *testing.T) {
	sut := parseInput(strings.Split(`#...
...
#..
..#`, "\n"))
	testCases := [][]int{
		// ida, idb, want
		{0, 1, 2},
		{0, 2, 5},
		{1, 2, 3},
	}
	for _, tc := range testCases {
		got1 := sut.getDistanceBetweenGalaxies(tc[0], tc[1])
		got2 := sut.getDistanceBetweenGalaxies(tc[1], tc[0])
		common.AssertEqual(tc[2], got1, t)
		common.AssertEqual(tc[2], got2, t)
	}
}
