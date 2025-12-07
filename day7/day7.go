package day7

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type node int

const (
	empty node = iota
	start
	splitter
	beam
)

type xy struct {
	x, y int
}

func (pt xy) add(vec xy) xy {
	return xy{pt.x + vec.x, pt.y + vec.y}
}

var (
	up    xy = xy{0, -1}
	right xy = xy{1, 0}
	down  xy = xy{0, 1}
	left  xy = xy{-1, 0}
)

type grid struct {
	nodes         map[xy]node
	width, height int
}

func (g grid) startPoint() (xy, error) {
	for pt, node := range g.nodes {
		if node == start {
			return pt, nil
		}
	}

	return xy{}, errors.New("no start point")
}

// Run simlation to fill the grid with beams as they move from the start and hit splitters
func (g grid) simulateBeams() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			pt := xy{x, y}
			if g.isStruckByBeam(pt) {
				switch g.nodes[pt] {
				case empty:
					g.placeBeam(pt)
				case splitter:
					g.splitBeam(pt)
				}
			}
		}
	}
}

func (g grid) isStruckByBeam(pt xy) bool {
	if node, set := g.nodes[pt.add(up)]; set {
		return node == beam || node == start
	}

	return false
}

func (g grid) placeBeam(pt xy) {
	if node, set := g.nodes[pt]; set {
		if node == empty {
			g.nodes[pt] = beam
		}
	}
}

func (g grid) splitBeam(pt xy) {
	g.placeBeam(pt.add(left))
	g.placeBeam(pt.add(right))
}

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	// Create initial grid from input
	grid := parseInput(string(contents))

	// Let the beams flow!
	grid.simulateBeams()

	fmt.Println("Part 1:", part1(grid))
	fmt.Println("Part 2:", part2(grid))
}

// Part 1 result is how many times the beam hits a splitter
func part1(grid grid) int {
	splits := 0

	// Look for any splitter getting hit by the beam
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			pt := xy{x, y}
			if node, ok := grid.nodes[pt]; ok {
				if node == splitter && grid.isStruckByBeam(pt) {
					splits += 1
				}
			}
		}
	}

	return splits
}

// Part 2 result is how many possible paths the beam can take from the start to the end
func part2(grid grid) int {
	// Find start of beam
	start, _ := grid.startPoint()

	// Pass around a cache to avoid computing the same splitter multiple times
	cache := make(map[xy]int)

	return 1 + numPaths(grid, start, cache)
}

// Determine number of paths a beam can take from a point
func numPaths(grid grid, pt xy, cache map[xy]int) int {
	initialPt := pt

	if cnt, ok := cache[initialPt]; ok {
		return cnt
	}

	// Keep moving the pointer down (the beam only moves down)
	// until we hit a splitter or the end of the grid
	for {
		if node, ok := grid.nodes[pt]; ok {
			if node == splitter {
				// A new path is formed if a splitter is hit
				// The beam is then followed down the left side (considered the same beam) and the right side (the new beam)
				// Add any other new paths created when following these two beams
				lPt := pt.add(left)
				rPt := pt.add(right)

				cnt := 1 + numPaths(grid, lPt, cache) + numPaths(grid, rPt, cache)
				cache[initialPt] = cnt

				return cnt
			}
		} else {
			return 0
		}

		pt = pt.add(down)
	}
}

// Parse into file into initial grid
func parseInput(input string) grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	nodes := make(map[xy]node)
	height := 0
	width := 0

	for y, line := range lines {
		height += 1
		width = 0
		for x, char := range line {
			width += 1
			pt := xy{x, y}
			switch char {
			case '.':
				nodes[pt] = empty
			case '^':
				nodes[pt] = splitter
			case 'S':
				nodes[pt] = start
			}
		}
	}

	return grid{nodes, width, height}
}
