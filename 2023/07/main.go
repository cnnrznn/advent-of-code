package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/cnnrznn/advent-of-code/util"
)

func main() {
	lines, _ := util.ReadLines("./input.txt")
	hands := makeHands(lines)

	// sort hands
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Class < hands[j].Class {
			return true
		} else if hands[i].Class > hands[j].Class {
			return false
		}

		for k := 0; k < 5; k++ {
			v1, v2 :=
				cardValPart2[hands[i].Cards[k]],
				cardValPart2[hands[j].Cards[k]]
			if v1 < v2 {
				return true
			} else if v1 > v2 {
				return false
			}
		}

		return false
	})

	fmt.Println(score(hands))
}

var (
	cardValPart2 = map[byte]int{
		'J': -1,
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'Q': 10,
		'K': 11,
		'A': 12,
	}
)

const (
	HIGHCARD     = 0
	ONEPAIR      = 1
	TWOPAIR      = 2
	THREEOFAKIND = 3
	FULLHOUSE    = 4
	FOUROFAKIND  = 5
	FIVEOFAKIND  = 6
)

type Hand struct {
	Class int
	Cards string
	Bid   int
}

func calcClassPart2(cards string) int {
	counts := make(map[rune]int)
	pairCount := make(map[int]int)

	for _, card := range cards {
		counts[card]++
		if card == 'J' {
			continue
		}

		pairCount[counts[card]]++
		pairCount[counts[card]-1]--
	}

	// Once we determine the number of pairs of other cards, find the next best pair size
	// The optimal strategy is to use the J's to increase the pair size of another pair,
	// because pair size beats all types of pairs of lesser size.
	// So, find the next best size of pairs.
	// Possibilities for existing pairs are [1, 5]
	// Subtract a pair count of the next best size, and increase the pair count of (# J's + next_best)
	nextBest := 0
	for k := 5; k > 0; k-- {
		if v, ok := pairCount[k]; ok && v > 0 {
			nextBest = k
			break
		}
	}

	pairCount[nextBest]--
	pairCount[nextBest+counts['J']]++

	for k := 5; k > 1; k-- {
		if v, ok := pairCount[k]; !ok || v == 0 {
			continue
		}

		v := pairCount[k]

		switch k {
		case 5:
			return FIVEOFAKIND
		case 4:
			return FOUROFAKIND
		case 3:
			if v, ok := pairCount[2]; ok && v > 0 {
				return FULLHOUSE
			} else {
				return THREEOFAKIND
			}
		case 2:
			if v == 2 {
				return TWOPAIR
			}
			return ONEPAIR
		}
	}

	return HIGHCARD
}
func calcClassPart1(cards string) int {
	counts := make(map[rune]int)
	pairCount := make(map[int]int)

	for _, card := range cards {
		counts[card]++
		pairCount[counts[card]]++
		pairCount[counts[card]-1]--
	}

	for k := 5; k > 0; k-- {
		if _, ok := pairCount[k]; !ok {
			continue
		}

		v := pairCount[k]

		switch k {
		case 5:
			return FIVEOFAKIND
		case 4:
			return FOUROFAKIND
		case 3:
			if v, ok := pairCount[2]; ok && v > 0 {
				return FULLHOUSE
			} else {
				return THREEOFAKIND
			}
		case 2:
			if v == 2 {
				return TWOPAIR
			}
			return ONEPAIR
		}
	}

	return HIGHCARD
}

func score(hands []Hand) int {
	result := 0

	for i, hand := range hands {
		result += (i + 1) * hand.Bid
	}

	return result
}

func makeHands(lines []string) []Hand {
	result := []Hand{}

	for _, line := range lines {
		fields := strings.Fields(line)
		bid, _ := strconv.Atoi(fields[1])
		cards := fields[0]

		result = append(result, Hand{
			Bid:   bid,
			Cards: cards,
			Class: calcClassPart2(cards),
		})
	}

	return result
}
