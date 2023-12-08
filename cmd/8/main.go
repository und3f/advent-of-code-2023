package main

import (
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

const startNode = "AAA"
const endNode = "ZZZ"

func main() {
	instr, nodes := readInput()

	part1(instr, nodes)
	part2(instr, nodes)
}

func part1(instr []rune, nodes NodeMap) {
	steps, _ := reachEnd(startNode, 0, instr, nodes)
	fwk.Solution(1, steps)
}

func part2(instr []rune, nodes NodeMap) {
	var curNodes []string
	var nodeCycles []int
	for node := range nodes {
		if node[len(node)-1] == 'A' {
			curNodes = append(curNodes, node)

			endSteps, _ := reachEnd(node, 0, instr, nodes)

			nodeCycles = append(nodeCycles, endSteps)
		}
	}

	fwk.Solution(2, fwk.LCM(nodeCycles...))
}

func reachEnd(startNode string, startCount int, instr []rune, nodes NodeMap) (int, string) {
	count := startCount

	node := nodes[startNode].walk(instr[count%len(instr)])
	for ; node[len(node)-1] != 'Z'; node = nodes[node].walk(instr[count%len(instr)]) {
		count++
	}

	return count + 1, node
}

type Node struct {
	left, right string
}

func (node Node) walk(instruction rune) string {
	switch instruction {
	case 'L':
		return node.left
	case 'R':
		return node.right
	}

	panic("Unexpected instruction")
}

type NodeMap map[string]Node

func readInput() ([]rune, NodeMap) {
	in := fwk.ReadInputLines()

	instructions := in[0]

	nodes := make(NodeMap)
	for _, node := range in[2:] {
		str := strings.Split(node, " = ")
		name := str[0]

		strs := strings.Split(str[1], ", ")
		left := strs[0][1:]
		right := strs[1][:len(strs[1])-1]

		nodes[name] = Node{
			left:  left,
			right: right,
		}
	}

	return []rune(instructions), nodes
}
