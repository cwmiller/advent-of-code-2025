package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type xy struct {
	x, y int
}

func (xy1 xy) add(xy2 xy) xy {
	return xy{
		x: xy1.x + xy2.x,
		y: xy1.y + xy2.y,
	}
}

var (
	up    = xy{0, -1}
	right = xy{1, 0}
	down  = xy{0, 1}
	left  = xy{-1, 0}
)

type pair struct {
	pt1, pt2 xy
}

type tile int

const (
	red tile = iota
	green
	border
	answer
)

type floor struct {
	grid          map[xy]tile
	width, height int
}

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	points := parseInput(string(contents))

	fmt.Println("Part 1:", part1(points))
	fmt.Println("Part 2:", part2(points))
}

func part1(points []xy) int {
	// Create a unique list of point pairs
	pairs := []pair{}

	maxArea := 0

	for i, pt1 := range points {
		for j, pt2 := range points {
			if j <= i {
				continue
			}

			pairs = append(pairs, pair{pt1, pt2})
		}
	}

	for _, pair := range pairs {
		area := area(pair)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func part2(points []xy) int {
	// Shrink grid to simplify the amount of pixels needed to process
	// Keep a map of the original point to the simplified one
	pointMap := make(map[xy]xy, len(points))
	simplifiedPoints := []xy{}

	for _, pt := range points {
		simplifiedPt := xy{pt.x / 10, pt.y / 10}
		simplifiedPoints = append(simplifiedPoints, simplifiedPt)
		pointMap[simplifiedPt] = pt
	}

	grid := make(map[xy]tile, 500000)
	width := 0
	height := 0

	// The points draw a line where each listed point is a red tile and all points in between are green
	for i := 0; i < len(simplifiedPoints); i += 1 {
		j := i + 1
		if j == len(simplifiedPoints) {
			j = 0
		}

		pt1 := simplifiedPoints[i]
		pt2 := simplifiedPoints[j]

		if pt1.x > width {
			width = pt1.x
		}

		if pt1.y > height {
			height = pt1.y
		}

		grid[pt1] = red
		grid[pt2] = red

		var v xy
		if pt1.x == pt2.x {
			if pt1.y < pt2.y {
				v = xy{0, 1}
			} else {
				v = xy{0, -1}
			}
		} else if pt1.y == pt2.y {
			if pt1.x < pt2.x {
				v = xy{1, 0}
			} else {
				v = xy{-1, 0}
			}
		} else {
			panic(fmt.Sprintf("Cannot connect %d,%d to %d,%d", pt1.x, pt1.y, pt2.x, pt2.y))
		}

		for pt := pt1.add(v); pt != pt2; pt = pt.add(v) {
			grid[pt] = green
		}
	}

	width += 1
	height += 1

	floor := floor{grid, width, height}

	// Mark pixels surrounding the border
	surroundBorder(floor)

	// Create pairing of all simplified points
	allPairs := []pair{}

	for i, pt1 := range simplifiedPoints {
		for j, pt2 := range simplifiedPoints {
			if j <= i {
				continue
			}

			allPairs = append(allPairs, pair{pt1, pt2})
		}
	}

	maxArea := 0
	for _, ptPair := range allPairs {
		if checkRect(floor, ptPair) {
			// Get original points
			pt1 := pointMap[ptPair.pt1]
			pt2 := pointMap[ptPair.pt2]
			area := area(pair{pt1, pt2})

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

// Get area of rectangle formed by two opposing corners
func area(pair pair) int {
	l := pair.pt2.x - pair.pt1.x
	if l < 0 {
		l *= -1
	}
	l += 1

	w := pair.pt2.y - pair.pt1.y
	if w < 0 {
		w *= -1
	}
	w += 1

	return l * w
}

// Create a stroke around the outside of the tiles
// This stroke is used to check if a rectangle goes outside the tiles
func surroundBorder(floor floor) {
	for x := 0; x < floor.width; x += 1 {
		// Send beam downward
		for y := 0; y < floor.height; y += 1 {
			pt := xy{x, y}

			// Check left
			if tile, set := floor.grid[pt.add(left)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
				}
			}

			// Check right
			if tile, set := floor.grid[pt.add(right)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
				}
			}

			// Check down. stop beam if found
			if tile, set := floor.grid[pt.add(down)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
					break
				}
			}
		}

		// Beam upward from bottom
		for y := floor.height; y > 0; y -= 1 {
			pt := xy{x, y}

			// Check left
			if tile, set := floor.grid[pt.add(left)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
				}
			}

			// Check right
			if tile, set := floor.grid[pt.add(right)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
				}
			}

			// Check up. Stop beam if hit
			if tile, set := floor.grid[pt.add(up)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
					break
				}
			}
		}
	}

	for y := 0; y < floor.height; y += 1 {
		for x := 0; x < floor.width; x += 1 {
			pt := xy{x, y}

			// Check up
			if tile, set := floor.grid[pt.add(up)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
				}
			}

			// Check down
			if tile, set := floor.grid[pt.add(down)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
				}
			}

			// Check right. Stop beam if hit
			if tile, set := floor.grid[pt.add(right)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
					break
				}
			}
		}

		for x := floor.width; x > 0; x -= 1 {
			pt := xy{x, y}

			// Check up
			if tile, set := floor.grid[pt.add(up)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
				}
			}

			// Check down
			if tile, set := floor.grid[pt.add(down)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
				}
			}

			// Check left. Stop beam if hit
			if tile, set := floor.grid[pt.add(left)]; set {
				if tile == green || tile == red {
					floor.grid[pt] = border
					break
				}
			}
		}
	}
}

// Check if a rectangle drawn using the given pair crosses the border
func checkRect(floor floor, pair pair) bool {
	minY := min(pair.pt1.y, pair.pt2.y)
	maxY := max(pair.pt1.y, pair.pt2.y)
	minX := min(pair.pt1.x, pair.pt2.x)
	maxX := max(pair.pt1.x, pair.pt2.x)

	for y := minY; y <= maxY; y += 1 {
		if floor.grid[xy{minX, y}] == border {
			return false
		}

		if floor.grid[xy{maxX, y}] == border {
			return false
		}
	}

	for x := minX; x <= maxX; x += 1 {
		if floor.grid[xy{x, minY}] == border {
			return false
		}

		if floor.grid[xy{x, maxY}] == border {
			return false
		}
	}

	return true
}

// Parse input file into points
func parseInput(input string) []xy {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pts := []xy{}

	for _, line := range lines {
		parsed := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(parsed[0])
		y, _ := strconv.Atoi(parsed[1])
		pts = append(pts, xy{x, y})
	}

	return pts
}
