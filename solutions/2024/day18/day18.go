package day18

import (
	"bufio"
	"fmt"
	"os"
)

var inputPath = "inputs/2024/day18.txt"

type Position struct {
	X, Y int
}

func input(amountOfRain *int) (map[Position]bool, []Position, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	positions := map[Position]bool{}
	positionsList := []Position{}

	if amountOfRain == nil {
		highAmountOfRain := 9999999999999999
		amountOfRain = &highAmountOfRain
	}

	for n := 1; scanner.Scan(); n++ {
		if n > *amountOfRain {
			break
		}

		position := Position{}
		line := scanner.Text()

		_, err = fmt.Sscanf(line, "%d,%d", &position.X, &position.Y)
		if err != nil {
			return nil, nil, err
		}

		positions[position] = true
		positionsList = append(positionsList, position)
	}

	return positions, positionsList, nil
}

func getAdjacentPositions(position *Position, maxX, maxY int) []Position {
	possibleAdjacentPositions := []Position{}

	possibleAdjacentPositions = append(possibleAdjacentPositions, Position{X: position.X + 1, Y: position.Y})
	possibleAdjacentPositions = append(possibleAdjacentPositions, Position{X: position.X - 1, Y: position.Y})
	possibleAdjacentPositions = append(possibleAdjacentPositions, Position{X: position.X, Y: position.Y + 1})
	possibleAdjacentPositions = append(possibleAdjacentPositions, Position{X: position.X, Y: position.Y - 1})

	adjacentPositions := []Position{}

	for _, adjacentPosition := range possibleAdjacentPositions {
		if !(adjacentPosition.X > maxX || adjacentPosition.Y > maxY || adjacentPosition.X < 0 || adjacentPosition.Y < 0) {
			adjacentPositions = append(adjacentPositions, adjacentPosition)
		}
	}

	return adjacentPositions
}

func findRoute(from *Position, to *Position, corruption map[Position]bool, maxX, maxY int) (int, error) {
	newHeads := map[Position]bool{}

	currentHeads := []Position{}
	currentHeads = append(currentHeads, *from)
	visited := map[Position]bool{}
	for steps := 1; ; steps++ {
		if len(currentHeads) == 0 {
			return 0, fmt.Errorf("no possible route")
		}

		for _, head := range currentHeads {
			visited[head] = true
			for _, adjacentPosition := range getAdjacentPositions(&head, maxX, maxY) {
				if adjacentPosition == *to {
					return steps, nil
				}

				if corruption[adjacentPosition] {
					continue
				}

				if !visited[adjacentPosition] {
					newHeads[adjacentPosition] = true
				}
			}
		}

		currentHeads = []Position{}

		for head := range newHeads {
			currentHeads = append(currentHeads, head)
		}

		newHeads = map[Position]bool{}
	}
}

func findPreventingByte(from *Position, to *Position, corruptionsList []Position, maxX, maxY int) *Position {
	for i := 1; i <= len(corruptionsList); i++ {
		corruptions := make(map[Position]bool, i)

		for _, corruption := range corruptionsList[:i] {
			corruptions[corruption] = true
		}

		_, err := findRoute(from, to, corruptions, maxX, maxY)
		if err != nil {
			return &corruptionsList[i-1]
		}
	}

	return nil
}

func Part1() error {
	kb := 1024
	corruption, _, err := input(&kb)
	if err != nil {
		return err
	}

	routeSize, err := findRoute(&Position{0, 0}, &Position{70, 70}, corruption, 70, 70)
	if err != nil {
		return err
	}

	fmt.Println("Day 18a:", routeSize)

	return nil
}

func Part2() error {
	_, corruptionsList, err := input(nil)
	if err != nil {
		return err
	}

	preventingByte := findPreventingByte(&Position{0, 0}, &Position{70, 70}, corruptionsList, 70, 70)

	fmt.Println(fmt.Sprintf("Day 18b: %d,%d", preventingByte.X, preventingByte.Y))

	return nil
}
