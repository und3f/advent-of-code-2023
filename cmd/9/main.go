package main

import (
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	values := readInputValues()

	sumNextValuesL := 0
	sumNextValuesR := 0

	for _, valueHist := range values {
		lVal, rVal := calNextValue(valueHist)
		sumNextValuesL += lVal
		sumNextValuesR += rVal
	}

	fwk.Solution(1, sumNextValuesR)
	fwk.Solution(2, sumNextValuesL)
}

func calNextValue(in []int) (int, int) {
	diffs := [][]int{}

	for diffL := in; !isAllEqual(diffL, 0); diffL = diff(diffL) {
		diffs = append(diffs, diffL)
	}

	rVal := 0
	lVal := 0
	for i := len(diffs) - 1; i >= 0; i-- {
		diffL := diffs[i]

		rVal += diffL[len(diffL)-1]
		lVal = diffL[0] - lVal
	}

	return lVal, rVal
}

func diff(in []int) []int {
	out := make([]int, len(in)-1)
	for i := 0; i < len(in)-1; i++ {
		out[i] = in[i+1] - in[i]
	}

	return out
}

func isAllEqual(in []int, val int) bool {
	for _, v := range in {
		if v != val {
			return false
		}
	}
	return true
}

func readInputValues() [][]int {
	var values [][]int
	for _, line := range fwk.ReadInputLines() {
		historyStrs := strings.Split(line, " ")
		valueHistory := make([]int, len(historyStrs))

		for i, str := range historyStrs {
			var err error
			valueHistory[i], err = strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
		}
		values = append(values, valueHistory)
	}

	return values
}
