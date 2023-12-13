package main

import (
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	patterns := readInput()
	fwk.Solution(1, calMirrors(patterns, 0))
	fwk.Solution(2, calMirrors(patterns, 1))
}

func calMirrors(patterns [][][]rune, smudge int) uint {
	sum := uint(0)

	for _, pattern := range patterns {
		if leftColumns, found := findVertReflection(pattern, smudge); found {
			sum += uint(leftColumns)
		} else if aboveColumns, found := findHorReflection(pattern, smudge); found {
			sum += uint(100 * aboveColumns)
		} else {
			fwk.PrintRunesLines(pattern)
			panic("Match not found")
		}
	}

	return sum
}

func findVertReflection(pattern [][]rune, smudge int) (int, bool) {
	for i := len(pattern[0]) - 1; i >= 1; i-- {
		if matchVert(pattern, i, smudge) {
			return i, true
		}
	}

	return 0, false
}

func matchVert(pattern [][]rune, vert int, smudge int) bool {
	maxJ := min(vert, len(pattern[0])-vert)
	for j := 0; j < maxJ; j++ {
		for i := 0; i < len(pattern); i++ {
			if pattern[i][vert-j-1] != pattern[i][vert+j] {
				if smudge--; smudge < 0 {
					return false
				}
			}
		}
	}
	return smudge == 0
}

func findHorReflection(pattern [][]rune, smudge int) (int, bool) {
	for i := len(pattern) - 1; i >= 1; i-- {
		if matchHor(pattern, i, smudge) {
			return i, true
		}
	}

	return 0, false
}

func matchHor(pattern [][]rune, hor int, smudge int) bool {
	maxI := min(hor, len(pattern)-hor)
	for i := 0; i < maxI; i++ {
		for j := 0; j < len(pattern[0]); j++ {
			if pattern[hor-i-1][j] != pattern[hor+i][j] {
				if smudge--; smudge < 0 {
					return false
				}
			}
		}
	}
	return smudge == 0
}

func readInput() [][][]rune {
	sections := strings.Split(strings.TrimSpace(fwk.ReadInput("")), "\n\n")

	patterns := make([][][]rune, len(sections))
	for i, section := range sections {
		lines := strings.Split(section, "\n")
		pattern := make([][]rune, len(lines))
		for j, line := range lines {
			pattern[j] = []rune(line)
		}
		patterns[i] = pattern
	}

	return patterns
}
