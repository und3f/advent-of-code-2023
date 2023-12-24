package main

import (
	"fmt"
	"slices"

	queue "github.com/Jcowwell/go-algorithm-club/PriorityQueue"
	"github.com/und3f/aoc/2023/fwk"
	"github.com/und3f/aoc/2023/fwk/twoD"
)

/* Solving this day I've understand
 * that my time is more important
 * than computation time.
 * So I've just left it running,
 * and waited for result.
 */

func main() {
	in := fwk.ReadInputRunesLines()
	part1(in)
	part2(in)
}

func part1(m [][]rune) {
	start := [2]int{0, 1}
	end := [2]int{len(m) - 1, len(m) - 2}
	path := walk(m, start, end)
	fwk.Solution(1, len(path)-1)
}

func part2(m [][]rune) {
	start := [2]int{0, 1}
	end := [2]int{len(m) - 1, len(m) - 2}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			switch m[i][j] {
			case '>', 'v', '<', '^':
				m[i][j] = '.'
			}
		}
	}
	path := walk(m, start, end)
	fwk.Solution(1, len(path)-1)
}

func adj(m [][]rune, pos []int) [][2]int {
	switch m[pos[0]][pos[1]] {
	case '>', 'v', '<', '^':
		forcedPos := fwk.AddVec(pos, twoD.AsciiDirections[m[pos[0]][pos[1]]])
		return [][2]int{{forcedPos[0], forcedPos[1]}}
	}

	var adjacency [][2]int
	for _, d := range twoD.CardinalDirections {
		npos := fwk.AddVec(pos, d)
		if !(npos[0] >= 0 && npos[0] < len(m) &&
			npos[1] >= 0 && npos[1] < len(m[0])) {
			continue
		}

		switch m[npos[0]][npos[1]] {
		case '.', '>', 'v', '<', '^':
			adjacency = append(adjacency, [2]int{npos[0], npos[1]})
		}
	}

	return adjacency
}

type Path [][2]int

func (p Path) Compare(b Path) int {
	return len(p) - len(b)
}

func walk(m [][]rune, startPos, endPos [2]int) [][2]int {
	var longest *Path
	pq := queue.PriorityQueueInit[*Path](func(a, b *Path) bool {
		return len(*a) > len(*b)
	})
	pq.Enqueue(&Path{startPos})

	for !pq.IsEmpty() {
		path, _ := pq.Dequeue()

		last := (*path)[len(*path)-1]
		if last == endPos {
			if longest == nil || len(*longest) < len(*path) {
				longest = path
				fmt.Println("Found solution", len(*longest)-1)
				continue
			}
		}

		adjacency := adj(m, last[:])
		var beforeLast [2]int
		if len(*path) >= 2 {
			beforeLast = (*path)[len(*path)-2]
		}
		for i := range adjacency {
			if adjacency[i] == beforeLast {
				adjacency = append(adjacency[:i], adjacency[i+1:]...)
				break
			}
		}

		for _, w := range adjacency {
			warr := [2]int{w[0], w[1]}
			if !slices.Contains([][2]int(*path), warr) {
				nPath := *path
				if len(adjacency) > 1 {
					nPath = slices.Clone(*path)
				}
				nPath = append(nPath, w)
				pq.Enqueue(&nPath)
			}
		}
	}

	return *longest
}
