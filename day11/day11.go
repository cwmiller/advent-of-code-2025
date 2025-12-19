package day11

import (
	"fmt"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/spf13/cobra"
)

type device struct {
	label   string
	outputs []string
}

func Run(cmd *cobra.Command, args []string) {
	contents, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	devices := parseInput(string(contents))

	g := graph.New(graph.StringHash, graph.Directed())

	for _, d := range devices {
		g.AddVertex(d.label)
		for _, o := range d.outputs {
			g.AddVertex(o)
			g.AddEdge(d.label, o)
		}
	}

	am, _ := g.AdjacencyMap()

	fmt.Println("Part 1:", part1(am))
	fmt.Println("Part 2:", part2(am))
}

func part1(am map[string]map[string]graph.Edge[string]) int {
	return countPaths(am, "you", "out")
}

func part2(am map[string]map[string]graph.Edge[string]) int {
	return countPaths(am, "svr", "fft") *
		countPaths(am, "fft", "dac") *
		countPaths(am, "dac", "out")
}

func countPaths(am map[string]map[string]graph.Edge[string], start, end string) int {
	edges := make(map[string]int)
	reachable := make(map[string]bool)
	reachable[start] = true

	queue := []string{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for neighbor := range am[node] {
			if !reachable[neighbor] {
				reachable[neighbor] = true
				queue = append(queue, neighbor)
			}
			edges[neighbor]++
		}
	}

	dp := make(map[string]int)
	dp[start] = 1

	topo := []string{}
	queue = []string{}

	for node := range reachable {
		if edges[node] == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		topo = append(topo, node)

		for neighbor := range am[node] {
			if !reachable[neighbor] {
				continue
			}

			edges[neighbor]--
			if edges[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	for _, node := range topo {
		for neighbor := range am[node] {
			if reachable[neighbor] {
				dp[neighbor] += dp[node]
			}
		}
	}

	return dp[end]
}

func parseInput(input string) []device {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	devices := []device{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ": ")
		label := parts[0]
		outputs := strings.Split(parts[1], " ")

		d := device{
			label,
			outputs,
		}

		devices = append(devices, d)
	}

	return devices
}
