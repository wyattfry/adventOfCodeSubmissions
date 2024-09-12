package main

import (
	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
	"aoc/day05"
	"aoc/day06"
	"aoc/day07"
	"aoc/day08"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Please specify a day\n")
		os.Exit(1)
	}

	i, err := strconv.Atoi(os.Args[1])

	if err != nil {
		os.Stderr.WriteString("Day must be expressed as a digit from 1 to 25\n")
		os.Exit(1)
	}

	inputFile := fmt.Sprintf("./day%02d/input.txt", i)

	f := []func(fileName string){
		day01.Solve,
		day02.Solve,
		day03.Solve,
		day04.Solve,
		day05.Solve,
		day06.Solve,
		day07.Solve,
		day08.Solve,
	}

	f[i-1](inputFile)
}
