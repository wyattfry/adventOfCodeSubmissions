package day08

import (
	"aoc/common"
	"fmt"
	"regexp"
	"slices"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	// part2Result := calculatePart2(lines)
	// fmt.Println("Solution for part 2:", part2Result, file)
}

func calculatePart1(lines []string) int {
	instructions, network := parseInput(lines)

	return len(traversGraph(network, instructions))
}

// func calculatePart2(lines []string) int {
// 	var totalWinnings int
// 	hands := parseInput(lines)
// 	sortHands(hands, false)
// 	for index, hand := range hands {
// 		// fmt.Println(hand)
// 		totalWinnings += hand.bid * (index + 1)
// 	}

// 	return totalWinnings
// }

type node struct {
	name  string
	left  *node
	right *node
}

func parseInput(lines []string) (instructions string, root node) {
	networkMap := map[string]*node{}

	for index, line := range lines {
		if index < 2 {
			continue
		}
		matches := regexp.MustCompile(`\w+`).FindAllString(line, -1)
		slices.Reverse(matches)
		for index, name := range matches {
			_, exists := networkMap[name]
			if !exists {
				newNode := node{
					name: name,
				}
				networkMap[name] = &newNode
			}
			if index == 2 {
				networkMap[name].left = networkMap[matches[1]]
				networkMap[name].right = networkMap[matches[0]]
			}
		}
	}

	return lines[0], *networkMap["AAA"]
}

func traversGraph(network node, instructions string) []string {
	currentNode := network
	var path []string
	var index int
	for ok := true; ok; ok = (currentNode.name != "ZZZ") {
		path = append(path, currentNode.name)
		if instructions[index%len(instructions)] == 'L' {
			currentNode = *currentNode.left
		} else {
			currentNode = *currentNode.right
		}
		index += 1
	}
	return path
}
