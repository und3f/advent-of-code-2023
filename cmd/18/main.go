package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
	"github.com/und3f/aoc/2023/fwk/twoD"
)

func main() {
	in := parseInput()
	fwk.Solution(1, bruteforce(in))

	plan := make([]Dig, len(in))
	for i := range in {
		color := in[i].color
		meters := color / 0x10
		dir := DirectionsFromNumber[color%0x10]
		plan[i] = Dig{
			meters: meters,
			dir:    dir,
		}
	}
	fwk.Solution(2, shoelace(plan))
}

func shoelace(plan []Dig) int {
	var sum int
	start := []int{0, 0}
	pos := slices.Clone(start)

	for _, dig := range plan {
		npos := fwk.AddVec(pos, fwk.MultVecByConstant(dig.dir, int(dig.meters)))

		sum += pos[1]*npos[0] - npos[1]*pos[0]
		sum += int(dig.meters)

		pos = npos
	}

	sum += pos[1]*start[0] - start[1]*pos[0]
	sum += fwk.Abs(start[0]-pos[0]) + fwk.Abs(start[1]-pos[1])

	return sum/2 + 1
}

func bruteforce(plan []Dig) uint {
	grid := fwk.NewInfiniteGrid[rune]()

	var prevDirection []int
	var turnPoints [][]int

	pos := []int{0, 0}

	for _, dig := range plan[0:1] {
		for i := uint64(0); i < dig.meters; i++ {
			grid.SetAt(pos, '#')
			pos = fwk.AddVec(pos, dig.dir)
		}
		prevDirection = plan[0].dir
	}

	for _, dig := range plan[1:] {
		turnPoints = append(turnPoints, calTurnPoint(pos, prevDirection, dig.dir)...)
		for i := uint64(0); i < dig.meters; i++ {
			grid.SetAt(pos, '#')
			pos = fwk.AddVec(pos, dig.dir)
		}
		prevDirection = dig.dir
	}

	/*
		for _, p := range turnPoints {
			floodFill(grid, p, '#')
		}
	*/
	for _, p := range turnPoints {
		grid.SetAt(p, 'T')
	}

	return grid.CountAll('#')
}

func fill(grid fwk.Grid[rune], tl, br []int, value rune) {
	for i := tl[0]; i <= br[0]; i++ {
		for j := tl[1]; j <= br[1]; j++ {
			grid.SetAt([]int{i, j}, value)
		}
	}
}

func floodFill(grid fwk.Grid[rune], pos []int, value rune) {
	if cur := grid.GetAt(pos); cur != '.' {
		return
	}
	grid.SetAt(pos, value)
	for _, dir := range Directions {
		floodFill(grid, fwk.AddVec(pos, dir), value)
	}
}

func calTurnPoint(pos, prevDir, dir []int) [][]int {
	if slices.Equal(twoD.RotateClockwise(prevDir), dir) {
		return [][]int{fwk.AddVec(fwk.AddVec(pos, twoD.Reverse(prevDir)), dir)}
	}

	return [][]int{
		fwk.AddVec(pos, prevDir),
		fwk.AddVec(pos, twoD.Reverse(dir)),
	}
}

type Dig struct {
	dir    []int
	meters uint64

	color uint64
}

var Directions = map[rune][]int{
	'U': twoD.DirectionNorth,
	'D': twoD.DirectionSouth,
	'L': twoD.DirectionWest,
	'R': twoD.DirectionEast,
}

var DirectionsFromNumber = map[uint64][]int{
	3: twoD.DirectionNorth,
	1: twoD.DirectionSouth,
	2: twoD.DirectionWest,
	0: twoD.DirectionEast,
}

func parseInput() []Dig {
	lines := fwk.ReadInputLines()
	plan := make([]Dig, len(lines))
	for i := range lines {
		strs := strings.Split(lines[i], " ")
		meters, _ := strconv.ParseUint(strs[1], 10, 64)
		color, _ := strconv.ParseUint(strs[2][2:8], 16, 64)
		plan[i] = Dig{
			dir:    Directions[[]rune(strs[0])[0]],
			meters: meters,
			color:  color,
		}
	}

	return plan
}
