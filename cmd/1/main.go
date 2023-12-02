package main

import (
	"strings"
	"unicode"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	lines := fwk.ReadInputLines()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		firstI := strings.IndexFunc(line, unicode.IsNumber)
		lastI := strings.LastIndexFunc(line, unicode.IsNumber)

		num := int((line[firstI]-'0')*10 + line[lastI] - '0')
		sum += num
	}

	fwk.Solution(1, sum)
}

var digits []string = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		firstI := len(line)
		firstD := -1

		lastI := -1
		lastD := -1
		for i := range digits {
			p := strings.Index(line, digits[i])
			if p >= 0 {
				if p < firstI {
					firstI = p
					firstD = i + 1
				}
			}
		}

		for i := range digits {
			p := strings.LastIndex(line, digits[i])
			if p >= 0 {
				if p > lastI {
					lastI = p
					lastD = i + 1
				}
			}
		}

		firstDigitPos := strings.IndexFunc(line, unicode.IsNumber)
		if firstDigitPos >= 0 && firstDigitPos < firstI {
			firstI = firstDigitPos
			firstD = int(line[firstDigitPos] - '0')
		}

		lastDigitPos := strings.LastIndexFunc(line, unicode.IsNumber)
		if lastDigitPos >= 0 && lastDigitPos > lastI {
			lastI = lastDigitPos
			lastD = int(line[lastDigitPos] - '0')
		}

		sum += firstD*10 + lastD
	}

	fwk.Solution(2, sum)
}
