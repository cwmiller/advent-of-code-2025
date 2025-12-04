package day2

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type idRange struct {
	start int
	end   int
}

func Run(cmd *cobra.Command, args []string) {
	ranges, err := parseRanges(args[0])
	if err != nil {
		fmt.Println("Unable to parse input:", err)
	}

	part1Result := part1(ranges)
	part2Result := part2(ranges)

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}

func part1(ranges []idRange) int {
	total := 0

	for _, idRange := range ranges {
		for id := idRange.start; id <= idRange.end; id++ {
			str := strconv.Itoa(id)
			if part1IsInvalid(str) {
				total += id
			}
		}
	}

	return total
}

func part1IsInvalid(id string) bool {
	if len(id)&1 != 0 {
		return false
	}

	midpoint := len(id) / 2
	firstHalf := id[:midpoint]
	secondHalf := id[midpoint:]

	return firstHalf == secondHalf
}

func part2(ranges []idRange) int {
	total := 0

	for _, idRange := range ranges {
		for id := idRange.start; id <= idRange.end; id++ {
			str := strconv.Itoa(id)
			if part2IsInvalid(str) {
				val, _ := strconv.Atoi(str)
				total += val
			}
		}
	}

	return total
}

func part2IsInvalid(id string) bool {
	for maskLength := len(id) / 2; maskLength > 0; maskLength-- {
		mask := id[:maskLength]

		test := id
		for len(test) > 0 {
			if strings.HasPrefix(test, mask) {
				test = test[maskLength:]
			} else {
				break
			}
		}

		if len(test) == 0 {
			return true
		}
	}

	return false
}

func parseRanges(filename string) ([]idRange, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ranges := []idRange{}

	for _, part := range strings.Split(strings.TrimSpace(string(contents)), ",") {
		bounds := strings.Split(part, "-")
		if len(bounds) != 2 {
			return nil, errors.New("invalid range: " + part)
		}
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, errors.New("invalid start of range: " + bounds[0])
		}
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, errors.New("invalid end of range: " + bounds[1])
		}

		ranges = append(ranges, idRange{
			start: start,
			end:   end,
		})
	}

	return ranges, nil
}
