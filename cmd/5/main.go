package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	seeds, maps := readInput()

	part1(seeds, maps)
	part2(seeds, maps)
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func part1(seeds []int, maps map[string]ConvertingMap) {
	lowestSeedLocation := MaxInt

	for _, seed := range seeds {
		number := convertSeed(seed, maps)
		if number < lowestSeedLocation {
			lowestSeedLocation = number
		}
	}
	fwk.Solution(1, lowestSeedLocation)
}

func part2(seeds []int, maps map[string]ConvertingMap) {
	lowestSeedLocation := MaxInt
	for i := 0; i < len(seeds); i += 2 {
		numbers := [][2]int{
			[2]int{seeds[i], seeds[i] + seeds[i+1]},
		}

		unit := "seed"
		for convertingMap, exists := maps[unit]; exists; convertingMap, exists = maps[unit] {
			var cnvNumbers [][2]int
			for _, number := range numbers {
				cnvNumbers = append(cnvNumbers, convertNumberRange(number, convertingMap.ranges)...)
			}
			numbers = cnvNumbers
			unit = convertingMap.unitDst
		}

		for _, number := range numbers {
			lowestSeedLocation = min(lowestSeedLocation, number[0])
		}
	}

	fwk.Solution(2, lowestSeedLocation)
}

func convertNumberRange(in [2]int, ranges []Range) [][2]int {
	number := in[0]
	numberLast := in[1]

	i := 0
	for ; i < len(ranges) && ranges[i].src+ranges[i].rng < number; i++ {
	}

	var cnvNumber [][2]int
	for number <= numberLast {
		if i >= len(ranges) {
			cnvNumber = append(cnvNumber, [2]int{number, numberLast})
			break
		}

		rng := ranges[i]
		if number < ranges[i].src {
			end := rng.src - 1

			if numberLast < end {
				cnvNumber = append(cnvNumber, [2]int{number, numberLast})
				break
			}

			cnvNumber = append(cnvNumber, [2]int{number, end})
			number = end + 1
		}

		end := min(numberLast, rng.src+rng.rng-1)
		offset := rng.dst - rng.src
		cnvNumber = append(cnvNumber, [2]int{number + offset, end + offset})
		number = end + 1

		i++
	}

	return cnvNumber
}

func convertSeed(seed int, maps map[string]ConvertingMap) int {

	unit := "seed"
	number := seed

	// fmt.Printf(" - %s %d", unit, number)
	for convertingMap, exists := maps[unit]; exists; convertingMap, exists = maps[unit] {
		unit = convertingMap.unitDst

		convertRng := Range{}
		for _, rng := range convertingMap.ranges {
			if rng.src <= number && number < rng.src+rng.rng {
				convertRng = rng
				break
			}
		}
		number += convertRng.dst - convertRng.src

		// fmt.Printf(", %s %d", unit, number)
	}
	// fmt.Println()

	return number
}

type Range struct {
	dst int
	src int
	rng int
}

type ConvertingMap struct {
	unitDst string
	ranges  []Range
}

func readInput() ([]int, map[string]ConvertingMap) {
	sections := strings.Split(strings.TrimSpace(fwk.ReadInput("")), "\n\n")

	seeds := parseNumbers(strings.Split(sections[0], ": ")[1])
	maps := make(map[string]ConvertingMap)

	for _, mapStr := range sections[1:] {
		lines := strings.Split(mapStr, "\n")
		name := strings.Split(lines[0], " ")[0]

		units := strings.Split(name, "-to-")
		unitSrc := units[0]
		unitDst := units[1]

		var convertingMap ConvertingMap
		convertingMap.unitDst = unitDst

		for _, mapRangeStr := range lines[1:] {
			nums := parseNumbers(mapRangeStr)

			convertingMap.ranges = append(convertingMap.ranges, Range{
				dst: nums[0],
				src: nums[1],
				rng: nums[2],
			})
		}

		sort.Slice(convertingMap.ranges, func(i, j int) bool {
			return convertingMap.ranges[i].src < convertingMap.ranges[j].src
		})

		maps[unitSrc] = convertingMap
	}

	return seeds, maps
}

func parseNumbers(str string) []int {
	var numbers []int

	for _, numStr := range strings.Split(str, " ") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
	}

	return numbers
}
