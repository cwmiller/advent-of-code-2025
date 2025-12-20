package day12

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type region struct {
	width, height int
	quantities    []int
}

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	regions := parseInput(string(contents))

	fmt.Println("Part 1:", part1(regions))
}

func part1(regions []region) int {
	total := 0

	for _, r := range regions {
		area := r.width * r.height
		totalShapes := 0
		for _, q := range r.quantities {
			totalShapes += q
		}
		required := totalShapes * 9

		if area >= required {
			total += 1
		}
	}

	return total
}

func parseInput(input string) []region {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	regions := []region{}

	lineRx := regexp.MustCompile(`(\d+)x(\d+): ([0-9 ]+)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		matches := lineRx.FindStringSubmatch(line)

		if len(matches) == 0 {
			continue
		}

		width, _ := strconv.Atoi(matches[1])
		height, _ := strconv.Atoi(matches[2])
		quantities := []int{}

		splits := strings.Split(matches[3], " ")
		for _, split := range splits {
			v, _ := strconv.Atoi(split)
			quantities = append(quantities, v)
		}

		region := region{
			width,
			height,
			quantities,
		}

		regions = append(regions, region)
	}

	return regions
}
