package main

import (
	queue "github.com/Jcowwell/go-algorithm-club/PriorityQueue"

	"github.com/und3f/aoc/2023/fwk"
	"github.com/und3f/aoc/2023/fwk/twoD"
)

var startPos = []int{0, 0}

func main() {
	in := fwk.ReadInputRunesLines()

	dijkstra := NewLavaDijkstra(in)
	startV := dijkstra.PosToV(startPos)
	endV := dijkstra.PosToV([]int{len(in) - 1, len(in[0]) - 1})

	fwk.Solution(1, dijkstra.FindPath(NewPath(startV), endV))
	fwk.Solution(2, dijkstra.FindPath(NewPathP2(startV), endV))

}

type LavaDijkstra struct {
	heatMap [][]int
	width   int
	height  int
}

func NewLavaDijkstra(heatMapRunes [][]rune) *LavaDijkstra {
	heatMap := make([][]int, len(heatMapRunes))
	for i, line := range heatMapRunes {
		heatMap[i] = make([]int, len(line))
		for j, char := range line {
			heatMap[i][j] = int(char) - '0'
		}
	}

	return &LavaDijkstra{
		heatMap: heatMap,
		width:   len(heatMap[0]),
		height:  len(heatMap),
	}
}

func (l *LavaDijkstra) PosToV(pos []int) int {
	return pos[0]*l.width + pos[1]
}

func (l *LavaDijkstra) VToPos(v int) []int {
	return []int{v / l.width, v % l.width}
}

func (l *LavaDijkstra) adj(v int) []int {
	pos := l.VToPos(v)

	var adj []int
	for _, dir := range [][]int{
		twoD.DirectionEast,
		twoD.DirectionSouth,
		twoD.DirectionWest,
		twoD.DirectionNorth,
	} {
		wPos := fwk.AddVec(pos, dir)
		if 0 <= wPos[0] && wPos[0] < l.height &&
			0 <= wPos[1] && wPos[1] < l.width {
			w := l.PosToV(wPos)
			adj = append(adj, w)
		}
	}

	return adj
}

type Path interface {
	Cur() int
	Loss() int

	Next(v, loss int) Path
	CacheKey() [3]int
	Compare(b Path) int
	CanEnd() bool
}

type P1Path struct {
	cur int

	loss int

	dir             int
	singleDirection int
}

func NewPath(v int) Path {
	return &P1Path{
		cur: v,
	}
}

func (p *P1Path) Cur() int {
	return p.cur
}

func (p *P1Path) Next(v int, loss int) Path {
	dir := v - p.cur

	if dir == -p.dir {
		return nil
	}

	singleDirection := 1
	if dir == p.dir {
		singleDirection = p.singleDirection + 1
		if singleDirection > 3 {
			return nil
		}
	}

	return &P1Path{
		cur:             v,
		loss:            p.loss + loss,
		dir:             dir,
		singleDirection: singleDirection,
	}
}

func (a *P1Path) Compare(_b Path) int {
	var b *P1Path
	switch i := _b.(type) {
	case *P1Path:
		b = i
	case *P2Path:
		b = &(i.P1Path)
	}

	if loss := a.loss - b.loss; loss != 0 {
		return loss
	}

	return a.Cur() - b.Cur()
}

func (p *P1Path) CacheKey() [3]int {
	return [3]int{
		p.cur,
		p.dir,
		p.singleDirection,
	}
}

func (p *P1Path) Loss() int {
	return p.loss
}

func (p *P1Path) CanEnd() bool {
	return true
}

type P2Path struct {
	P1Path
}

func NewPathP2(v int) Path {
	return &P2Path{
		P1Path{
			cur:             v,
			singleDirection: 4,
		},
	}
}

func (p *P2Path) CanEnd() bool {
	return p.P1Path.singleDirection >= 4
}

func (p *P2Path) Next(v int, loss int) Path {
	dir := v - p.cur

	if dir == -p.dir {
		return nil
	}

	singleDirection := 1
	if dir == p.dir {
		singleDirection = p.singleDirection + 1
		if singleDirection > 10 {
			return nil
		}
	} else {
		if p.singleDirection < 4 {
			return nil
		}
	}

	return &P2Path{
		P1Path{
			cur:             v,
			loss:            p.loss + loss,
			dir:             dir,
			singleDirection: singleDirection,
		},
	}
}

const MaxInt = int(^uint(0) >> 1)

func (l *LavaDijkstra) FindPath(startPath Path, end int) int {
	best := MaxInt

	posLoss := make(map[[3]int]int)

	pq := queue.PriorityQueueInit[Path](func(a, b Path) bool {
		return a.Compare(b) <= 0
	})
	pq.Enqueue(startPath)

	for !pq.IsEmpty() {
		p, _ := pq.Dequeue()

		if best <= p.Loss() {
			continue
		}

		if p.Cur() == end && p.CanEnd() {
			best = p.Loss()
		}

		for _, w := range l.adj(p.Cur()) {
			if nextP := p.Next(w, l.LossAt(w)); nextP != nil {
				cacheKey := nextP.CacheKey()
				if loss, exists := posLoss[cacheKey]; exists && loss <= p.Loss() {
					continue
				}
				posLoss[cacheKey] = p.Loss()

				pq.Enqueue(nextP)
			}
		}
	}

	return best
}

func (l *LavaDijkstra) LossAt(v int) int {
	pos := l.VToPos(v)
	return l.heatMap[pos[0]][pos[1]]
}
