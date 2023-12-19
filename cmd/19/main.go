package main

import (
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	rules, parts := parseInput()
	part1(rules, parts)
	part2(rules)
}

func part1(workflows map[string][]Rule, parts []map[string]int) {
	var sum int
	for _, part := range parts {
		if evaluate(workflows, part) {
			sum += calRating(part)
		}
	}
	fwk.Solution(1, sum)
}

const startWorkflow = "in"

func evaluate(workflows map[string][]Rule, part map[string]int) bool {
	workflow := startWorkflow

	for !slices.Contains([]string{"A", "R"}, workflow) {

		rules := workflows[workflow]
		matched := false
		for _, rule := range rules {
			switch rule.operation {
			case '>':
				if part[rule.variable] > rule.value {
					matched = true
				}
			case '<':
				if part[rule.variable] < rule.value {
					matched = true
				}
			default:
				matched = true
				workflow = rule.resolution
				break
			}

			if matched {
				workflow = rule.resolution
				break
			}
		}
	}

	return strings.Compare("A", workflow) == 0
}

func calRating(part map[string]int) int {
	var sum int
	for _, v := range part {
		sum += v
	}
	return sum
}

func part2(workflows map[string][]Rule) {
	initRange := make(map[string][2]int)
	for _, variable := range []string{"x", "m", "a", "s"} {
		initRange[variable] = [2]int{1, 4000}
	}

	fwk.Solution(2, evaluateRange(workflows, startWorkflow, initRange))
}

func evaluateRange(
	workflows map[string][]Rule,
	workflow string,
	initialRange map[string][2]int,
) uint {
	switch workflow {
	case "A":
		return calRangeRating(initialRange)
	case "R":
		return 0
	}

	var sum uint
	curRange := maps.Clone(initialRange)
	for _, part := range workflows[workflow] {
		switch part.operation {
		case '<':
			varRng := curRange[part.variable]
			if varRng[0] < part.value {
				matchRange := maps.Clone(curRange)
				matchRange[part.variable] = [2]int{
					varRng[0], min(varRng[1], part.value-1),
				}
				sum += evaluateRange(workflows, part.resolution, matchRange)
			}

			varRng[0] = part.value
			if varRng[0] > varRng[1] {
				return sum
			}
			curRange[part.variable] = varRng
		case '>':
			varRng := curRange[part.variable]
			if part.value < varRng[1] {
				matchRange := maps.Clone(curRange)
				matchRange[part.variable] = [2]int{
					max(varRng[0], part.value+1), varRng[1],
				}
				sum += evaluateRange(workflows, part.resolution, matchRange)
			}

			varRng[1] = part.value
			if varRng[0] > varRng[1] {
				return sum
			}
			curRange[part.variable] = varRng
		case ':':
			sum += evaluateRange(workflows, part.resolution, curRange)
		default:
			panic("Not expected")
		}
	}

	return sum
}

func calRangeRating(part map[string][2]int) uint {
	var sum uint = 1
	for _, v := range part {
		sum *= uint(v[1]-v[0]) + 1
	}
	return sum
}

type Rule struct {
	variable  string
	operation rune
	value     int

	resolution string
}

var workflowRe = regexp.MustCompile("^(\\w+){(.+)}$")
var ruleRe = regexp.MustCompile("^(\\w+)(.)(-?\\d+):(\\w+)")

func parseInput() (map[string][]Rule, []map[string]int) {
	workflows := make(map[string][]Rule)
	var vars []map[string]int

	in := strings.TrimSpace(fwk.ReadInput(""))

	sections := strings.Split(in, "\n\n")
	for _, l := range strings.Split(sections[0], "\n") {
		match := workflowRe.FindStringSubmatch(l)
		name := match[1]

		rulesStrs := strings.Split(match[2], ",")
		rules := make([]Rule, len(rulesStrs))
		for i, str := range rulesStrs[:len(rulesStrs)-1] {
			match := ruleRe.FindStringSubmatch(str)
			value, _ := strconv.Atoi(match[3])
			rules[i] = Rule{
				variable:   match[1],
				operation:  []rune(match[2])[0],
				value:      value,
				resolution: match[4],
			}
			rules[len(rules)-1] = Rule{
				operation:  ':',
				resolution: rulesStrs[len(rulesStrs)-1],
			}
		}
		workflows[name] = rules
	}

	for _, l := range strings.Split(sections[1], "\n") {
		data := l[1 : len(l)-1]
		varMap := make(map[string]int)
		for _, s := range strings.Split(data, ",") {
			v := strings.Split(s, "=")
			value, _ := strconv.Atoi(v[1])
			varMap[v[0]] = value
		}
		vars = append(vars, varMap)
	}

	return workflows, vars
}
