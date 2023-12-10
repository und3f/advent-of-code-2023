package main

import (
	"slices"

	"github.com/und3f/aoc/2023/fwk"
	"github.com/und3f/aoc/2023/fwk/twoD"
)

const (
	StartPosition = 'S'
	VertPipe      = '|'
	HorPipe       = '-'
	BendNE        = 'L'
	BendNW        = 'J'
	BendSW        = '7'
	BendSE        = 'F'
	Ground        = '.'
)

func main() {
	maze := fwk.ReadInputRunesLines()
	startPos := findStartPos(maze)
	maze[startPos[0]][startPos[1]] = BendSE

	part1(maze, startPos)
	part2(maze, startPos)
}

func part1(maze [][]rune, startPos []int) {
	visits := NewVisitor(maze, startPos).Dijkstra()
	fwk.Solution(1, len(visits)-1)
}

func part2(maze [][]rune, startPos []int) {
	visits := NewVisitor(maze, startPos).Dijkstra()
	simplifiedMaze := make([][]rune, len(maze))
	for i := range maze {
		simplifiedMaze[i] = make([]rune, len(maze[i]))
		for j := range simplifiedMaze[i] {
			simplifiedMaze[i][j] = Ground
		}
	}

	for _, iteration := range visits {
		for _, p := range iteration {
			simplifiedMaze[p[0]][p[1]] = maze[p[0]][p[1]]
		}
	}

	insideTiles := NewIsideTracer(simplifiedMaze, startPos).FindAllInsideTiles()

	/*
			Print simplified maze with the inside marks

		for _, p := range insideTiles {
			simplifiedMaze[p[0]][p[1]] = 'I'
		}

		for _, line := range simplifiedMaze {
			fmt.Println(string(line))
		}
	*/

	fwk.Solution(2, len(insideTiles))
}

var SymToDirections = map[rune][][]int{
	HorPipe:  [][]int{twoD.DirectionWest, twoD.DirectionEast},
	VertPipe: [][]int{twoD.DirectionNorth, twoD.DirectionSouth},
	BendNE:   [][]int{twoD.DirectionNorth, twoD.DirectionEast},
	BendNW:   [][]int{twoD.DirectionNorth, twoD.DirectionWest},
	BendSE:   [][]int{twoD.DirectionSouth, twoD.DirectionEast},
	BendSW:   [][]int{twoD.DirectionSouth, twoD.DirectionWest},
}

type Visitor struct {
	maze    [][]rune
	pos     []int
	visited [][]bool
}

func NewVisitor(maze [][]rune, pos []int) *Visitor {
	visited := make([][]bool, len(maze))
	for i := range maze {
		visited[i] = make([]bool, len(maze[i]))
	}

	return &Visitor{
		maze:    maze,
		pos:     pos,
		visited: visited,
	}
}

func (v *Visitor) Dijkstra() [][][]int {
	var res [][][]int
	i := 0

	pq := [][]int{v.pos}
	for len(pq) > 0 {
		res = append(res, [][]int{})
		var newPq [][]int

		for _, pos := range pq {
			if v.IsVisited(pos) {
				continue
			}
			v.visited[pos[0]][pos[1]] = true
			res[i] = append(res[i], pos)

			sym := v.maze[pos[0]][pos[1]]
			dirs := SymToDirections[sym]

			for _, dir := range dirs {
				pos := fwk.AddVec(pos, dir)
				newPq = append(newPq, pos)
			}
		}
		pq = newPq
		i++
	}
	return res[:len(res)-1]
}

func (v *Visitor) IsVisited(pos []int) bool {
	return v.visited[pos[0]][pos[1]]
}

type InsideTracer struct {
	maze   [][]rune
	pos    []int
	inside [][]bool
}

func NewIsideTracer(maze [][]rune, pos []int) *InsideTracer {
	inside := make([][]bool, len(maze))
	for i := range maze {
		inside[i] = make([]bool, len(maze[i]))
	}
	return &InsideTracer{
		maze:   maze,
		pos:    pos,
		inside: inside,
	}
}

func (t *InsideTracer) FindAllInsideTiles() [][]int {
	t.trace()
	t.floodFillAllInsides()

	return t.getInsides()
}

func (t *InsideTracer) trace() {
	dir := twoD.DirectionWest
	pos := fwk.AddVec(t.pos, dir)
	dir = t.nextDir(pos, dir)

	for !slices.Equal(t.pos, pos) {
		pos = fwk.AddVec(pos, dir)
		dir = t.nextDir(pos, dir)

		for _, f := range []func([]int) []int{
			twoD.RotateClockwise45,
			twoD.RotateClockwise,
		} {
			iDir := f(dir)
			iPos := fwk.AddVec(pos, iDir)

			if !(0 <= iPos[0] && iPos[0] < len(t.maze) && 0 <= iPos[1] && iPos[1] < len(t.maze[iPos[0]])) {
				continue
			}
			// fmt.Println(pos, "->", iPos, dir, "->", iDir)

			sym := t.maze[iPos[0]][iPos[1]]
			if sym == Ground {
				t.inside[iPos[0]][iPos[1]] = true
				// fmt.Println("Adding --^")
			}
		}
	}
}

func (t *InsideTracer) nextDir(pos []int, dir []int) []int {
	sym := t.maze[pos[0]][pos[1]]
	dirs := SymToDirections[sym]

	negDir := twoD.Reverse(dir)
	for _, dir := range dirs {
		if !slices.Equal(dir, negDir) {
			return dir
		}
	}

	return nil
}

func (t *InsideTracer) floodFillAllInsides() {
	for i := range t.maze {
		for j := range t.maze[i] {
			if t.inside[i][j] {
				t.floodFillAround([]int{i, j})
			}
		}
	}
}

func (t *InsideTracer) floodFillAround(p []int) {
	for _, dir := range [][]int{
		twoD.DirectionWest,
		twoD.DirectionEast,
		twoD.DirectionNorth,
		twoD.DirectionSouth,
	} {
		t.floodFill(fwk.AddVec(p, dir))
	}
}

func (t *InsideTracer) floodFill(p []int) {
	if t.maze[p[0]][p[1]] != Ground || t.inside[p[0]][p[1]] {
		return
	}

	t.inside[p[0]][p[1]] = true
	t.floodFillAround(p)
}

func (t *InsideTracer) getInsides() [][]int {
	var poss [][]int

	for i := range t.maze {
		for j := range t.maze[i] {
			if t.inside[i][j] {
				poss = append(poss, []int{i, j})
			}
		}
	}

	return poss
}

func findStartPos(maze [][]rune) []int {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == StartPosition {
				return []int{i, j}
			}
		}
	}

	panic("Start pos not found")
}
