package main

import (
	"github.com/und3f/aoc/2023/fwk"
)

const Rounds = 1000000000

func main() {
	puzzle := NewPuzzle()
	puzzle.SlideNorth()
	fwk.Solution(1, puzzle.CalTotalLoad())

	cache := make(map[string]int)
	puzzle = NewPuzzle()
	for i := 0; i < Rounds; i++ {
		puzzle.SlideNorth()
		puzzle.SlideWest()
		puzzle.SlideSouth()
		puzzle.SlideEast()

		s := puzzle.String()
		if v, exists := cache[s]; exists {
			cycleSize := i - v
			roundsLeft := Rounds - i
			useCycleTimes := roundsLeft / cycleSize
			i += useCycleTimes * cycleSize
		} else {
			cache[puzzle.String()] = i
		}
	}
	fwk.Solution(2, puzzle.CalTotalLoad())
}

type Puzzle struct {
	puzzle [][]rune
}

func NewPuzzle() *Puzzle {
	return &Puzzle{fwk.ReadInputRunesLines()}
}

func (p *Puzzle) SlideNorth() {
	obstacle := make([]int, len(p.puzzle[0]))
	for j := 0; j < len(p.puzzle[0]); j++ {
		obstacle[j] = -1
	}

	for i := 0; i < len(p.puzzle); i++ {
		for j := 0; j < len(p.puzzle[i]); j++ {
			switch p.puzzle[i][j] {
			case '#':
				obstacle[j] = i
			case 'O':
				p.puzzle[i][j] = '.'
				p.puzzle[obstacle[j]+1][j] = 'O'
				obstacle[j]++
			}
		}
	}
}

func (p *Puzzle) SlideSouth() {
	obstacle := make([]int, len(p.puzzle[0]))
	for j := 0; j < len(p.puzzle[0]); j++ {
		obstacle[j] = len(p.puzzle)
	}

	for i := len(p.puzzle) - 1; i >= 0; i-- {
		for j := 0; j < len(p.puzzle[i]); j++ {
			switch p.puzzle[i][j] {
			case '#':
				obstacle[j] = i
			case 'O':
				p.puzzle[i][j] = '.'
				p.puzzle[obstacle[j]-1][j] = 'O'
				obstacle[j]--
			}
		}
	}
}

func (p *Puzzle) SlideWest() {
	obstacle := make([]int, len(p.puzzle[0]))
	for j := 0; j < len(p.puzzle); j++ {
		obstacle[j] = -1
	}

	for j := 0; j < len(p.puzzle[0]); j++ {
		for i := len(p.puzzle) - 1; i >= 0; i-- {
			switch p.puzzle[i][j] {
			case '#':
				obstacle[i] = j
			case 'O':
				p.puzzle[i][j] = '.'
				p.puzzle[i][obstacle[i]+1] = 'O'
				obstacle[i]++
			}
		}
	}
}

func (p *Puzzle) SlideEast() {
	obstacle := make([]int, len(p.puzzle[0]))
	for j := 0; j < len(p.puzzle); j++ {
		obstacle[j] = len(p.puzzle[0])
	}

	for j := len(p.puzzle[0]) - 1; j >= 0; j-- {
		for i := len(p.puzzle) - 1; i >= 0; i-- {
			switch p.puzzle[i][j] {
			case '#':
				obstacle[i] = j
			case 'O':
				p.puzzle[i][j] = '.'
				p.puzzle[i][obstacle[i]-1] = 'O'
				obstacle[i]--
			}
		}
	}
}

func (p *Puzzle) CalTotalLoad() uint {
	load := uint(0)
	for i := 0; i < len(p.puzzle); i++ {
		for j := 0; j < len(p.puzzle[i]); j++ {
			if p.puzzle[i][j] == 'O' {
				load += uint(len(p.puzzle) - i)
			}
		}
	}

	return load
}

func (p *Puzzle) String() string {
	return fwk.StringifyRunesLines(p.puzzle)
}
