package day6

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type problem struct {
	operation operation
	operands  []int
}

type operation int

const (
	addition operation = iota
	multiplication
)

func (p problem) solve() int {
	var result int

	// Perform the operation on all operands
	// First operand is used as the seed value
	for i, operand := range p.operands {
		if i == 0 {
			result = operand
		} else {
			if p.operation == addition {
				result += operand
			} else {
				result *= operand
			}
		}
	}

	return result
}

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	part1Problems := part1ParseInput(string(contents))
	fmt.Println("Part 1:", part1(part1Problems))

	part2Problems := part2ParseInput(string(contents))
	fmt.Println("Part 2:", part1(part2Problems))
}

func part1(problems []problem) int {
	total := 0
	for _, problem := range problems {
		total += problem.solve()
	}

	return total
}

func part1ParseInput(input string) []problem {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Problems are layed out in vertical columns
	// Keep track of the problems using their column index
	indexedProblems := make(map[int]problem)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		columns := strings.Split(line, " ")

		idx := 0

		for _, col := range columns {
			if col == "" {
				continue
			}

			if _, set := indexedProblems[idx]; !set {
				indexedProblems[idx] = problem{}
			}

			// Read either a operand or an operation from the column
			problem := indexedProblems[idx]
			switch col {
			case "+":
				problem.operation = addition
			case "*":
				problem.operation = multiplication
			default:
				operand, _ := strconv.Atoi(col)
				problem.operands = append(problem.operands, operand)
			}
			indexedProblems[idx] = problem

			idx++
		}
	}

	values := maps.Values(indexedProblems)
	return slices.Collect(values)
}

func part2ParseInput(input string) []problem {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Part 2 changes how numbers are read from the input
	// Instead of a number being read left-to-right, it's read top-to-bottom

	// Determine max width of each column
	columnWidths := columnWidths(lines)

	problems := make(map[int]problem)

	for problemIdx, problemWidth := range columnWidths {
		var problem problem

		// Determine which column the problem starts on
		colIdx := 0
		for i := range problemIdx {
			colIdx += columnWidths[i] + 1
		}

		// Operation is on the last line at the start of the column
		switch lines[len(lines)-1][colIdx] {
		case '+':
			problem.operation = addition
		case '*':
			problem.operation = multiplication
		}

		// Read each digit vertically and build out full number
		for digitIdx := range problemWidth {
			var digits string
			for lineIdx := 0; lineIdx < len(lines)-1; lineIdx++ {
				digit := string(lines[lineIdx][colIdx+digitIdx])

				if digit != " " {
					digits += digit
				}
			}

			if len(digits) > 0 {
				operand, err := strconv.Atoi(digits)
				if err != nil {
					panic(err)
				}

				problem.operands = append(problem.operands, operand)
			}
		}

		problems[problemIdx] = problem
	}

	values := maps.Values(problems)
	return slices.Collect(values)
}

func columnWidths(lines []string) map[int]int {
	widths := make(map[int]int)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		columns := strings.Split(line, " ")

		colIdx := 0

		for _, col := range columns {
			if col == "" {
				continue
			}

			if len(col) > widths[colIdx] {
				widths[colIdx] = len(col)
			}

			colIdx++
		}
	}

	return widths
}
