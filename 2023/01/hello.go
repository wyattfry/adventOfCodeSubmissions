package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	solve("./input-test.txt")
	solve("./input.txt")
}

func solve(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		value := getFirstDigit(line)*10 + getLastDigit(line)
		sum += value
	}
	fmt.Println("Solution:", sum, "File:", fileName)
}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return -1
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
