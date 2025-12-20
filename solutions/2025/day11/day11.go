package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputPath = "inputs/2025/day11.txt"

type Graph map[string][]string

func parseInput(filePath string) (Graph, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	graph := make(Graph)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}

		device := parts[0]
		outputs := strings.Fields(parts[1])
		graph[device] = outputs
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}

	return graph, nil
}

func dfs(graph Graph, current, end string, memo map[string]int) int {
	if current == end {
		return 1
	}

	if result, ok := memo[current]; ok {
		return result
	}

	count := 0
	for _, neighbor := range graph[current] {
		count += dfs(graph, neighbor, end, memo)
	}

	memo[current] = count

	return count
}

func countPaths(graph Graph, start, end string) int {
	return dfs(graph, start, end, make(map[string]int))
}

func Part1() error {
	graph, err := parseInput(inputPath)
	if err != nil {
		return err
	}

	fmt.Println("Day11a:", countPaths(graph, "you", "out"))

	return nil
}

func countPathsWithRequiredNodes(graph Graph, start, end string) int {
	return dfsWithRequired(graph, start, end, make(map[state]int), false, false)
}

type state struct {
	name string
	fft  bool
	dac  bool
}

func dfsWithRequired(graph Graph, current, end string, memo map[state]int, visitedFFT, visitedDAC bool) int {
	if current == end {
		if visitedFFT && visitedDAC {
			return 1
		}

		return 0
	}

	state := state{name: current, fft: visitedFFT, dac: visitedDAC}

	if result, ok := memo[state]; ok {
		return result
	}

	count := 0
	for _, neighbor := range graph[current] {
		newFFT := visitedFFT || neighbor == "fft"
		newDAC := visitedDAC || neighbor == "dac"
		count += dfsWithRequired(graph, neighbor, end, memo, newFFT, newDAC)
	}

	memo[state] = count

	return count
}

func Part2() error {
	graph, err := parseInput(inputPath)
	if err != nil {
		return err
	}

	fmt.Println("Day11b:", countPathsWithRequiredNodes(graph, "svr", "out"))

	return nil
}
