package day8

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// 3d point in space
type xyz struct {
	x, y, z int
}

// Fuse box
type box struct {
	id        uuid.UUID
	pos       xyz
	circuitId uuid.UUID
}

// Euclidian distance between two boxes
type measurement struct {
	box1Id, box2Id uuid.UUID
	distance       float64
}

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	iterations, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	boxes := parseInput(string(contents))
	measurements := measureBoxes(boxes)

	fmt.Println("Part 1:", part1(boxes, measurements, iterations))
	fmt.Println("Part 2:", part2(boxes, measurements))
}

// Part 1 result is the product of the top 3 circuits after connecting the closest boxes
func part1(boxes map[uuid.UUID]box, measurements []measurement, iterations int) int {
	circuits := make(map[uuid.UUID]int)
	cnt := 0

	// Add all boxes as single circuits
	for _, box := range boxes {
		circuits[box.circuitId] = 1
	}

	for _, measurement := range measurements {
		if cnt == iterations {
			break
		}

		cnt++

		box1 := boxes[measurement.box1Id]
		box2 := boxes[measurement.box2Id]

		// Skip if already attached to the same circuit
		if box1.circuitId == box2.circuitId {
			continue
		}

		// Merge circuits into new circuit with new ID
		mergedCircuitId := uuid.New()

		// Update all boxes in the merged circuit with the new ID
		for boxId, box := range boxes {
			switch box.circuitId {
			case box1.circuitId:
				box.circuitId = mergedCircuitId
			case box2.circuitId:
				box.circuitId = mergedCircuitId
			}

			boxes[boxId] = box
		}

		// Update circuit box counts
		circuits[mergedCircuitId] = circuits[box1.circuitId] + circuits[box2.circuitId]
		delete(circuits, box1.circuitId)
		delete(circuits, box2.circuitId)
	}

	// Sort circuit sizes to get largest ones
	circuitSizes := []int{}
	for _, size := range circuits {
		circuitSizes = append(circuitSizes, size)
	}

	slices.Sort(circuitSizes)
	slices.Reverse(circuitSizes)

	result := circuitSizes[0]
	for i := 1; i < 3; i++ {
		result *= circuitSizes[i]
	}

	return result
}

// Part 2 connects all the remaining boxes together
// Result is the product of the X coordinates of the last two boxes to connect
func part2(boxes map[uuid.UUID]box, measurements []measurement) int {
	circuits := make(map[uuid.UUID]bool)

	// Populate list with all current circuits
	for _, box := range boxes {
		circuits[box.circuitId] = true
	}

	lastMergeProduct := 0

	// Keep connecting the closest fuse boxes until everything is connected in a single circuit
	for len(circuits) > 1 {
		for _, measurement := range measurements {
			box1 := boxes[measurement.box1Id]
			box2 := boxes[measurement.box2Id]

			// Skip if already attached to the same circuit
			if box1.circuitId == box2.circuitId {
				continue
			}

			// Merge circuits into new circuit with new ID
			mergedCircuitId := uuid.New()

			// Update all boxes in the merged circuit with the new ID
			for boxId, box := range boxes {
				switch box.circuitId {
				case box1.circuitId:
					box.circuitId = mergedCircuitId
				case box2.circuitId:
					box.circuitId = mergedCircuitId
				}

				boxes[boxId] = box
			}

			circuits[mergedCircuitId] = true
			delete(circuits, box1.circuitId)
			delete(circuits, box2.circuitId)

			lastMergeProduct = box1.pos.x * box2.pos.x
		}
	}

	return lastMergeProduct
}

// Get straight-line distance between two points
func distance(pt1 xyz, pt2 xyz) float64 {
	return math.Sqrt(math.Pow(float64(pt2.x-pt1.x), 2) + math.Pow(float64(pt2.y-pt1.y), 2) + math.Pow(float64(pt2.z-pt1.z), 2))
}

// Measure distances between all boxes
func measureBoxes(boxes map[uuid.UUID]box) []measurement {
	measurementsMap := make(map[string]measurement)

	for _, box1 := range boxes {
		for _, box2 := range boxes {
			if box1 == box2 {
				continue
			}

			var mapKey string
			if box1.id.String() < box2.id.String() {
				mapKey = box1.id.String() + "-" + box2.id.String()
			} else {
				mapKey = box2.id.String() + "-" + box1.id.String()
			}

			// Have we already measured these boxes?
			if _, set := measurementsMap[mapKey]; set {
				continue
			}

			distance := distance(box1.pos, box2.pos)
			measurement := measurement{box1.id, box2.id, distance}

			measurementsMap[mapKey] = measurement
		}
	}

	measurements := make([]measurement, 0, len(measurementsMap))
	for _, measurement := range measurementsMap {
		measurements = append(measurements, measurement)
	}

	// Sort by distance
	sort.Slice(measurements, func(i, j int) bool {
		return measurements[i].distance < measurements[j].distance
	})

	return measurements
}

// Parse input file into junction boxes
func parseInput(input string) map[uuid.UUID]box {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	//boxes := []box{}
	boxes := make(map[uuid.UUID]box)

	for _, line := range lines {
		parsed := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(parsed[0])
		y, _ := strconv.Atoi(parsed[1])
		z, _ := strconv.Atoi(parsed[2])
		box := box{
			id:        uuid.New(),
			pos:       xyz{x, y, z},
			circuitId: uuid.New(),
		}
		boxes[box.id] = box
	}

	return boxes
}
