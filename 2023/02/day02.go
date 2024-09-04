package day02

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
)

// type cubeset struct {
// 	red   int
// 	green int
// 	blue  int
// }

// type game struct {
// 	id   int
// 	sets []cubeset
// }

// func parseGame(gamestring string) game {
// 	// Game 94: 14 red, 10 green; 15 red; 4 red; 4 green, 7 red, 1 blue; 6 red, 5 green; 1 red, 2 green
// 	idstring := regexp.MustCompile(`^Game (\d+):`).FindString(gamestring)
// 	id, _ := strconv.Atoi(idstring)

// 	sets :=

// 	for _, set := range strings.Split(strings.SplitAfterN(gamestring, ":", strings.Index(gamestring, ":"))[1], ";") {

// 		red, _ := strconv.Atoi(regexp.MustCompile(`(\d+) red`).FindString(set))
// 		blue, _ := strconv.Atoi(regexp.MustCompile(`(\d+) blue`).FindString(set))
// 		green, _ := strconv.Atoi(regexp.MustCompile(`(\d+) green`).FindString(set))

// 		sets.append(cubeset{
// 			red:   red,
// 			blue:  blue,
// 			green: green,
// 		})

// 	}
// 	parsedGame := game{
// 		id:   id,
// 		sets: sets,
// 	}

// 	return parsedGame
// }

// Determine which games would have been possible if the bag had been loaded
// with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of
// the IDs of those games?

var (
	redCount   = 12
	greenCount = 13
	blueCount  = 14
)

func Solve(file string) {
	var gameIdSum int
	for index, line := range common.Readlines(file) {
		if getHighestForColor(line, "red") > redCount {
			continue
		} else if getHighestForColor(line, "green") > greenCount {
			continue
		} else if getHighestForColor(line, "blue") > blueCount {
			continue
		}
		gameId := index + 1
		gameIdSum += gameId
	}
	fmt.Println("Solution for part 1:", gameIdSum)
}

func getHighestForColor(gamestring string, color string) int {
	matches := regexp.MustCompile(fmt.Sprintf("(\\d+)( %s)", color)).FindAllStringSubmatch(gamestring, -1)
	var highest int
	for _, str := range matches {
		if num, err := strconv.Atoi(str[1]); err == nil {
			if num > highest {
				highest = num
			}
		}
	}
	fmt.Println(matches, "highest:", highest)

	return highest
}
