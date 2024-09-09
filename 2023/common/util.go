package common

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
)

func Readlines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func GetEnvNewline() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func ExtractInts(strWithInts string) []int {
	var output []int
	for _, numstr := range regexp.MustCompile(`\d+`).FindAllString(strWithInts, -1) {
		num, _ := strconv.Atoi(numstr)
		output = append(output, num)
	}

	return output
}
