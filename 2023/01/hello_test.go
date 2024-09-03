package main

import (
	"testing"
)

type testCase struct {
	input  string
	output int
}

func TestStringToInt(t *testing.T) {
	cases := []testCase{
		{input: "one", output: 1},
		{input: "1", output: 1},
		{input: "two", output: 2},
		{input: "2", output: 2},
		{input: "three", output: 3},
		{input: "3", output: 3},
		{input: "four", output: 4},
		{input: "asdfasdf", output: -1},
	}
	for _, element := range cases {
		want := element.output
		result := stringToInt(element.input)
		if want != result {
			t.Fatalf(`stringToInt("%s") = %d, wanted %d`, element.input, result, want)
		}
	}
}

func TestGetLastDigitPart2(t *testing.T) {
	want := 1
	result := getLastDigit2("abc1", 1)
	if want != result {
		t.Fatalf(`getLastDigitPart2("%s", 1) = %d, wanted %d`, "abc1", result, want)
	}
}

func TestGetLastDigitPart2_2(t *testing.T) {
	want := 1
	result := getLastDigit2("abcone", 1)
	if want != result {
		t.Fatalf(`getLastDigitPart2("%s", 1) = %d, wanted %d`, "abc1", result, want)
	}
}

func TestGetLastDigitPart2_3(t *testing.T) {
	want := 8
	result := getLastDigit2("oneight", 1)
	if want != result {
		t.Fatalf(`getLastDigitPart2("%s", 1) = %d, wanted %d`, "abc1", result, want)
	}
}

func TestGetLastDigitPart2_4(t *testing.T) {
	want := 3
	result := getLastDigit2("asdf3asdf", 1)
	if want != result {
		t.Fatalf(`getLastDigitPart2("%s", 1) = %d, wanted %d`, "abc1", result, want)
	}
}

func TestFirstDigit2(t *testing.T) {
	cases := []testCase{
		{input: "two1nine", output: 29},
		{input: "eightwothree", output: 83},
		{input: "abcone2threexyz", output: 13},
		{input: "xtwone3four", output: 24},
		{input: "4nineeightseven2", output: 42},
		{input: "zoneight234", output: 14},
		{input: "7pqrstsixteen", output: 76},
		{input: "2oneight", output: 28},
		{input: "oneight", output: 18},
		{input: "eighthree", output: 83},
	}
	for _, element := range cases {
		want := element.output / 10
		result := getFirstDigit2(element.input)
		want2 := element.output % 10
		result2 := getLastDigit2(element.input, 1)
		if want != result {
			t.Fatalf(`getFirstDigit2("%s") = %d, wanted %d`, element.input, result, want)
		}
		if want2 != result2 {
			t.Fatalf(`getLastDigit2("%s") = %d, wanted %d`, element.input, result2, want2)
		}
	}
}
