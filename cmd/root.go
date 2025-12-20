package cmd

import (
	"os"

	"github.com/cwmiller/advent-of-code-2025/day1"
	"github.com/cwmiller/advent-of-code-2025/day10"
	"github.com/cwmiller/advent-of-code-2025/day11"
	"github.com/cwmiller/advent-of-code-2025/day12"
	"github.com/cwmiller/advent-of-code-2025/day2"
	"github.com/cwmiller/advent-of-code-2025/day3"
	"github.com/cwmiller/advent-of-code-2025/day4"
	"github.com/cwmiller/advent-of-code-2025/day5"
	"github.com/cwmiller/advent-of-code-2025/day6"
	"github.com/cwmiller/advent-of-code-2025/day7"
	"github.com/cwmiller/advent-of-code-2025/day8"
	"github.com/cwmiller/advent-of-code-2025/day9"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "advent-of-code-2025",
	Short: "Advent of Code 2025 Solutions",
	Long:  `Solutions for the Advent of Code 2025 programming challenge.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var day1Cmd = &cobra.Command{
	Use:   "day1 [input file]",
	Short: "Day 1: Secret Entrance",
	Args:  cobra.ExactArgs(1),
	Run:   day1.Run,
}

var day2Cmd = &cobra.Command{
	Use:   "day2 [input file]",
	Short: "Day 2: Gift Shop",
	Args:  cobra.ExactArgs(1),
	Run:   day2.Run,
}

var day3Cmd = &cobra.Command{
	Use:   "day3 [input file]",
	Short: "Day 3: Lobby",
	Args:  cobra.ExactArgs(1),
	Run:   day3.Run,
}

var day4Cmd = &cobra.Command{
	Use:   "day4 [input file]",
	Short: "Day 4: Printing Department",
	Args:  cobra.ExactArgs(1),
	Run:   day4.Run,
}

var day5Cmd = &cobra.Command{
	Use:   "day5 [input file]",
	Short: "Day 5: Cafeteria",
	Args:  cobra.ExactArgs(1),
	Run:   day5.Run,
}

var day6Cmd = &cobra.Command{
	Use:   "day6 [input file]",
	Short: "Day 6: Cafeteria",
	Args:  cobra.ExactArgs(1),
	Run:   day6.Run,
}

var day7Cmd = &cobra.Command{
	Use:   "day7 [input file]",
	Short: "Day 7: Laboratories",
	Args:  cobra.ExactArgs(1),
	Run:   day7.Run,
}

var day8Cmd = &cobra.Command{
	Use:   "day8 [input file] [iterations]",
	Short: "Day 8: Playground",
	Args:  cobra.ExactArgs(2),
	Run:   day8.Run,
}

var day9Cmd = &cobra.Command{
	Use:   "day9 [input file]",
	Short: "Day 9: Movie Theater",
	Args:  cobra.ExactArgs(1),
	Run:   day9.Run,
}

var day10Cmd = &cobra.Command{
	Use:   "day10 [input file]",
	Short: "Day 10: Factory",
	Args:  cobra.ExactArgs(1),
	Run:   day10.Run,
}

var day11Cmd = &cobra.Command{
	Use:   "day11 [input file]",
	Short: "Day 11: Reactor",
	Args:  cobra.ExactArgs(1),
	Run:   day11.Run,
}

var day12Cmd = &cobra.Command{
	Use:   "day12 [input file]",
	Short: "Day 12: Christmas Tree Farm",
	Args:  cobra.ExactArgs(1),
	Run:   day12.Run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent-of-code-2025.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(day1Cmd)
	rootCmd.AddCommand(day2Cmd)
	rootCmd.AddCommand(day3Cmd)
	rootCmd.AddCommand(day4Cmd)
	rootCmd.AddCommand(day5Cmd)
	rootCmd.AddCommand(day6Cmd)
	rootCmd.AddCommand(day7Cmd)
	rootCmd.AddCommand(day8Cmd)
	rootCmd.AddCommand(day9Cmd)
	rootCmd.AddCommand(day10Cmd)
	rootCmd.AddCommand(day11Cmd)
	rootCmd.AddCommand(day12Cmd)
}
