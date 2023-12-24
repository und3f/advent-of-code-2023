package main

import (
	"fmt"
	"slices"

	"github.com/und3f/aoc/2023/cmd/9/calculator"
	"github.com/und3f/aoc/2023/fwk"
	"github.com/und3f/aoc/2023/fwk/twoD"
)

func main() {
	part1(fwk.ReadInputRunesLines())
	part2(fwk.ReadInputRunesLines())
}

const Part1Steps = 64

func part1(gardenMap [][]rune) {
	startPos := findStartPos(gardenMap)
	gardenMap[startPos[0]][startPos[1]] = '.'
	positions := steps(gardenMap, [][]int{startPos}, Part1Steps, gardenMapClip(gardenMap))
	fwk.Solution(1, len(positions))
}

func gardenMapClip(gardenMap [][]rune) func([]int) bool {
	return func(pos []int) bool {
		return !(0 <= pos[0] && pos[0] < len(gardenMap) &&
			0 <= pos[1] && pos[1] < len(gardenMap[0]))
	}
}

const Part2Steps = 26501365
const magicNumber = 65

/* Part2 solution is based on spoiler
 * https://www.reddit.com/r/adventofcode/comments/18orn0s/2023_day_21_part_2_links_between_days/
 * My contribution is just the implementation of infinite positions computations
 */
func part2(gardenMap [][]rune) {
	startPos := findStartPos(gardenMap)
	gardenMap[startPos[0]][startPos[1]] = '.'

	mapSize := len(gardenMap)
	mapSizeSteps := (Part2Steps - magicNumber) / mapSize

	positions := steps(gardenMap, [][]int{startPos},
		magicNumber,
		func([]int) bool { return false },
	)

	poly := []uint{uint(len(positions))}
	for i := 0; i < 2; i++ {
		positions = steps(gardenMap, positions,
			mapSize, func([]int) bool { return false },
		)

		poly = append(poly, uint(len(positions)))
		mapSizeSteps--
	}

	for ; mapSizeSteps > 0; mapSizeSteps-- {
		_, rVal := calculator.CalNextValue(poly)
		poly = append(poly[1:], rVal)
	}

	fwk.Solution(2, poly[len(poly)-1])
}

func steps(gardenMap [][]rune, startPos [][]int, steps int, clip func([]int) bool) [][]int {
	positions := slices.Clone(startPos)
	dimenstions := []int{len(gardenMap), len(gardenMap[0])}

	for i := 0; i < steps; i++ {
		var newPos [][]int

		for _, pos := range positions {
			for _, dir := range twoD.CardinalDirections {
				nPos := fwk.AddVec(pos, dir)

				if clip(nPos) {
					continue
				}

				modPos := ModVect(nPos, dimenstions)
				if gardenMap[modPos[0]][modPos[1]] == '.' {
					newPos = append(newPos, nPos)
				}
			}
		}

		slices.SortFunc(newPos, func(a, b []int) int {
			if y := a[0] - b[0]; y != 0 {
				return y
			}
			return a[1] - b[1]
		})
		newPos = slices.CompactFunc(newPos, func(a, b []int) bool {
			return a[0] == b[0] && a[1] == b[1]
		})

		positions = newPos
		// plotMapAmount(gardenMap, positions)
	}

	return positions
}

func mod(d, m int) int {
	res := d % m
	if res < 0 {
		res += m
	}
	return res
}

func ModVect(d, m []int) []int {
	r := make([]int, len(d))
	for i := range d {
		r[i] = mod(d[i], m[i])
	}
	return r
}

func findStartPos(gardenMap [][]rune) []int {
	for i := range gardenMap {
		for j := range gardenMap[i] {
			if gardenMap[i][j] == 'S' {
				return []int{i, j}
			}
		}
	}
	return nil
}

func plotMapAmount(gardenMap [][]rune, positions [][]int) {
	grid := fwk.NewCustomInfiniteGrid[int](0, func(v int) string {
		if v < 0 {
			v = 0
		}
		return fmt.Sprintf("%4d", v)
	})

	for i := -4; i <= 4; i++ {
		for j := -4; j <= 4; j++ {
			grid.SetAt([]int{i, j}, -1)
		}
	}

	for _, p := range positions {
		gridP := []int{p[0] / len(gardenMap), p[1] / len(gardenMap[0])}
		grid.SetAt(gridP, 1+grid.GetAt(gridP))
	}
	fmt.Println(grid)
}

func plotMap(gardenMap [][]rune, positions [][]int) {
	grid := fwk.NewInfiniteGrid[rune]()

	var tl, br []int = []int{0, 0}, []int{0, 0}
	for _, p := range positions {
		tl[0] = min(tl[0], p[0])
		tl[1] = min(tl[1], p[1])

		br[0] = max(tl[0], p[0])
		br[1] = max(tl[1], p[1])
	}

	/*
		for i := tl[0] / len(gardenMap); i < int(math.Ceil(float64(br[0])/float64(len(gardenMap)))); i++ {
			for j := tl[1] / len(gardenMap[0]); j < int(math.Ceil(float64(br[1])/float64(len(gardenMap[0])))); j++ {
				plotMapAt(grid, gardenMap, []int{j, i})
			}
		}
		plotMapAt(grid, gardenMap, []int{0, 0})
	*/

	for _, p := range positions {
		grid.SetAt(p, 'O')
	}
	fmt.Println(grid.String())
}

func plotMapAt(grid fwk.Grid[rune], gardenMap [][]rune, pos []int) {
	offset := []int{
		len(gardenMap) * pos[0],
		len(gardenMap[0]) * pos[1],
	}
	for i := range gardenMap {
		for j, v := range gardenMap[i] {
			grid.SetAt(
				fwk.AddVec([]int{i, j}, offset),
				v,
			)
		}
	}
}
