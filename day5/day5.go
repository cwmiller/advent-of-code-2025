package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type productId int

type productRange struct {
	min productId
	max productId
}

func (r productRange) Contains(value productId) bool {
	return value >= r.min && value <= r.max
}

func (r productRange) FitsIn(other productRange) bool {
	return r.min >= other.min && r.max <= other.max
}

func (r productRange) IntersectsStartOf(other productRange) bool {
	return r.min < other.min && r.max >= other.min
}

func (r productRange) IntersectsEndOf(other productRange) bool {
	return r.min < other.max && r.max >= other.max
}

func Run(cmd *cobra.Command, args []string) {

	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	freshProducts, availableProducts := parseInput(string(contents))

	part1Result := part1(freshProducts, availableProducts)
	part2Result := part2(freshProducts)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}

func part1(freshProducts []productRange, availableProducts []productId) int {
	total := 0

	for _, product := range availableProducts {
		isFresh := false
		for _, freshRange := range freshProducts {
			if freshRange.Contains(product) {
				isFresh = true
				break
			}
		}

		if isFresh {
			total += 1
		}
	}

	return total
}

func part2(freshProducts []productRange) int {
	total := 0
	adjusting := true

	for adjusting {
		freshProducts, adjusting = part2AdjustRanges(freshProducts)
	}

	for _, productRange := range freshProducts {
		total += (int(productRange.max) - int(productRange.min)) + 1
	}

	return total
}

func part2AdjustRanges(productRanges []productRange) ([]productRange, bool) {
	adjustedRanges := []productRange{}
	adjusted := false

	for _, productRange := range productRanges {
		for _, otherProductRange := range productRanges {
			if productRange == otherProductRange {
				continue
			}

			if productRange.FitsIn(otherProductRange) {
				adjusted = true
				goto skip
			}

			if productRange.IntersectsStartOf(otherProductRange) {
				productRange.max = max(productRange.max, otherProductRange.max)
				adjusted = true
			}

			if productRange.IntersectsEndOf(otherProductRange) {
				productRange.min = min(productRange.min, otherProductRange.min)
				adjusted = true
			}
		}

		adjustedRanges = append(adjustedRanges, productRange)
	skip:
	}

	return uniqueRanges(adjustedRanges), adjusted
}

func uniqueRanges(productRanges []productRange) []productRange {
	seen := make(map[productRange]bool)
	uniques := []productRange{}

	for _, item := range productRanges {
		if _, ok := seen[item]; !ok {
			uniques = append(uniques, item)
			seen[item] = true
		}
	}

	return uniques
}

func parseInput(input string) ([]productRange, []productId) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	freshProducts := []productRange{}
	availableProducts := []productId{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			min, _ := strconv.Atoi(parts[0])
			max, _ := strconv.Atoi(parts[1])
			freshProducts = append(freshProducts, productRange{productId(min), productId(max)})
		} else if line != "" {
			val, _ := strconv.Atoi(line)
			availableProducts = append(availableProducts, productId(val))
		}
	}

	return freshProducts, availableProducts
}
