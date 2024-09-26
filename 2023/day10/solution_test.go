package day10

import (
	"aoc/common"
	"strings"
	"testing"
)

func Test_indexOf(t *testing.T) {
	input := pipeSystem{
		system: []string{
			"...",
			"T..",
			"...",
		},
	}
	wantRow := 1
	wantCol := 0
	gotRow, gotCol := input.indexOf('T')
	common.AssertEqual(wantRow, gotRow, t)
	common.AssertEqual(wantCol, gotCol, t)
}

func Test_getNeighbor(t *testing.T) {
	input := pipeSystem{
		system: []string{
			"apc",
			"tea",
			"gsi",
		},
	}
	want := "past"
	got := input.getNeighbor(1, 1, "NESW")
	common.AssertEqual(want, got, t)
}

type canGoCase struct {
	system    pipeSystem
	row, col  int
	direction string
	want      bool
}

func Test_canGo(t *testing.T) {
	for _, tc := range []canGoCase{
		{
			system: pipeSystem{
				system: []string{
					"F-.",
					"|-7",
					"..|",
				},
			},
			row:       1,
			col:       1,
			direction: "E",
			want:      true,
		},
		{
			system: pipeSystem{
				system: []string{
					"F-.",
					"|-7",
					"..|",
				},
			},
			row:       1,
			col:       1,
			direction: "W",
			want:      false,
		},
		{
			system: pipeSystem{
				system: []string{
					"F-.",
					"|-7",
					"..|",
				},
			},
			row:       1,
			col:       1,
			direction: "N",
			want:      false,
		},
		{
			system: pipeSystem{
				system: []string{
					"F-.",
					"|-7",
					"..|",
				},
			},
			row:       1,
			col:       2,
			direction: "S",
			want:      true,
		},
	} {
		got := tc.system.canGo(tc.row, tc.col, tc.direction)
		common.AssertEqual(tc.want, got, t)
	}
}

type stepTestCase struct {
	sys      pipeSystem
	row, col int
	want     int
}

func Test_calcLoopLength(t *testing.T) {
	for _, tc := range []stepTestCase{
		{
			sys: pipeSystem{
				system: strings.Split(`.....
..-7.
.|.|.
.L-J.
.....`, "\n"),
			},
			row:  1,
			col:  2,
			want: 6,
		},
		{
			sys: pipeSystem{
				system: strings.Split(`.....
.S-7.
.|.|.
.L-J.
.....`, "\n"),
			},
			row:  1,
			col:  1,
			want: 7,
		},
	} {
		got := tc.sys.calcLoopLength(tc.row, tc.col)
		common.AssertEqual(tc.want, got, t)
	}
}

type calcPart1TestCase struct {
	lines []string
	want  int
}

func Test_calcpart1(t *testing.T) {
	for _, tc := range []calcPart1TestCase{
		{
			lines: strings.Split(`.....
.S-7.
.|.|.
.L-J.
.....`, "\n"),
			want: 4,
		},
		{
			lines: strings.Split(`..F7.
.FJ|.
SJ.L7
|F--J
LJ...`, "\n"),
			want: 8,
		},
	} {
		got := calculatePart1(tc.lines)
		common.AssertEqual(tc.want, got, t)
	}
}

func Test_parse(t *testing.T) {
	var input = strings.Split(`.....
.....
..F-7
..|.|
..S-J`, "\n")
	result := parseInput(input)
	path := result[string(START_TILE)].traceLoop()
	area := calculateAreaShoelace(path)
	interiorPointCount := calculateInteriorPointsPicks(area, len(path))
	common.AssertEqual(4, area, t)
	common.AssertEqual(1, interiorPointCount, t)
}
