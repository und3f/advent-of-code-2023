package main

import (
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/cmd/9/calculator"
	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	values := readInputValues()

	sumNextValuesL := 0
	sumNextValuesR := 0

	for _, valueHist := range values {
		lVal, rVal := calculator.CalNextValue(valueHist)
		sumNextValuesL += lVal
		sumNextValuesR += rVal
	}

	fwk.Solution(1, sumNextValuesR)
	fwk.Solution(2, sumNextValuesL)
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
