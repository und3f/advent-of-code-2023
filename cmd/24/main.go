package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	hailstones := parseInput()

	part1(hailstones)
	part2(hailstones)
}

const (
	COORD_LEAST = 200000000000000
	COORD_MOST  = 400000000000000
)

func part1(hailstones [][2][3]float64) {
	fwk.Solution(1, findIntersection2D(
		hailstones,
		[]float64{COORD_LEAST, COORD_LEAST},
		[]float64{COORD_MOST, COORD_MOST},
	))
}

var coords = []string{"x", "y", "z"}

func part2(hailstones [][2][3]float64) {
	file, err := os.CreateTemp("", "aoc-day24.*.smt2")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file, "(declare-const solution Real)\n")
	for _, coord := range coords {
		fmt.Fprintf(file, "(declare-const %s Real)\n", coord)
		fmt.Fprintf(file, "(declare-const v%s Real)\n", coord)
		fmt.Fprintf(file, "(assert (>= v%s -300))\n", coord)
		fmt.Fprintf(file, "(assert (<= v%s 300))\n", coord)
	}
	fmt.Fprintf(file, "(assert (= solution (+ x y z)))\n")

	for i := range hailstones {
		h := hailstones[i]
		timeConst := fmt.Sprintf("t%d", i)
		fmt.Fprintf(file, "(declare-const %s Real)\n", timeConst)
		fmt.Fprintf(file, "(assert (>= %s 0.0))\n", timeConst)
		for dim, coord := range coords {
			fmt.Fprintf(file,
				"(assert (= (+ %s (* v%s %s)) (+ %f (* %f %s))))\n",
				coord, coord,
				timeConst,
				h[0][dim],
				h[1][dim],
				timeConst,
			)
		}
	}

	fmt.Fprintln(file, "(check-sat)")
	fmt.Fprintln(file, "(get-model)")
	fmt.Printf("Run: z3 -smt2 %q\n", file.Name())
}

func findIntersection2D(hailstones [][2][3]float64, tl, br []float64) int {
	var sum int

	for i, a := range hailstones {
		for j := i + 1; j < len(hailstones); j++ {
			b := hailstones[j]

			posA := a[0]
			velA := a[1]
			posB := b[0]
			velB := b[1]

			if isParallel2D(velA, velB) {
				continue
			}

			t2 := (velA[0]*(posB[1]-posA[1]) - velA[1]*(posB[0]-posA[0])) / (velA[1]*velB[0] - velA[0]*velB[1])
			t1 := (posB[0] - posA[0] + velB[0]*t2) / velA[0]

			intersection := [4]float64{
				posB[0] + t2*velB[0],
				posB[1] + t2*velB[1],
				posA[0] + t1*velA[0],
				posA[1] + t1*velA[1],
			}

			if t2 >= 0. && t1 >= 0. {
				if tl[0] <= intersection[0] && intersection[0] <= br[0] &&
					tl[1] <= intersection[1] && intersection[1] <= br[1] {
					sum++
				}
			}
		}
	}

	return sum
}

func isParallel2D(velA, velB [3]float64) bool {
	return int64(velA[0])*int64(velB[1]) == int64(velB[0])*int64(velA[1])
}

var positionVelicitySplitRe = regexp.MustCompile("\\s+@\\s+")

func parseInput() [][2][3]float64 {
	lines := fwk.ReadInputLines()
	r := make([][2][3]float64, len(lines))

	for i := range lines {
		s := positionVelicitySplitRe.Split(lines[i], -1)
		r[i][0] = parseThreeCoords(s[0])
		r[i][1] = parseThreeCoords(s[1])
	}

	return r
}

var threeCoordsSplitRe = regexp.MustCompile(", +")

func parseThreeCoords(s string) [3]float64 {
	numStrs := threeCoordsSplitRe.Split(s, -1)
	var r [3]float64
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi(numStrs[i])
		if err != nil {
			panic(err)
		}
		r[i] = float64(num)
	}
	return r
}
