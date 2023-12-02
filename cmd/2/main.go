package main

import (
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

type Round map[string]int

type Game []Round

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	games := readGames()

	sum := 0
	fewestCubesSum := 0

	for i, game := range games {

		maxAmount := make(map[string]int)
		for _, round := range game {
			for color, amount := range round {
				maxAmount[color] = max(maxAmount[color], amount)
			}
		}

		if isMapGE(limits, maxAmount) {
			sum += i + 1
		}

		minSet := 1
		for _, amount := range maxAmount {
			minSet *= amount
		}

		fewestCubesSum += minSet
	}

	fwk.Solution(1, sum)
	fwk.Solution(2, fewestCubesSum)
}

func isMapGE(a, b map[string]int) bool {
	for k, v := range b {
		if a[k] < v {
			return false
		}
	}

	return true
}

func readGames() []Game {
	lines := fwk.ReadInputLines()
	games := make([]Game, len(lines))

	for i, line := range lines {
		data := strings.Split(line, ": ")
		rounds := strings.Split(data[1], "; ")

		games[i] = make(Game, len(rounds))

		for j, roundStrs := range rounds {
			round := make(Round)

			for _, subsetStrs := range strings.Split(roundStrs, ", ") {
				dataStrs := strings.Split(subsetStrs, " ")
				round[dataStrs[1]], _ = strconv.Atoi(dataStrs[0])
			}

			games[i][j] = round
		}
	}

	return games
}
