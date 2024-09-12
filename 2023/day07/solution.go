package day07

import (
	"aoc/common"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type cardHandType int

const (
	highCard cardHandType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var (
	cardNames = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
	// 251133021 is too low
}

// Using closure to make two versions of the compare function, for parts 1 and 2
func makeCompareFunc(usePart1Rules bool) func(a, b cardHand) int {

	// Using a cache was a huge performance improvement
	cache := map[string]string{}

	return func(a, b cardHand) int {

		if usePart1Rules {
			// Use Part 1 Rules, 'J's are not wild
			if int(a.handType) != int(b.handType) {
				// a & b have different hand types
				c := cmp.Compare(int(a.handType), int(b.handType))
				return c
			}
		} else {
			// Use Part 2 Rules, 'J's are wild when comparing hand types
			var strongestHands []cardHand

			for _, ch := range []cardHand{a, b} {
				val, ok := cache[ch.cardHandString]
				if ok {
					strongestHands = append(strongestHands, cardHand{
						bid:            ch.bid,
						cardHandString: val,
						handType:       calculateHandType(val),
					})
				} else {
					strongest := calculateStrongestHand(ch)
					cache[ch.cardHandString] = strongest.cardHandString
					strongestHands = append(strongestHands, strongest)
				}
			}
			if strongestHands[0].handType != strongestHands[1].handType {
				return cmp.Compare(strongestHands[0].handType, strongestHands[1].handType)
			}
		}
		an := cardHandStringToNumber(a.cardHandString, usePart1Rules)
		bn := cardHandStringToNumber(b.cardHandString, usePart1Rules)
		c := cmp.Compare(an, bn)

		return c
	}
}

// Given a card hand and 'J's are wild, return the strongest card hand possible
func calculateStrongestHand(ch cardHand) cardHand {
	if !strings.ContainsRune(ch.cardHandString, 'J') {
		return ch
	}
	hands := []cardHand{}
	for _, chs := range calculateAllHands(ch.cardHandString, 0) {
		if !strings.ContainsRune(chs, 'J') {
			hands = append(hands, cardHand{
				bid:            ch.bid,
				cardHandString: chs,
				handType:       calculateHandType(chs),
			})
		}
	}
	sortHands(hands, true)

	return hands[len(hands)-1] // or it may be the last one, idk
}

// Given a card hand string (e.g. "34JQK") and 'J's are wild, return an array of
// possible card hand strings. It uses recursion to form all possible
// combinations. I wonder if there's a more efficient way.
func calculateAllHands(cardHandString string, startIdx int) []string {
	result := []string{}
	if startIdx >= len(cardHandString) {
		return []string{cardHandString}
	}
	if cardHandString[startIdx] != 'J' {
		return calculateAllHands(cardHandString, startIdx+1)
	} else {
		for _, n := range cardNames {
			variant := cardHandString[:startIdx] + string(n) + cardHandString[startIdx+1:]
			result = append(result, calculateAllHands(variant, startIdx+1)...)
		}
	}

	return result
}

// Sorts card hands by hand type, ascending in strength
func sortHands(hands []cardHand, usePart1Rules bool) {
	compareCardHands := makeCompareFunc(usePart1Rules)
	slices.SortFunc(hands, compareCardHands)
}

func calculatePart1(lines []string) int {
	var totalWinnings int
	hands := parseInput(lines)
	sortHands(hands, true)
	for index, hand := range hands {
		totalWinnings += hand.bid * (index + 1)
	}

	return totalWinnings
}

func calculatePart2(lines []string) int {
	var totalWinnings int
	hands := parseInput(lines)
	sortHands(hands, false)
	for index, hand := range hands {
		// fmt.Println(hand)
		totalWinnings += hand.bid * (index + 1)
	}

	return totalWinnings
}

type cardHand struct {
	bid            int
	cardHandString string
	handType       cardHandType
}

// Converts a cardhand string, like "QKA23" to a hexidecimal integer where the
// cards with letter names (T,J,Q,K,A) become the hex digits A-E
func cardHandStringToNumber(c string, usePart1Rules bool) int {
	var cardHandNumber int
	letterNumberMap := map[rune]int{
		'T': 0xA,
		'J': 0xB,
		'Q': 0xC,
		'K': 0xD,
		'A': 0xE,
	}
	if !usePart1Rules {
		letterNumberMap['J'] = 0x1
	}
	for _, r := range c {
		cardHandNumber *= 0x10
		currentDigit := int(r - '0')
		if currentDigit > 1 && currentDigit < 10 {
			cardHandNumber += currentDigit
		} else {
			cardHandNumber += letterNumberMap[r]
		}
	}

	return cardHandNumber
}

func parseInput(lines []string) []cardHand {
	hands := []cardHand{}
	for _, line := range lines {
		s := strings.Split(line, " ")
		hands = append(hands, cardHand{
			bid:            common.ExtractInts(s[1])[0],
			cardHandString: s[0],
			handType:       calculateHandType(s[0]),
		})
	}
	return hands
}

func calculateHandType(cardHandString string) cardHandType {
	counts := map[rune]int{}
	for _, r := range cardHandString {
		counts[r] += 1
	}
	l := len(counts)
	switch l {
	case 1:
		// {A: 5}
		return fiveOfAKind
	case 2:
		for _, quant := range counts {
			if quant == 4 {
				// {A: 4, 8: 1}
				return fourOfAKind
			}
		}
		// {2: 2, 3: 3}
		return fullHouse
	case 3:
		for _, quant := range counts {
			if quant == 3 {
				// {T: 3, 9: 1, 8: 1}
				return threeOfAKind
			}
		}
		// {2: 2, 3: 2, 4: 1}
		return twoPair
	case 4:
		return onePair
	case 5:
		return highCard
	}
	return -1
}
