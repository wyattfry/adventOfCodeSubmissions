package common

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"testing"
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

func ExtractInts(strWithInts string) []int {
	var output []int
	for _, numstr := range regexp.MustCompile(`-?\d+`).FindAllString(strWithInts, -1) {
		num, err := strconv.Atoi(numstr)
		if err != nil {
			panic(err)
		}
		output = append(output, num)
	}

	return output
}

func AssertEqual(want any, got any, t *testing.T) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf(`Got %#v, but wanted %#v`, got, want)
	}
}

func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
