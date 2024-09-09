package main

import (
	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
	"aoc/day05"
	"fmt"
	"os"
)

func main() {
	a := os.Args[1]
	if a == "1" {
		day01.Solve("./day01/input.txt")
	} else if a == "2" {
		day02.Solve("./day02/input.txt")
	} else if a == "3" {
		day03.Solve("./day03/input.txt")
	} else if a == "4" {
		day04.Solve("./day04/input.txt")
	} else if a == "5" {
		day05.Solve("./day05/input.txt")
	} else {
		fmt.Println("Please specify a day")
	}

}
