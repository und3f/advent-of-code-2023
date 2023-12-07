package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

var cardsRank = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func main() {
	players := readInput()

	fwk.Solution(1, calTotalWinnings(players))

	cardsRank['J'] = -1
	for i := range players {
		players[i].combination = evaluateHandWithJokers(players[i].hand)
	}
	fwk.Solution(2, calTotalWinnings(players))
}

func calTotalWinnings(players []Player) uint64 {
	slices.SortFunc(players, func(a, b Player) int {
		combDiff := a.combination - b.combination
		if combDiff != 0 {
			return combDiff
		}

		for i := range a.hand {
			diff := cardsRank[a.hand[i]] - cardsRank[b.hand[i]]
			if diff != 0 {
				return diff
			}
		}

		return 0
	})

	res := uint64(0)
	for i, player := range players {
		// fmt.Printf("Hand %s rank: %d, bid: %d\n", string(player.hand), i+1, player.bid)
		res += (uint64(i) + 1) * player.bid
	}

	return res

}

type Player struct {
	hand        []rune
	bid         uint64
	combination int
}

func readInput() []Player {
	var players []Player
	for _, line := range fwk.ReadInputLines() {
		strs := strings.Split(line, " ")
		bid, _ := strconv.ParseUint(strs[1], 10, 64)
		hand := []rune(strs[0])
		combination := evaluateHand(hand)
		// sortHand(hand)
		players = append(players, Player{
			hand:        hand,
			bid:         bid,
			combination: combination,
		})
	}

	return players
}

const FiveOfAKind = 7
const FourOfAKind = FiveOfAKind - 1
const FullHouse = FourOfAKind - 1
const ThreeOfAKind = FullHouse - 1
const TwoPair = ThreeOfAKind - 1
const OnePair = TwoPair - 1

func sortHand(hand []rune) {
	cardCount := make(map[rune]int)

	for _, card := range hand {
		cardCount[card]++
	}
	slices.SortFunc(hand, func(b, a rune) int {
		countDiff := cardCount[a] - cardCount[b]
		if countDiff != 0 {
			return countDiff
		}

		return cardsRank[a] - cardsRank[b]
	})
}

func evaluateHandWithJokers(hand []rune) int {
	cardCount := make(map[rune]int)

	for _, card := range hand {
		cardCount[card]++
	}

	jokers := cardCount['J']
	delete(cardCount, 'J')

	combinations := make([]int, len(hand)+1)

	for _, v := range cardCount {
		combinations[v]++
	}

	if jokers == len(hand) {
		combinations[jokers] = 1
	} else {
		for i := len(combinations) - 1; i > 0; i-- {
			if combinations[i] > 0 {
				combinations[i]--
				combinations[i+jokers]++
				break
			}
		}
	}

	return evaluateCombinations(combinations)
}

func evaluateHand(hand []rune) int {
	cardCount := make(map[rune]int)

	for _, card := range hand {
		cardCount[card]++
	}

	combinations := make([]int, len(hand)+1)

	for _, v := range cardCount {
		combinations[v]++
	}

	return evaluateCombinations(combinations)
}

func evaluateCombinations(combinations []int) int {
	if combinations[5] > 0 {
		return FiveOfAKind
	}
	if combinations[4] > 0 {
		return FourOfAKind
	}
	if combinations[3] == 1 && combinations[2] == 1 {
		return FullHouse
	}
	if combinations[3] == 1 {
		return ThreeOfAKind
	}
	if combinations[2] == 2 {
		return TwoPair
	}
	if combinations[2] == 1 {
		return OnePair
	}

	return 0
}
