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

	part2Result := calculatePart2(lines, true)
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

type checkpoint struct {
	Step  int
	State []string
}

func calculatePart2(lines []string, useCheckpoints bool) int {
	instructions, networkMap := parseInput(lines)
	for nodeName, pnode := range networkMap {
		if strings.HasSuffix(nodeName, "A") {
			fmt.Printf("Starting at %s\n", nodeName)
			var lastStep int
			for step, n := range traverseGraphIter(*pnode, instructions) {
				if step > 999999 {
					break
				}
				if strings.HasSuffix(n.name, "Z") {
					fmt.Printf("Step %d, diff: %d  Node %s\n", step, step-lastStep, n.name)
					lastStep = step
					break
				}
			}
		}
	}
	return -1
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
