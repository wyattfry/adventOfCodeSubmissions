package day08

import (
	"aoc/common"
	"encoding/json"
	"fmt"
	"iter"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
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

func newProgressBar(max int) *progressbar.ProgressBar {
	return progressbar.NewOptions(max,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetDescription("Traversing the desert..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]🐪[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
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
	var startingNodes []node
	var step int
	var startingStep int

	data, err := readJSON("checkpoint.json")
	if !useCheckpoints || err != nil {
		fmt.Println("Not using checkpoints")
		startingNodes = getStartingNodes(networkMap)
		startingStep = 0
	} else {
		for _, name := range data.State {
			startingNodes = append(startingNodes, *networkMap[name])
		}
		step = data.Step
		startingStep = step
		fmt.Printf("Resuming from step %d\n", step)
	}

	ITER_LIMIT := int(math.Pow(2, 30)) // 1,073,741,824
	bar := newProgressBar(ITER_LIMIT)

	// infinite loop until all end with 'Z'
	for ok := true; ok; ok = true {
		// for _, n := range startingNodes {
		// 	fmt.Print(n.name + " ")
		// }
		// fmt.Println(step)

		if step-startingStep > ITER_LIMIT {
			bar.Finish()
			os.Stderr.WriteString("\nERROR: Hit max iterations allowed\n")
			var currentNodesNames []string
			for _, n := range startingNodes {
				currentNodesNames = append(currentNodesNames, n.name)
			}
			fmt.Printf("Step: %d\nCurrent Nodes: %#v\n", step, currentNodesNames)
			jsonString, _ := json.Marshal(checkpoint{
				Step:  step,
				State: currentNodesNames,
			})
			os.WriteFile("checkpoint.json", jsonString, os.ModePerm)
			os.WriteFile(fmt.Sprintf("checkpoint.%d.json", step), jsonString, os.ModePerm)

			os.Exit(1)
		}
		nextNodes := []node{}
		var countEndWithZ int
		// iterate over current nodes
		for _, n := range startingNodes {
			i := 0
			// get each node's next node
			for nextNode := range traverseGraphIter(n, instructions, step) {
				if i == 1 {
					if strings.HasSuffix(nextNode.name, "Z") {
						countEndWithZ += 1
					}
					nextNodes = append(nextNodes, nextNode)
					break
				}
				i += 1
			}
		}
		if countEndWithZ == len(startingNodes) {
			return step + 1
		}
		startingNodes = nextNodes[:]
		step += 1
		bar.Add(1)
	}

	return -1
}

func readJSON(fileName string) (checkpoint, error) {
	datas := checkpoint{}

	file, err := os.ReadFile(fileName)
	if err != nil {
		return checkpoint{}, err
	}
	json.Unmarshal(file, &datas)

	return datas, nil
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

func traverseGraphIter(root node, instructions string, step int) iter.Seq[node] {
	return func(yield func(node) bool) {
		currentNode := root
		// var index int
		for ok := true; ok; ok = true {
			if !yield(currentNode) {
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
