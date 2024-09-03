package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	numberWords = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	// Pre-compile regex patterns to reuse them
	firstDigitPattern = regexp.MustCompile(`([a-z]*)(\d)`)
	lastDigitPattern  = regexp.MustCompile(`(\d)([a-z]*)$`)
	anyDigitPattern   = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
)

func main() {
	// solve("./input-test.txt")
	solve("./input.txt")
}

func solve(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sumPart1, sumPart2 int

	for scanner.Scan() {
		line := scanner.Text()
		sumPart1 += calculateSum(line, getFirstDigit, getLastDigit)
		sumPart2 += calculateSum(line, getFirstMatchDigit, getLastMatchDigit)
	}

	fmt.Printf("Solution Part 1: %d (File: %s)\n", sumPart1, fileName)
	fmt.Printf("Solution Part 2: %d (File: %s)\n", sumPart2, fileName)
}

func calculateSum(input string, firstDigitFunc, lastDigitFunc func(string) int) int {
	first := firstDigitFunc(input)
	last := lastDigitFunc(input)
	if first == -1 || last == -1 {
		return 0
	}
	return first*10 + last
}

func stringToInt(str string) int {
	if num, err := strconv.Atoi(str); err == nil {
		return num
	}
	if num, found := numberWords[str]; found {
		return num
	}
	return -1
}

func getFirstDigit(input string) int {
	matches := firstDigitPattern.FindStringSubmatch(input)
	if matches != nil {
		return stringToInt(matches[2])
	}
	return -1
}

func getLastDigit(input string) int {
	matches := lastDigitPattern.FindStringSubmatch(input)
	if matches != nil {
		return stringToInt(matches[1])
	}
	return -1
}

func getFirstMatchDigit(input string) int {
	matches := anyDigitPattern.FindString(input)
	return stringToInt(matches)
}

func getLastMatchDigit(input string) int {
	for length := 1; length <= len(input); length++ {
		substr := input[len(input)-length:]
		if digit := getFirstMatchDigit(substr); digit > 0 {
			return digit
		}
	}
	return -1
}
