package main

import (
	"bytes"
	"regexp"
	"strconv"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	documents := readInput()

	part1(documents)
	part2(documents)
}

func part1(documents []Document) {
	mul := uint(1)
	for _, document := range documents {
		mul = mul * calPossibleSolutions(document)
	}

	fwk.Solution(1, mul)
}

func part2(documents []Document) {
	var timeBuf bytes.Buffer
	var distBuf bytes.Buffer
	for _, document := range documents {
		timeBuf.WriteString(strconv.FormatUint(uint64(document.time), 10))
		distBuf.WriteString(strconv.FormatUint(uint64(document.distance), 10))
	}

	time, _ := strconv.ParseUint(timeBuf.String(), 10, 64)
	dist, _ := strconv.ParseUint(distBuf.String(), 10, 64)

	fwk.Solution(2, calPossibleSolutions(Document{time: uint(time), distance: uint(dist)}))
}

func calPossibleSolutions(doc Document) uint {
	var holdTime uint
	for holdTime = 1; holdTime < doc.time && !isBeatingRecord(doc, holdTime); holdTime++ {
	}
	if !isBeatingRecord(doc, holdTime) {
		return 0
	}
	minSpeed := holdTime

	for ; holdTime < doc.time && isBeatingRecord(doc, holdTime); holdTime++ {
	}
	maxSpeed := holdTime - 1

	return maxSpeed - minSpeed + 1
}

func isBeatingRecord(doc Document, holdTime uint) bool {
	return (doc.time-holdTime)*holdTime > doc.distance
}

type Document struct {
	time, distance uint
}

var whitespaceRe = regexp.MustCompile(" +")

func readInput() []Document {
	lines := fwk.ReadInputLines()
	times := whitespaceRe.Split(lines[0], -1)[1:]
	distances := whitespaceRe.Split(lines[1], -1)[1:]
	documents := make([]Document, len(times))
	for i := range times {
		time, _ := strconv.ParseUint(times[i], 10, 64)
		distance, _ := strconv.ParseUint(distances[i], 10, 64)
		documents[i] = Document{time: uint(time), distance: uint(distance)}
	}

	return documents
}
