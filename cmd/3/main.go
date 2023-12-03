package main

import (
	"strconv"
	"unicode"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	schematic := fwk.ReadInputRunesLines()

	part1(schematic)
	part2(schematic)
}

func part1(schematic [][]rune) {
	engineSum := 0

	forEachSymbolPositions(schematic, func(i, j int) {
		for _, num := range findAdjancedNumber(schematic, i, j) {
			engineSum += num
		}
	})

	fwk.Solution(1, engineSum)
}

func part2(schematic [][]rune) {
	gearRatioSum := 0

	forEachSymbolPositions(schematic, func(i, j int) {
		if schematic[i][j] == '*' {
			adj := findAdjancedNumber(schematic, i, j)
			if len(adj) == 2 {
				gearRatioSum += adj[0] * adj[1]
			}
		}
	})

	fwk.Solution(2, gearRatioSum)
}

func findAdjancedNumber(schematic [][]rune, i, j int) []int {
	var nums []int
	for iOffset := -1; iOffset <= 1; iOffset++ {
		for jOffset := -1; jOffset <= 1; jOffset++ {
			if unicode.IsNumber(schematic[i+iOffset][j+jOffset]) {
				row := schematic[i+iOffset]

				left, right := findNumberBoundaries(row, j+jOffset)

				num, _ := strconv.Atoi(string(row[left : right+1]))
				// fmt.Printf("Number: %d\n", num)
				nums = append(nums, num)

				jOffset = right - j
			}
		}
	}

	return nums
}

func forEachSymbolPositions(schematic [][]rune, cb func(y, x int)) {
	for i := range schematic {
		for j, sym := range schematic[i] {
			if isSymbol(sym) {
				cb(i, j)
			}
		}
	}

}

func findNumberBoundaries(row []rune, j int) (left, right int) {
	left = j
	right = j

	for ; left > 0 && unicode.IsNumber(row[left-1]); left-- {
	}

	for ; right < len(row)-1 && unicode.IsNumber(row[right+1]); right++ {
	}

	return
}

func isSymbol(sym rune) bool {
	return sym != '.' && !unicode.IsNumber(sym)
}
