package day08

import (
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
)

func Test_parseInput(t *testing.T) {
	for _, tc := range exampleCases {

		instructions, network := parseInput(tc.lines)
		if instructions != tc.instructions {
			t.Errorf(`instructions = "%s" but wanted "%s"`, instructions, tc.instructions)
		}
		if network.left == nil || network.right == nil {
			t.Errorf(`network L&R nodes (%v, %v) contains null but shouldn't`, network.left, network.right)
		}
	}
}

func Test_traverseGraph(t *testing.T) {
	for _, tc := range exampleCases {

		instructions, network := parseInput(tc.lines)
		path := traversGraph(network, instructions)
		if len(path) == 0 {
			t.Errorf(`len(path) = 0 but wanted > 0`)
		}
		if network.left == nil || network.right == nil {
			t.Errorf(`network L&R nodes (%v, %v) contains null but shouldn't`, network.left, network.right)
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

//	func Test_calculatePart2(t *testing.T) {
//		expect := 5905
//		result := calculatePart2(example)
//		if result != expect {
//			t.Error("f() =", result, "but wanted", expect)
//		}
//	}
