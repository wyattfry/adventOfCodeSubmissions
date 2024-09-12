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
		result := cardHandStringToNumber(input, true)
		if result != expect {
			t.Errorf(`f(%s) = 0x%x, but wanted 0x%x`, input, result, expect)
		}
	}
}

func Test_cardHandStringToNumberPart2(t *testing.T) {
	cases := map[string]int{
		"22222": 0x22222,
		"TTTTT": 0xAAAAA,
		"32T3K": 0x32A3D,
		"JJT3K": 0x11A3D,
	}
	for input, expect := range cases {
		result := cardHandStringToNumber(input, false)
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
	sortHands(hands, true)
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

func Test_calculateAllHands(t *testing.T) {
	cases := map[string][]string{
		"J":     cardNames,
		"2J":    {"22", "23", "24", "25", "26", "27", "28", "29", "2T", "2J", "2Q", "2K", "2A"},
		"KTJJT": {"KT22T", "KT23T", "KT24T", "KT25T", "KT26T", "KT27T", "KT28T", "KT29T", "KT2TT", "KT2JT", "KT2QT", "KT2KT", "KT2AT", "KT32T", "KT33T", "KT34T", "KT35T", "KT36T", "KT37T", "KT38T", "KT39T", "KT3TT", "KT3JT", "KT3QT", "KT3KT", "KT3AT", "KT42T", "KT43T", "KT44T", "KT45T", "KT46T", "KT47T", "KT48T", "KT49T", "KT4TT", "KT4JT", "KT4QT", "KT4KT", "KT4AT", "KT52T", "KT53T", "KT54T", "KT55T", "KT56T", "KT57T", "KT58T", "KT59T", "KT5TT", "KT5JT", "KT5QT", "KT5KT", "KT5AT", "KT62T", "KT63T", "KT64T", "KT65T", "KT66T", "KT67T", "KT68T", "KT69T", "KT6TT", "KT6JT", "KT6QT", "KT6KT", "KT6AT", "KT72T", "KT73T", "KT74T", "KT75T", "KT76T", "KT77T", "KT78T", "KT79T", "KT7TT", "KT7JT", "KT7QT", "KT7KT", "KT7AT", "KT82T", "KT83T", "KT84T", "KT85T", "KT86T", "KT87T", "KT88T", "KT89T", "KT8TT", "KT8JT", "KT8QT", "KT8KT", "KT8AT", "KT92T", "KT93T", "KT94T", "KT95T", "KT96T", "KT97T", "KT98T", "KT99T", "KT9TT", "KT9JT", "KT9QT", "KT9KT", "KT9AT", "KTT2T", "KTT3T", "KTT4T", "KTT5T", "KTT6T", "KTT7T", "KTT8T", "KTT9T", "KTTTT", "KTTJT", "KTTQT", "KTTKT", "KTTAT", "KTJ2T", "KTJ3T", "KTJ4T", "KTJ5T", "KTJ6T", "KTJ7T", "KTJ8T", "KTJ9T", "KTJTT", "KTJJT", "KTJQT", "KTJKT", "KTJAT", "KTQ2T", "KTQ3T", "KTQ4T", "KTQ5T", "KTQ6T", "KTQ7T", "KTQ8T", "KTQ9T", "KTQTT", "KTQJT", "KTQQT", "KTQKT", "KTQAT", "KTK2T", "KTK3T", "KTK4T", "KTK5T", "KTK6T", "KTK7T", "KTK8T", "KTK9T", "KTKTT", "KTKJT", "KTKQT", "KTKKT", "KTKAT", "KTA2T", "KTA3T", "KTA4T", "KTA5T", "KTA6T", "KTA7T", "KTA8T", "KTA9T", "KTATT", "KTAJT", "KTAQT", "KTAKT", "KTAAT"},
	}
	for input, expect := range cases {
		result := calculateAllHands(input, 0)
		if !reflect.DeepEqual(result, expect) {
			t.Errorf(`f("%s")=%#v but wanted %#v`, input, result, expect)
		}
	}
}

func Test_calculateStrongestHand(t *testing.T) {
	cases := map[string]string{
		"QJJQ2": "QQQQ2",
	}
	for input, expect := range cases {
		result := calculateStrongestHand(cardHand{cardHandString: input})
		if result.cardHandString != expect {
			t.Errorf(`("%s")=%s but wanted %s`, input, result.cardHandString, expect)
		}
	}
}

func Test_calculatePart2(t *testing.T) {
	expect := 5905
	result := calculatePart2(example)
	if result != expect {
		t.Error("f() =", result, "but wanted", expect)
	}
}
