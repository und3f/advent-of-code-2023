package fwk

import "fmt"

func PrintRunesLines(pattern [][]rune) {
	for _, line := range pattern {
		fmt.Println(string(line))
	}
	fmt.Println()
}
