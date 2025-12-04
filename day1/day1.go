package day1

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type direction int

const (
	left  direction = -1
	right direction = 1
)

type rotation struct {
	direction direction
	steps     int
}

func (r rotation) String() string {
	b := new(strings.Builder)
	if r.direction == left {
		b.WriteRune('L')
	} else {
		b.WriteRune('R')
	}

	b.WriteString(strconv.Itoa(r.steps))

	return b.String()
}

func Run(cmd *cobra.Command, args []string) {
	rotations, err := parseRotations(args[0])
	if err != nil {
		fmt.Println("Unable to parse input:", err)
	}

	position := 50
	zeroLands := 0
	zeroClicks := 0

	for _, rotation := range rotations {
		newPosition := position

		for step := 0; step < rotation.steps; step++ {
			if rotation.direction == left {
				newPosition = (newPosition - 1 + 100) % 100
			} else {
				newPosition = (newPosition + 1) % 100
			}

			if newPosition == 0 {
				zeroClicks += 1
			}
		}

		if newPosition == 0 {
			zeroLands += 1
		}

		fmt.Println(position, "→", rotation, "→", newPosition)

		position = newPosition
	}

	fmt.Println("Step 1 Password:", zeroLands)
	fmt.Println("Step 2 Password:", zeroClicks)
}

func parseRotations(filename string) ([]rotation, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rotations := []rotation{}

	for _, line := range strings.Split(string(contents), "\n") {
		line = strings.TrimSpace(line)

		if len(line) < 2 {
			return nil, errors.New("invalid rotation: " + line)
		}

		dirChar := line[0]
		var dir direction
		switch dirChar {
		case 'L':
			dir = left
		case 'R':
			dir = right
		default:
			return nil, errors.New("invalid direction: " + string(dirChar))
		}

		stepsStr := line[1:]
		steps, err := strconv.Atoi(stepsStr)
		if err != nil {
			return nil, errors.New("invalid steps: " + stepsStr)
		}

		rotations = append(rotations, rotation{
			direction: dir,
			steps:     steps,
		})
	}

	return rotations, nil
}
