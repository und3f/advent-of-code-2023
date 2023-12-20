package main

import (
	"bytes"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/und3f/aoc/2023/fwk"
)

func main() {
	modules := parseInput()
	part1(modules)
	part2(modules)
}

const PART1 = 1000

func part1(modules map[string]Module) {
	emulator := NewEmulator(modules)

	var low, high, sum uint
	for i := 0; i < PART1; i++ {
		pulses := emulator.PressButton()
		sum += uint(len(pulses))
		for _, s := range pulses {
			if s.Value {
				high++
			} else {
				low++
			}
		}
	}
	fwk.Solution(1, high*low)
}

const Part2wire = "rx"
const Part2Value = false

func part2(modules map[string]Module) {
	var components []string
	wire := Part2wire
	reqValue := Part2Value
	for name, module := range modules {
		if slices.Contains(module.Outputs, wire) {
			wire = name
			reqValue = !Part2Value
			break
		}
	}

	for name, module := range modules {
		if slices.Contains(module.Outputs, wire) {
			components = append(components, name)
		}
	}

	componentCycle := make([]int, len(components))
	for i, comp := range components {
		emulator := NewEmulator(modules)
		presses := 0

		found := false
		for !found {
			pulses := emulator.PressButton()
			presses++
			for _, s := range pulses {
				if strings.Compare(comp, s.Src) == 0 && s.Value == reqValue {
					found = true
					componentCycle[i] = presses
					break
				}
			}
		}
	}
	fwk.Solution(2, fwk.LCM(componentCycle...))
}

const startModule = "broadcaster"

type Signal struct {
	Src   string
	Value bool
	Dest  string
}

type Emulator struct {
	flipFlopState   map[string]bool
	conjuctionState map[string]map[string]bool
	modules         map[string]Module
}

func NewEmulator(modules map[string]Module) *Emulator {
	flipFlopState := make(map[string]bool)
	conjuctionState := make(map[string]map[string]bool)
	for moduleName, module := range modules {
		switch module.Type {
		case '%':
			flipFlopState[moduleName] = false
		case '&':
			conjuctionState[moduleName] = make(map[string]bool)
		}
	}
	for srcModule, module := range modules {
		for _, dstModule := range module.Outputs {
			if modules[dstModule].Type == '&' {
				conjuctionState[dstModule][srcModule] = false
			}
		}
	}
	return &Emulator{
		modules:         modules,
		conjuctionState: conjuctionState,
		flipFlopState:   flipFlopState,
	}
}

func (e *Emulator) PressButton() []Signal {
	signals := []Signal{
		{Src: "button", Dest: startModule, Value: false},
	}
	var signalHistory []Signal

	for len(signals) > 0 {
		var nextSignals []Signal

		for _, signal := range signals {
			signalHistory = append(signalHistory, signal)

			signalValue := signal.Value
			moduleName := signal.Dest
			module := e.modules[moduleName]

			switch module.Type {
			case '%':
				if signalValue == true {
					// Ignore
					continue
				}
				signalValue = !e.flipFlopState[moduleName]
				e.flipFlopState[moduleName] = signalValue
			case '&':
				e.conjuctionState[moduleName][signal.Src] = signalValue
				signalValue = true
				for _, v := range e.conjuctionState[moduleName] {
					signalValue = signalValue && v
				}
				signalValue = !signalValue
			}

			for _, destModule := range module.Outputs {
				nextSignals = append(nextSignals, Signal{
					Src:   moduleName,
					Dest:  destModule,
					Value: signalValue,
				})
			}
		}

		signals = nextSignals
	}

	return signalHistory
}

func (e *Emulator) String() string {
	var buf bytes.Buffer

	buf.WriteString("Flip Flops: ")
	var flipFlops []string
	for m := range e.flipFlopState {
		flipFlops = append(flipFlops, m)
	}
	sort.Strings(flipFlops)
	for _, m := range flipFlops {
		if e.flipFlopState[m] {
			buf.WriteString(fmt.Sprintf("%s:%s ", m, "1"))
		}
	}
	buf.WriteString("\n")

	buf.WriteString("Conjuctions: ")
	var conjuctions []string
	for m := range e.conjuctionState {
		conjuctions = append(conjuctions, m)
	}
	sort.Strings(conjuctions)
	for _, m := range conjuctions {
		wires := []string{}

		for w := range e.conjuctionState[m] {
			if e.conjuctionState[m][w] {
				wires = append(wires, w)
			}
		}
		if len(wires) == 0 {
			continue
		}

		buf.WriteString(m)
		buf.WriteString("(")

		sort.Strings(wires)
		for _, w := range wires {
			buf.WriteString(fmt.Sprintf("%s:%s ", w, "1"))
		}
		buf.WriteString(") ")
	}
	buf.WriteString("\n")

	return buf.String()
}

type Module struct {
	Type    rune
	Outputs []string
}

func parseInput() map[string]Module {
	lines := fwk.ReadInputLines()
	r := make(map[string]Module)
	for _, line := range lines {
		s := strings.Split(line, " -> ")

		var moduleType rune
		moduleName := s[0]
		if slices.Contains([]rune{'%', '&'}, []rune(moduleName)[0]) {
			moduleType = []rune(moduleName)[0]
			moduleName = moduleName[1:]
		}
		r[moduleName] = Module{
			Type:    moduleType,
			Outputs: strings.Split(s[1], ", "),
		}
	}

	return r
}
