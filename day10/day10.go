package day10

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/spf13/cobra"
)

type machineIndicatorState []bool
type machineJoltageState []int

func (state machineIndicatorState) String() string {
	var indicators string
	for _, indicator := range state {
		if indicator {
			indicators += "#"
		} else {
			indicators += "."
		}
	}

	return indicators
}

func (state machineJoltageState) String() string {
	var joltages []string
	for _, joltage := range state {
		joltages = append(joltages, strconv.Itoa(joltage))
	}

	return strings.Join(joltages, ",")
}

// Compare two joltage states
// If all joltages in `state` and `other` match, return 0
// If any joltages in `other` exceed `state`, return 1
// Else return -1
func (state machineJoltageState) compare(other machineJoltageState) int {
	equals := true

	for i, j := range state {
		if other[i] > j {
			return 1
		}

		if other[i] != j {
			equals = false
		}
	}

	if equals {
		return 0
	} else {
		return -1
	}
}

type machine struct {
	indicatorTarget machineIndicatorState
	indicatorState  machineIndicatorState
	joltageTarget   machineJoltageState
	joltageState    machineJoltageState
	buttons         []button
}

type button struct {
	wires []int
}

func (btn button) pressIndicator(state machineIndicatorState) machineIndicatorState {
	newState := make(machineIndicatorState, len(state))
	copy(newState, state)

	for _, wire := range btn.wires {
		newState[wire] = !state[wire]
	}

	return newState
}

func (btn button) pressJoltage(state machineJoltageState) machineJoltageState {
	newState := make(machineJoltageState, len(state))
	copy(newState, state)

	for _, wire := range btn.wires {
		newState[wire] += 1
	}

	return newState
}

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	machines := parseInput(string(contents))

	fmt.Println("Part 1:", part1(machines))
	fmt.Println("Run day10-part2.cs for part 2")
}

func part1(machines []machine) int {
	result := 0

	for _, machine := range machines {
		sourceState := machine.indicatorState
		targetState := machine.indicatorTarget

		g := part1Graph(machine)
		path, err := graph.ShortestPath(g, sourceState.String(), targetState.String())

		if err != nil {
			panic(err)
		}

		result += (len(path) - 1)
	}

	return result
}

func part1Graph(m machine) graph.Graph[string, machineIndicatorState] {
	g := graph.New(func(ms machineIndicatorState) string {
		return ms.String()
	})

	// Add source state
	g.AddVertex(m.indicatorState)

	part1PressButtons(g, m.indicatorState, m.buttons, m.indicatorTarget)

	return g
}

func part1PressButtons(g graph.Graph[string, machineIndicatorState], state machineIndicatorState, btns []button, target machineIndicatorState) {
	queue := []machineIndicatorState{}

	for _, btn := range btns {
		nextState := btn.pressIndicator(state)
		err := g.AddVertex(nextState)

		if err == nil {
			queue = append(queue, nextState)
		}
		g.AddEdge(state.String(), nextState.String())

		// End now if we've hit the target
		if nextState.String() == target.String() {
			return
		}
	}

	for _, nextState := range queue {
		part1PressButtons(g, nextState, btns, target)
	}
}

// Parse input into machines
func parseInput(input string) []machine {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	machines := []machine{}

	indicatorPattern := regexp.MustCompile(`\[([\.#]+)\]`)
	buttonsPattern := regexp.MustCompile(`\(([\d,]+)\)`)
	joltagesPattern := regexp.MustCompile(`\{([\d,]+)\}`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		indicatorsResult := indicatorPattern.FindStringSubmatch(line)
		joltagesResult := joltagesPattern.FindStringSubmatch(line)
		buttonResults := buttonsPattern.FindAllStringSubmatch(line, -1)

		indicators := []bool{}
		joltages := []int{}
		buttons := []button{}

		for _, light := range indicatorsResult[1] {
			indicators = append(indicators, light == '#')
		}

		joltagesSplits := strings.Split(joltagesResult[1], ",")
		for _, joltage := range joltagesSplits {
			joltageVal, _ := strconv.Atoi(joltage)
			joltages = append(joltages, joltageVal)
		}

		/*
			for _, joltage := range joltagesResult[1] {
				joltageVal, err := strconv.Atoi(string(joltage))
				if err != nil {
					panic(err)
				}
				joltages = append(joltages, joltageVal)
			}*/

		for _, buttonResult := range buttonResults {
			parsed := strings.Split(buttonResult[1], ",")
			wires := []int{}
			for _, wire := range parsed {
				v, _ := strconv.Atoi(wire)
				wires = append(wires, v)
			}
			buttons = append(buttons, button{wires})
		}

		indicatorState := make(machineIndicatorState, len(indicators))
		joltageState := make(machineJoltageState, len(joltages))

		machine := machine{
			indicatorState:  indicatorState,
			indicatorTarget: indicators,
			joltageState:    joltageState,
			joltageTarget:   joltages,
			buttons:         buttons,
		}
		machines = append(machines, machine)
	}

	return machines
}
