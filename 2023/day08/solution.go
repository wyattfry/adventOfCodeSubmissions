package day08

import (
	"aoc/common"
	"fmt"
	"iter"
	"regexp"
	"slices"
	"strings"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
	// answer > 9,999,999
}

func calculatePart1(lines []string) int {
	instructions, network := parseInput(lines)

	startingNode := *network["AAA"]

	return len(traverseGraph(startingNode, instructions))
}

func getStartingNodes(networkMap map[string]*node) []node {
	startingNodes := []node{}
	for nodeName, pnode := range networkMap {
		// get all nodes with names that end with 'A'
		if strings.HasSuffix(nodeName, "A") {
			startingNodes = append(startingNodes, *pnode)
		}
	}
	return startingNodes
}

// Greatest Common Divisor via Euclidian Algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Least Common Multiple via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func calculatePart2(lines []string) int {
	instructions, networkMap := parseInput(lines)
	firstZSteps := []int{}
	for nodeName, pnode := range networkMap {
		if strings.HasSuffix(nodeName, "A") {
			for step, n := range traverseGraphIter(*pnode, instructions) {
				if strings.HasSuffix(n.name, "Z") {
					firstZSteps = append(firstZSteps, step)
					break
				}
			}
		}
	}
	return LCM(firstZSteps[0], firstZSteps[1], firstZSteps...)
}

type node struct {
	name  string
	left  *node
	right *node
}

func parseInput(lines []string) (instructions string, root map[string]*node) {
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

	return lines[0], networkMap
}

func traverseGraph(network node, instructions string) []string {
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

func traverseGraphIter(root node, instructions string) iter.Seq2[int, node] {
	return func(yield func(int, node) bool) {
		currentNode := root
		var step int
		for ok := true; ok; ok = true {
			if !yield(step, currentNode) {
				return
			}
			if instructions[step%len(instructions)] == 'L' {
				currentNode = *currentNode.left
			} else {
				currentNode = *currentNode.right
			}
			step += 1
		}
	}
}
