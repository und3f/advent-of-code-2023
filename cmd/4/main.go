package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	in := readInput()

	part1(in)
	part2(in)
}

func part1(in [][2][]int) {
	sum := 0
	for _, card := range in {
		sum += calWinningNumbers(card[0], card[1])
	}
	fwk.Solution(1, sum)
}

func part2(in [][2][]int) {
	instances := make([]int, len(in))

	var pile []int
	for i := range in {
		pile = append(pile, i)
	}

	for len(pile) > 0 {
		i := pile[0]
		card := in[i]
		pile = pile[1:]

		instances[i]++

		num := len(findWinningNumbers(card[0], card[1]))

		for j := i + 1; j < i+1+num; j++ {
			pile = append(pile, j)
		}
	}

	sum := 0
	for _, count := range instances {
		sum += count
	}

	fwk.Solution(2, sum)
}

var numbersSplitRe = regexp.MustCompile("\\s+")

func readInput() [][2][]int {
	lines := fwk.ReadInputLines()
	cards := make([][2][]int, len(lines))

	for i, line := range lines {
		cardStrs := strings.Split(line, ": ")
		numbersStrs := strings.Split(cardStrs[1], " | ")

		var wins []int
		for _, numStr := range numbersSplitRe.Split(numbersStrs[0], -1) {
			if len(numStr) > 0 {
				num, _ := strconv.Atoi(numStr)
				wins = append(wins, num)
			}
		}

		var myNums []int
		for _, numStr := range numbersSplitRe.Split(numbersStrs[1], -1) {
			if len(numStr) > 0 {
				num, _ := strconv.Atoi(numStr)
				myNums = append(myNums, num)
			}
		}

		cards[i] = [2][]int{wins, myNums}
	}

	return cards
}

func calWinningNumbers(wins, owns []int) int {
	won := findWinningNumbers(wins, owns)
	if len(won) == 0 {
		return 0
	}

	count := 1
	for i := 1; i < len(won); i++ {
		count = count * 2
	}

	return count
}
func findWinningNumbers(wins, owns []int) []int {
	var won []int
	for _, win := range wins {
		if slices.Index(owns, win) >= 0 {
			won = append(won, win)
		}
	}

	return won
}
