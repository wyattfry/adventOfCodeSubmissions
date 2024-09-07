package templateday01

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve() {
	lines := readlines("./day01/input.txt")

	// Your Solution Here
	// `lines` is a string array of the input text file

	fmt.Println(lines)
}

func readlines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
