package main

import (
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	line := fwk.ReadInputLines()[0]
	steps := strings.Split(line, ",")

	part1(steps)
	part2(steps)
}

func part1(steps []string) {
	sum := uint(0)
	for _, step := range steps {
		sum += hash(step)
	}
	fwk.Solution(1, sum)
}

type HashMap struct {
	label string
	value uint64
}

func part2(steps []string) {
	boxes := make([][]HashMap, 256)

	for _, stepS := range steps {
		step := []rune(stepS)
		opIndex := strings.IndexAny(stepS, "=-")
		if opIndex == -1 {
			panic("Not found")
		}

		op := step[opIndex]
		label := string(step[:opIndex])
		labelHash := hash(label)
		value, _ := strconv.ParseUint(string(step[opIndex+1:]), 10, 64)

		switch op {
		case '-':
			for i, hm := range boxes[labelHash] {
				if strings.Compare(hm.label, label) == 0 {
					boxes[labelHash] = append(boxes[labelHash][:i], boxes[labelHash][i+1:]...)
					break
				}
			}
		case '=':
			found := false
			for i, hm := range boxes[labelHash] {
				if strings.Compare(hm.label, label) == 0 {
					found = true
					boxes[labelHash][i].value = value
					break
				}
			}
			if !found {
				boxes[labelHash] = append(boxes[labelHash], HashMap{
					label: label,
					value: value,
				})
			}
		}
	}

	sum := uint64(0)
	for boxI, box := range boxes {
		for slotI, lense := range box {
			value := uint64(boxI+1) * (uint64(slotI) + 1) * lense.value
			sum += value
		}
	}

	fwk.Solution(2, sum)
}

func hash(s string) uint {
	var val uint
	for _, c := range s {
		val += uint(c)
		val = val * 17
		val = val % 256
	}

	return val
}
