package main

import (
	"encoding/json"
	"log"

	"github.com/und3f/aoc/2023/fwk"
	"github.com/und3f/aoc/2023/fwk/twoD"
)

func main() {
	maze := fwk.ReadInputRunesLines()

	fwk.Solution(1, len(traceBeam(maze, []int{0, 0}, twoD.DirectionWest)))
	part2(maze)
}

func part2(maze [][]rune) {
	var best int

	for _, dir := range [][]int{
		twoD.DirectionSouth,
		twoD.DirectionNorth,
	} {
		i := 0
		if dir[0] < 0 {
			i = len(maze) - 1
		}

		for j := range maze[0] {
			energized := traceBeam(maze, []int{i, j}, dir)
			best = max(best, len(energized))
		}
	}

	for _, dir := range [][]int{
		twoD.DirectionEast,
		twoD.DirectionWest,
	} {
		j := 0
		if dir[1] < 0 {
			j = len(maze[0]) - 1
		}

		for i := range maze[0] {
			energized := traceBeam(maze, []int{i, j}, dir)
			best = max(best, len(energized))
		}
	}

	fwk.Solution(2, best)
}

type Beam struct {
	pos []int
	dir []int
}

func (b Beam) String() string {
	s, _ := json.Marshal([][]int{b.pos, b.dir})
	return string(s)
}

func traceBeam(maze [][]rune, pos []int, dir []int) [][]int {
	energizedMap := make([][]bool, len(maze))
	for i := range maze {
		energizedMap[i] = make([]bool, len(maze[i]))
	}

	visited := make(map[string]any)

	beams := []Beam{{pos: pos, dir: dir}}
	for len(beams) > 0 {
		newBeams := []Beam{}
		for _, beam := range beams {
			p := beam.pos
			if !(0 <= p[0] && p[0] < len(maze) && 0 <= p[1] && p[1] < len(maze[0])) {
				continue
			}
			if _, exists := visited[beam.String()]; exists {
				continue
			}
			visited[beam.String()] = struct{}{}
			energizedMap[p[0]][p[1]] = true

			rBeams := reflectBeams(maze[p[0]][p[1]], beam)
			for i := range rBeams {
				rBeams[i].pos = fwk.AddVec(rBeams[i].pos, rBeams[i].dir)
			}

			newBeams = append(newBeams, rBeams...)
		}

		beams = newBeams
	}

	var energized [][]int
	for i := range energizedMap {
		for j, v := range energizedMap[i] {
			if v {
				energized = append(energized, []int{i, j})
			}
		}
	}

	return energized
}

func reflectBeams(sym rune, beam Beam) []Beam {
	switch sym {
	case '.':
		return []Beam{beam}
	case '|':
		if beam.dir[1] != 0 {
			return []Beam{{
				pos: beam.pos,
				dir: twoD.DirectionSouth,
			}, {
				pos: beam.pos,
				dir: twoD.DirectionNorth,
			}}
		}
		return []Beam{beam}
	case '-':
		if beam.dir[0] != 0 {
			return []Beam{{
				pos: beam.pos,
				dir: twoD.DirectionEast,
			}, {
				pos: beam.pos,
				dir: twoD.DirectionWest,
			}}
		}
		return []Beam{beam}
	case '/':
		if beam.dir[0] != 0 {
			return []Beam{{
				pos: beam.pos,
				dir: twoD.RotateClockwise(beam.dir),
			}}
		} else {
			return []Beam{{
				pos: beam.pos,
				dir: twoD.RotateCounterclockwise(beam.dir),
			}}
		}
	case '\\':
		if beam.dir[0] != 0 {
			return []Beam{{
				pos: beam.pos,
				dir: twoD.RotateCounterclockwise(beam.dir),
			}}
		} else {
			return []Beam{{
				pos: beam.pos,
				dir: twoD.RotateClockwise(beam.dir),
			}}
		}

	default:
		log.Fatalf("Unexpected symbol: %c", sym)
	}

	return nil
}
