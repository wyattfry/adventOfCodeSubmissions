package day07

import (
	"aoc/common"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func Solve(file string) {
	lines := common.Readlines(file)

	part1Result := calculatePart1(lines)
	fmt.Println("Solution for part 1:", part1Result, file)

	part2Result := calculatePart2(lines)
	fmt.Println("Solution for part 2:", part2Result, file)
}

func compareCardHands(a, b cardHand) int {
	if int(a.handType) != int(b.handType) {
		c := cmp.Compare(int(a.handType), int(b.handType))
		return c
	}
	an := cardHandStringToNumber(a.cardHandString)
	bn := cardHandStringToNumber(b.cardHandString)
	c := cmp.Compare(an, bn)
	return c
}

func sortHands(hands []cardHand) {
	slices.SortFunc(hands, compareCardHands)
}

func calculatePart1(lines []string) int {
	var totalWinnings int
	hands := parseInput(lines)
	sortHands(hands)
	for index, hand := range hands {
		totalWinnings += hand.bid * (index + 1)
	}

	return totalWinnings
}

func calculatePart2(lines []string) int {
	return -1
}

type cardHand struct {
	bid            int
	cardHandString string
	handType       cardHandType
}

// Converts a cardhand string, like "QKA23" to a hexidecimal integer where the
// cards with letter names (T,J,Q,K,A) become the hex digits A-E
func cardHandStringToNumber(c string) int {
	letterNumberMap := map[rune]int{
		'T': 0xA,
		'J': 0xB,
		'Q': 0xC,
		'K': 0xD,
		'A': 0xE,
	}
	var cardHandNumber int
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
