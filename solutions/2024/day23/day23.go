package day23

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var inputPath string = "inputs/2024/day23.txt"

func input() (map[string][]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open %s: %v", inputPath, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	computerConnections := map[string][]string{}

	for scanner.Scan() {
		computersFromTo := strings.Split(scanner.Text(), "-")

		computerConnections[computersFromTo[0]] = append(computerConnections[computersFromTo[0]], computersFromTo[1])
		computerConnections[computersFromTo[1]] = append(computerConnections[computersFromTo[1]], computersFromTo[0])
	}

	return computerConnections, scanner.Err()
}

func Part1() error {
	computerConnections, err := input()
	if err != nil {
		return err
	}

	cycles := make(map[string][]string)

	for computer := range computerConnections {
		if computer[0] != 't' {
			continue
		}

		for _, connectedComputer := range computerConnections[computer] {
			for _, computerConnectedToConnectedComputer := range computerConnections[connectedComputer] {
				if slices.Contains(computerConnections[computerConnectedToConnectedComputer], computer) {
					computerGroup := []string{computer, connectedComputer, computerConnectedToConnectedComputer}
					slices.Sort(computerGroup)
					key := fmt.Sprintf("%s-%s-%s", computerGroup[0], computerGroup[1], computerGroup[2])
					cycles[key] = computerGroup
				}
			}
		}
	}

	fmt.Println("Day23a:", len(cycles))

	return nil
}

func Part2() error {
	computerConnections, err := input()
	if err != nil {
		return err
	}

	computers := make([]string, len(computerConnections))

	for computer := range computerConnections {
		computers = append(computers, computer)
	}

	maxLenComputerList := []string{}

	BronKerbosch([]string{}, computers, []string{}, &maxLenComputerList, computerConnections)

	slices.Sort(maxLenComputerList)

	fmt.Println("Day23b:", strings.Join(maxLenComputerList, ","))

	return nil
}

func BronKerbosch(computersInChain, possibleAdditions, excludedComputers []string, maxLenComputerList *[]string, computerConnections map[string][]string) {
	if len(possibleAdditions) == 0 && len(excludedComputers) == 0 && len(*maxLenComputerList) < len(computersInChain) {
		newComputersInChain := make([]string, len(computersInChain))
		copy(newComputersInChain, computersInChain)
		*maxLenComputerList = slices.Clone(newComputersInChain)

		return
	}

	possibleAdditionsCopy := make([]string, len(possibleAdditions))
	copy(possibleAdditionsCopy, possibleAdditions)

	for _, possibleComputerToAdd := range possibleAdditionsCopy {
		computersInChain := append(computersInChain, possibleComputerToAdd)
		newComputerConnections := computerConnections[possibleComputerToAdd]

		newPossibleAdditions := findConnectable(possibleAdditions, newComputerConnections)
		newExcludedComputers := findConnectable(excludedComputers, newComputerConnections)

		BronKerbosch(computersInChain, newPossibleAdditions, newExcludedComputers, maxLenComputerList, computerConnections)

		computerIndex := slices.Index(possibleAdditions, possibleComputerToAdd)
		possibleAdditions = slices.Delete(possibleAdditions, computerIndex, computerIndex+1)

		excludedComputers = append(excludedComputers, possibleComputerToAdd)
	}
}

func findConnectable(computers, connections []string) []string {
	connectableMap := map[string]bool{}

	for _, val := range computers {
		connectableMap[val] = true
	}

	connectable := []string{}

	for _, val := range connections {
		if connectableMap[val] {
			connectable = append(connectable, val)
		}
	}

	return connectable
}
