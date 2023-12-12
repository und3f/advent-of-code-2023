package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2023/cmd/12/state"
	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	conditions := readInput()

	fwk.Solution(1, calAllCombinations(conditions))
	fwk.Solution(2, calAllCombinations(unfoldAll(conditions)))
}

func calAllCombinations(conds []Condition) uint {
	var sum uint

	for _, cond := range conds {
		sum += bruteforce(cond)
	}

	return sum
}

func bruteforce(cond Condition) uint {
	matcher := state.New(cond.damagedSpringsSizes)

	cache := make(map[string]map[state.State]uint)
	combinations := bruteforceMatch(cond.springs, matcher, cache)
	// fmt.Printf("%s -- %d\n", cond.String(), combinations)

	return combinations
}

func bruteforceMatch(s []rune, matcher state.State, cache map[string]map[state.State]uint) uint {
	if matcher == nil {
		return 0
	}

	if stateCache, exists := cache[string(s)]; exists {
		if value, exists := stateCache[matcher]; exists {
			return value
		}
	}

	curState := matcher
	for i, c := range s {
		if c == '?' {
			nextS := s[i+1:]

			v := bruteforceMatch(nextS, curState.Next('#'), cache) +
				bruteforceMatch(nextS, curState.Next('.'), cache)

			createCacheEntry(cache, s, matcher, v)
			return v
		}

		curState = curState.Next(c)
		if curState == nil {
			return 0
		}
	}

	if curState.IsEnd() {
		return 1
	}
	return 0
}

func createCacheEntry(cache map[string]map[state.State]uint, s []rune, matcher state.State, value uint) {
	var stateCache map[state.State]uint
	var exists bool
	if stateCache, exists = cache[string(s)]; !exists {
		stateCache = make(map[state.State]uint)
		cache[string(s)] = stateCache
	}

	stateCache[matcher] = value
}

type Condition struct {
	springs             []rune
	damagedSpringsSizes []uint64
}

func (c Condition) String() string {
	return fmt.Sprintf("%s %v", string(c.springs), c.damagedSpringsSizes)
}

func unfoldAll(_conds []Condition) []Condition {
	conds := make([]Condition, len(_conds))
	for i, cond := range _conds {
		conds[i] = unfold(cond)
	}
	return conds
}

const mult = 5

func unfold(cond Condition) Condition {
	springs := make([]rune, len(cond.springs)*mult+mult-1)
	damaged := make([]uint64, len(cond.damagedSpringsSizes)*mult)
	for i := 0; i < mult; i++ {
		copy(springs[i*len(cond.springs)+i:], cond.springs)
		if i < mult-1 {
			springs[(i+1)*len(cond.springs)+i] = '?'
		}

		copy(damaged[i*len(cond.damagedSpringsSizes):], cond.damagedSpringsSizes)
	}

	return Condition{
		springs:             springs,
		damagedSpringsSizes: damaged,
	}
}

func readInput() []Condition {
	lines := fwk.ReadInputLines()
	conditions := make([]Condition, len(lines))

	for i, line := range lines {
		s := strings.Split(line, " ")
		numStrs := strings.Split(s[1], ",")
		nums := make([]uint64, len(numStrs))
		for i, str := range numStrs {
			var err error
			nums[i], err = strconv.ParseUint(str, 10, 64)
			if err != nil {
				panic(err)
			}
		}
		conditions[i] = Condition{
			springs:             []rune(s[0]),
			damagedSpringsSizes: nums,
		}
	}

	return conditions
}
