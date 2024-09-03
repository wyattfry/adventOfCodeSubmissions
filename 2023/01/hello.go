package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func numberWords() []string {
	return []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
}

func main() {
	// solve("./input-test.txt")
	solve("./input.txt")

	// part 2: 52844 is too high
}

func solve(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	var sumPart2 int
	for scanner.Scan() {
		line := scanner.Text()
		value := getFirstDigit(line)*10 + getLastDigit(line)
		sum += value

		calNumber := getFirstDigit2(line)*10 + getLastDigit2(line, 1)
		// fmt.Println(calNumber)
		sumPart2 += calNumber
		// fmt.Print(sumPart2, " ")
	}
	fmt.Println("Solution Part 1:", sum, "File:", fileName)
	fmt.Println("Solution Part 2:", sumPart2, "File:", fileName)
}

func stringToInt(str string) int {
	// fmt.Println("strotint input", str)
	i, err := strconv.Atoi(str)
	if err != nil {
		result := slices.Index(numberWords(), str) + 1
		if result != 0 {
			// fmt.Println("parsed", str, "to", result)
			return result
		} else {
			// fmt.Println("couldn't parse", str, "into an integer")
			return -1
		}
	}
	return i
}

func getFirstDigit(input string) int {
	re := regexp.MustCompile(`([a-z]*)(\d)`)
	matches := re.FindStringSubmatch(input)
	if matches != nil {
		return stringToInt(matches[2])
	} else {
		return -1
	}
}

func getLastDigit(input string) int {
	re := regexp.MustCompile(`(\d)([a-z]*)$`)
	matches := re.FindStringSubmatch(input)
	if matches != nil {
		return stringToInt(matches[1])
	} else {
		return -1
	}
}

func getFirstDigit2(input string) int {
	// go does not support positive lookahead: https://github.com/google/re2/wiki/Syntax
	re := regexp.MustCompile(`(=?(one|two|three|four|five|six|seven|eight|nine|\d))`)
	matches := re.FindAllString(input, -1)
	if matches != nil {
		return stringToInt(matches[0])
	} else {
		return -1
	}
}

func getLastDigit2(input string, length int) int {
	// fmt.Println("length", length, "    string length", len(input))
	if length > len(input) {
		// fmt.Println("HERE")
		return -1
	}
	var output int
	for output == 0 {
		toParse := input[len(input)-length:]
		// fmt.Println("toParse", toParse)
		output = getFirstDigit2(toParse)
		// fmt.Println("output", output)
		if output > 0 {
			return output
		} else {
			return getLastDigit2(input, length+1)
		}
	}
	return output
}
