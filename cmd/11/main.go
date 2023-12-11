package main

import (
	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	image := fwk.ReadInputRunesLines()
	fwk.Solution(1, calDistancesAfterExpansion(image, 2))
	fwk.Solution(2, calDistancesAfterExpansion(image, 1000000))
}

const SymGalaxy = '#'
const SymEmpty = '.'

func calDistancesAfterExpansion(image [][]rune, expandMultiplier uint) uint {
	colsExpand, rowsExpand := calExpands(image, expandMultiplier)

	var galaxies [][]uint
	for i := range image {
		for j, v := range image[i] {
			if v == SymGalaxy {
				pos := []uint{uint(i) + rowsExpand[i], uint(j) + colsExpand[j]}
				galaxies = append(galaxies, pos)
			}
		}
	}

	sum := uint(0)
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := CalManhattan(galaxies[j], galaxies[i])
			sum += dist
			// fmt.Printf("Compare star %d and %d => %d\n", i+1, j+1, dist)
		}
	}

	return sum
}

func CalManhattan(a, b []uint) uint {
	manhattan := uint(0)
	for i := range a {
		minuend := max(a[i], b[i])
		subtrahend := min(a[i], b[i])
		manhattan += minuend - subtrahend
	}

	return manhattan
}

func calExpands(image [][]rune, expandMultiplier uint) ([]uint, []uint) {
	colsExpand := make([]uint, len(image))
	rowsExpand := make([]uint, len(image[0]))

	rowExpand := uint(0)
	for i := range image {
		isRowEmpty := true
		for _, v := range image[i] {
			if v != SymEmpty {
				isRowEmpty = false
				break
			}
		}
		if isRowEmpty {
			rowExpand += 1*expandMultiplier - 1
		}
		rowsExpand[i] = rowExpand
	}

	colExpand := uint(0)
	for j := range image[0] {
		isColEmpty := true
		for _, row := range image {
			if row[j] != SymEmpty {
				isColEmpty = false
				break
			}
		}
		if isColEmpty {
			colExpand += 1*expandMultiplier - 1
		}
		colsExpand[j] = colExpand
	}

	return colsExpand, rowsExpand
}
