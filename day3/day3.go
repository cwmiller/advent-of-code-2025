package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	part1TotalJoltage := 0
	part2TotalJoltage := 0

	for _, line := range strings.Split(string(contents), "\n") {
		part1TotalJoltage += part1MaxJoltage(line)
		part2TotalJoltage += part2MaxJoltage(line)
	}

	fmt.Println("Part 1:", part1TotalJoltage)
	fmt.Println("Part 2:", part2TotalJoltage)
}

func part1MaxJoltage(bank string) int {
	max := 0

	for i, char := range bank {
		for j := i + 1; j < len(bank); j++ {
			nextChar := bank[j]
			value, _ := strconv.Atoi(string(char) + string(nextChar))

			if value > max {
				max = value
			}
		}
	}

	return max
}

func part2MaxJoltage(bank string) int {
	max := 0
	remainingBank := strings.TrimSpace(bank)
	working := []string{}

	for len(working) < 12 {
		// Record the lowest index for each value found in the remaining bank
		indexes := make(map[int]int)
		for i, char := range remainingBank {
			val, _ := strconv.Atoi(string(char))
			if _, set := indexes[val]; !set {
				indexes[val] = i
			}
		}

		for val := 9; val >= 0; val-- {
			if index, set := indexes[val]; set {
				if len(working)+len(remainingBank[index:]) >= 12 {
					working = append(working, strconv.Itoa(val))
					remainingBank = remainingBank[index+1:]
					break
				}
			}
		}

	}

	max, err := strconv.Atoi(strings.Join(working, ""))

	if err != nil {
		panic(err)
	}

	return max
}
