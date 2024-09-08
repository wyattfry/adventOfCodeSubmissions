package day04

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var (
	example = strings.Split(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`, "\n")
)

func TestParse(t *testing.T) {
	result := parseLine(example[0])
	fmt.Println(result)

	if !reflect.DeepEqual(result.winners, []int{41, 48, 83, 86, 17}) {
		t.Error("Wrong winners")
	}

	myNumbersExpect := []int{83, 86, 6, 31, 17, 9, 48, 53}

	if !reflect.DeepEqual(result.myNumbers, myNumbersExpect) {
		t.Error("Wrong myNumbers, expected", myNumbersExpect, "got", result.myNumbers)
	}
}

func TestCalculatePart1(t *testing.T) {
	want := 13
	result := calculatePart1(example)
	if result != want {
		t.Error("calculatePart1() =", result, "but wanted", want)
	}
}

func Test_getWinnerCount(t *testing.T) {
	sc := scratchcard{
		winners:   []int{1, 2, 3},
		myNumbers: []int{1, 2, 4, 5, 6},
	}
	result := getWinnerCount(sc)
	if result != 2 {
		t.Error("getWinnerCount() =", result, " but wanted 2")
	}
}

func Test_calculatePart2(t *testing.T) {
	result := calculatePart2(example)
	want := 30
	if result != want {
		t.Error("calculatePart2() =", result, "but wanted", want)
	}
}
