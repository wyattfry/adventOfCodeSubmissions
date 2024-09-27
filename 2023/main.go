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
	"aoc/day09"
	"aoc/day10"
	"aoc/day11"
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
		os.Stderr.WriteString("Day must be expressed as a integer\n")
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
		day09.Solve,
		day10.Solve,
		day11.Solve,
	}

	if len(f) < i {
		os.Stderr.WriteString(fmt.Sprintf("Day must be expressed as a integer from 1 to %d\n", i-1))
		os.Exit(1)
	}

	f[i-1](inputFile)
}
