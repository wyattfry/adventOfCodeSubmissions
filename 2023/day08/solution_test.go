package day08

import (
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	lines        []string
	instructions string
	solution     int
}

var (
	exampleCases = []testCase{
		{
			lines: strings.Split(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`, "\n"),
			instructions: "RL",
			solution:     2,
		},
		{
			lines: strings.Split(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`, "\n"),
			instructions: "LLR",
			solution:     6,
		},
	}
	exampleCasesPart2 = []testCase{
		{
			lines: strings.Split(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`, "\n"),
			instructions: "LR",
			solution:     6,
		},
	}
)

func Test_parseInput(t *testing.T) {
	for _, tc := range exampleCases {

		instructions, network := parseInput(tc.lines)
		if instructions != tc.instructions {
			t.Errorf(`instructions = "%s" but wanted "%s"`, instructions, tc.instructions)
		}
		if network["AAA"].left == nil || network["AAA"].right == nil {
			t.Errorf(`network L&R nodes (%v, %v) contains null but shouldn't`, network["AAA"].left, network["AAA"].right)
		}
	}
}

func Test_traverseGraph(t *testing.T) {
	for _, tc := range exampleCases {

		instructions, network := parseInput(tc.lines)
		path := traverseGraph(*network["AAA"], instructions)
		if len(path) == 0 {
			t.Errorf(`len(path) = 0 but wanted > 0`)
		}
		node := *network["AAA"]
		if node.left == nil || node.right == nil {
			t.Errorf(`node L&R nodes (%v, %v) contains null but shouldn't`, node.left, node.right)
		}
	}
}

func Test_calculatePart1(t *testing.T) {
	for _, tc := range exampleCases {

		result := calculatePart1(tc.lines)
		if result != tc.solution {
			t.Errorf(`=%d but wanted %d`, result, tc.solution)
		}
	}
}

func Test_traverseGraphIter(t *testing.T) {
	instructions, networkMap := parseInput(exampleCases[0].lines)
	for n := range traverseGraphIter(*networkMap["AAA"], instructions, 0) {
		fmt.Println(n.name)
		if n.name == "ZZZ" {
			break
		}
	}
}

func Test_calculatePart2(t *testing.T) {
	for _, tc := range exampleCasesPart2 {

		result := calculatePart2(tc.lines, false)
		if result != tc.solution {
			t.Error("f() =", result, "but wanted", tc.solution)
		}
	}
}
