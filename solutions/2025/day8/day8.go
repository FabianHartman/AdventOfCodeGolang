package day8

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputPath = "inputs/2025/day8.txt"

type JunctionBox struct {
	X int
	Y int
	Z int
}

type JunctionBoxDistance struct {
	From     JunctionBox
	To       JunctionBox
	Distance int
}

func squared(n int) int {
	return n * n
}

func (this JunctionBox) distance(that JunctionBox) JunctionBoxDistance {
	return JunctionBoxDistance{
		From:     this,
		To:       that,
		Distance: squared(this.X-that.X) + squared(this.Y-that.Y) + squared(this.Z-that.Z),
	}
}

func isSameCircuit(circuitA, circuitB map[JunctionBox]struct{}) bool {
	if len(circuitA) != len(circuitB) {
		return false
	}
	for box := range circuitA {
		_, ok := circuitB[box]
		if !ok {
			return false
		}
	}

	return true
}

func ParseInput() ([]JunctionBox, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var junctionBoxes []JunctionBox

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructionParts := strings.Split(scanner.Text(), ",")

		xIntValue, err := strconv.Atoi(instructionParts[0])
		if err != nil {
			return nil, fmt.Errorf("error converting x int value: %s", err)
		}

		yIntValue, err := strconv.Atoi(instructionParts[1])
		if err != nil {
			return nil, fmt.Errorf("error converting y int value: %s", err)
		}

		zIntValue, err := strconv.Atoi(instructionParts[2])
		if err != nil {
			return nil, fmt.Errorf("error converting z int value: %s", err)
		}

		junctionBoxes = append(junctionBoxes, JunctionBox{
			X: xIntValue,
			Y: yIntValue,
			Z: zIntValue,
		})
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error reading the file: %s", err)
	}

	return junctionBoxes, nil
}

func calculateDistances(junctionBoxes []JunctionBox) []JunctionBoxDistance {
	var distances []JunctionBoxDistance

	for i := range junctionBoxes {
		boxA := junctionBoxes[i]

		for j := i + 1; j < len(junctionBoxes); j++ {
			boxB := junctionBoxes[j]

			distances = append(distances, boxA.distance(boxB))
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})
	return distances
}

func initializeCircuits(junctionBoxes []JunctionBox) []map[JunctionBox]struct{} {
	var circuits []map[JunctionBox]struct{}

	for _, box := range junctionBoxes {
		circuits = append(circuits, map[JunctionBox]struct{}{box: {}})
	}
	return circuits
}

func initializeBoxesCircuitsAndDistances() ([]JunctionBoxDistance, []map[JunctionBox]struct{}, error) {
	junctionBoxes, err := ParseInput()
	if err != nil {
		return nil, nil, err
	}

	return calculateDistances(junctionBoxes), initializeCircuits(junctionBoxes), nil
}

func deleteByIndex(circuits []map[JunctionBox]struct{}, index int) []map[JunctionBox]struct{} {
	return append(circuits[:index], circuits[index+1:]...)
}

func mergeIfPossible(circuits []map[JunctionBox]struct{}, from JunctionBox, to JunctionBox) []map[JunctionBox]struct{} {
	var circuitA, circuitB map[JunctionBox]struct{}
	var indexB int

	for i, circuit := range circuits {
		_, ok := circuit[from]
		if ok {
			circuitA = circuit
		}

		_, ok = circuit[to]
		if ok {
			circuitB = circuit
			indexB = i
		}
	}

	if circuitA != nil && circuitB != nil && !isSameCircuit(circuitA, circuitB) {
		for box := range circuitB {
			circuitA[box] = struct{}{}
		}

		circuits = deleteByIndex(circuits, indexB)

		sort.Slice(circuits, func(i, j int) bool {
			return len(circuits[i]) > len(circuits[j])
		})
	}
	return circuits
}

func Part1() error {
	distances, circuits, err := initializeBoxesCircuitsAndDistances()
	if err != nil {
		return err
	}

	for iteration := 0; iteration < 1000; iteration++ {
		from := distances[iteration].From
		to := distances[iteration].To

		mergeIfPossible(circuits, from, to)
	}

	fmt.Println("Day 8a", len(circuits[0])*len(circuits[1])*len(circuits[2]))

	return nil
}

func Part2() error {
	distances, circuits, err := initializeBoxesCircuitsAndDistances()
	if err != nil {
		return err
	}

	for iteration := 0; true; iteration++ {
		from := distances[iteration].From
		to := distances[iteration].To

		circuits = mergeIfPossible(circuits, from, to)

		allCircuitsValid := true
		for _, circuit := range circuits {
			if len(circuit) == 1 {
				allCircuitsValid = false

				break
			}
		}

		if allCircuitsValid {
			fmt.Println("Day 8b", from.X*to.X)

			break
		}
	}

	return nil
}
