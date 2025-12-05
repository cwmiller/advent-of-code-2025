package day4

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type point struct {
	x int
	y int
}

func (p point) Add(v vec) point {
	return point{p.x + v.x, p.y + v.y}
}

type vec struct {
	x int
	y int
}

var (
	n vec = vec{0, -1}
	e vec = vec{1, 0}
	s vec = vec{0, 1}
	w vec = vec{-1, 0}

	nw vec = vec{-1, -1}
	ne vec = vec{1, -1}
	sw vec = vec{-1, 1}
	se vec = vec{1, 1}

	directions = []vec{n, ne, e, se, s, sw, w, nw}
)

type node int

const (
	empty node = iota
	paper
)

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	grid, width, height := parseInput(string(contents))

	part1Result := part1(grid, width, height)
	part2Result := part2(grid, width, height)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}

func part1(grid map[point]node, width int, height int) int {
	count := 0
	for y := range height {
		for x := range width {
			pt := point{x, y}

			if accessible(grid, pt) {
				count += 1
			}
		}
	}

	return count
}

func part2(grid map[point]node, width int, height int) int {
	count := 0

	for {
		removed := removeAccessible(grid, width, height)

		if removed == 0 {
			break
		}

		count += removed
	}

	return count
}

func removeAccessible(grid map[point]node, width int, height int) int {
	removed := 0

	for y := range height {
		for x := range width {
			pt := point{x, y}

			if accessible(grid, pt) {
				grid[pt] = empty
				removed += 1
			}
		}
	}

	return removed
}

func accessible(grid map[point]node, pt point) bool {
	if node, exists := grid[pt]; !exists || node != paper {
		return false
	}

	count := 0
	for _, dir := range directions {
		adjPt := pt.Add(dir)
		if node, exists := grid[adjPt]; exists && node == paper {
			count += 1
		}
	}

	return count < 4
}

func parseInput(input string) (map[point]node, int, int) {
	grid := make(map[point]node)
	height := 0
	width := 0

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		height += 1
		width = 0
		for x, char := range line {
			width += 1
			pt := point{x, y}
			switch char {
			case '.':
				grid[pt] = empty
			case '@':
				grid[pt] = paper
			}
		}
	}

	return grid, width, height
}
