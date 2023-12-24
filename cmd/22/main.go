package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	part1(parseInput())
	part2(parseInput())
}

func part1(bricks [][][]int) {
	var sum int

	isReliance := make([]bool, len(bricks))
	for _, deps := range buildBricksDep(bricks) {
		if len(deps) == 1 {
			isReliance[deps[0]] = true
		}
	}

	for _, v := range isReliance {
		if !v {
			sum++
		}
	}

	fwk.Solution(1, sum)
}

func part2(bricks [][][]int) {
	var sum int

	depends := buildBricksDep(bricks)
	adj := make([][]bool, len(bricks))
	for i, deps := range depends {
		adj[i] = make([]bool, len(bricks))
		for _, j := range deps {
			adj[i][j] = true
		}
	}

	for i, _ := range bricks {
		disintigrated := make(map[int]any)
		disintigrated[i] = struct{}{}

		for j := i + 1; j < len(bricks); j++ {
			shouldFall := false

			for k := 0; k < j; k++ {
				if adj[j][k] {
					shouldFall = true
					if _, isDisintigrated := disintigrated[k]; !isDisintigrated {
						shouldFall = false
						break
					}
				}
			}

			if shouldFall {
				disintigrated[j] = struct{}{}
			}
		}

		sum += len(disintigrated) - 1
	}

	fwk.Solution(2, sum)
}

func buildBricksDep(bricks [][][]int) [][]int {
	fallBricks(bricks)
	slices.SortFunc(bricks, compareBricksHighZ)

	reliances := make([][]int, len(bricks))
	for i, brick := range bricks {
		for j := i - 1; j >= 0; j-- {
			relianceBrick := bricks[j]
			if isOverlapXY(brick, relianceBrick) && relianceBrick[1][2]+1 == brick[0][2] {
				reliances[i] = append(reliances[i], j)
			}
		}
	}

	return reliances
}

func fallBricks(bricks [][][]int) {
	slices.SortFunc(bricks, compareBricksLowZ)
	ground := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	for i, brick := range bricks {
		reliance := ground
		for j := i - 1; j >= 0; j-- {
			prevBrick := bricks[j]
			if prevBrick[1][2] > reliance[1][2] && isOverlapXY(prevBrick, brick) {
				reliance = prevBrick
			}
		}

		yOffset := brick[0][2] - reliance[1][2] - 1
		brick[0][2] -= yOffset
		brick[1][2] -= yOffset
	}
}

func isOverlapXY(a, b [][]int) bool {
	return (a[0][0] <= b[1][0] && a[0][1] <= b[1][1]) &&
		(a[1][0] >= b[0][0] && a[1][1] >= b[0][1])
}

func parseInput() [][][]int {
	lines := fwk.ReadInputLines()
	snapshot := make([][][]int, len(lines))
	for i := range lines {
		s := strings.Split(lines[i], "~")
		snapshot[i] = make([][]int, len(s))
		for j := range s {
			snapshot[i][j] = parseSnapshot(s[j])
		}
	}
	return snapshot
}

func parseSnapshot(s string) []int {
	numStrs := strings.Split(s, ",")
	nums := make([]int, len(numStrs))

	for i := range numStrs {
		nums[i], _ = strconv.Atoi(numStrs[i])
	}

	return nums
}

func compareBricksLowZ(a, b [][]int) int {
	if z := a[0][2] - b[0][2]; z != 0 {
		return z
	}
	if x := a[0][0] - b[0][0]; x != 0 {
		return x
	}
	return a[0][1] - b[0][1]
}

func compareBricksHighZ(a, b [][]int) int {
	if z := a[1][2] - b[1][2]; z != 0 {
		return z
	}
	if x := a[0][0] - b[0][0]; x != 0 {
		return x
	}
	return a[0][1] - b[0][1]
}
