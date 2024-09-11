package day07

import (
	"reflect"
	"strings"
	"testing"
)

var (
	example = strings.Split(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`, "\n")
)

func Test_cardHandStringToNumber(t *testing.T) {
	cases := map[string]int{
		"22222": 0x22222,
		"TTTTT": 0xAAAAA,
		"32T3K": 0x32A3D,
	}
	for input, expect := range cases {
		result := cardHandStringToNumber(input)
		if result != expect {
			t.Errorf(`f(%s) = 0x%x, but wanted 0x%x`, input, result, expect)
		}
	}
}

func Test_parseInput(t *testing.T) {
	hands := parseInput(example)
	expects := []cardHand{
		{
			bid:            765,
			cardHandString: "32T3K",
			handType:       onePair,
		},
		{
			bid:            684,
			cardHandString: "T55J5",
			handType:       threeOfAKind,
		},
	}
	for index, expect := range expects {
		if !reflect.DeepEqual(hands[index], expect) {
			t.Errorf(`f() = %#v but wanted %#v`, hands[index], expect)
		}
	}
}

func Test_sortHands(t *testing.T) {
	expect := []cardHand{
		{bid: 765, cardHandString: "32T3K", handType: 1},
		{bid: 220, cardHandString: "KTJJT", handType: 2},
		{bid: 28, cardHandString: "KK677", handType: 2},
		{bid: 684, cardHandString: "T55J5", handType: 3},
		{bid: 483, cardHandString: "QQQJA", handType: 3},
	}
	hands := parseInput(example)
	sortHands(hands)
	if !reflect.DeepEqual(hands, expect) {
		t.Errorf(`f() = %#v but wanted %#v`, hands, expect)
	}
}

func Test_calculateHandType(t *testing.T) {
	cases := map[string]cardHandType{
		"AAAAA": fiveOfAKind,
		"AA8AA": fourOfAKind,
		"23332": fullHouse,
		"TTT98": threeOfAKind,
		"23432": twoPair,
		"A23A4": onePair,
		"23456": highCard,
	}
	for input, expect := range cases {
		result := calculateHandType(input)
		if result != expect {
			t.Errorf(`f() = %#v but wanted %#v`, result, expect)
		}
	}
}

func Test_calculatePart1(t *testing.T) {
	expect := 765*1 + 220*2 + 28*3 + 684*4 + 483*5 // = 6440
	result := calculatePart1(example)
	if result != expect {
		t.Error("calculatePart1 =", result, ", but wanted", expect)
	}
}

// func Test_calculatePart2(t *testing.T) {
// 	expect := 71503
// 	result := calculatePart2(example)
// 	if result != expect {
// 		t.Error("calculatePart2 =", result, "but wanted", expect)
// 	}
// }
