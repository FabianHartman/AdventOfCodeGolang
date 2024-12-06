package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	Row, Col int
}

type Map struct {
	Obstructions   []Position
	MaxRow, MaxCol int
}

func (m *Map) isObstruction(position Position) bool {
	for _, o := range m.Obstructions {
		if o.Row == position.Row && o.Col == position.Col {
			return true
		}
	}

	return false
}

func (m *Map) isOnMap(position Position) bool {
	return !(position.Row < 0 || position.Col < 0 || position.Row > m.MaxRow || position.Col > m.MaxCol)
}

func (m *Map) findWalkOver(position *Position, direction string) ([]Position, bool) {
	visitedLocations := []Position{}
	walkingPosition := Position{position.Row, position.Col}
	for true {
		switch direction {
		case "up":
			walkingPosition.Row--
		case "right":
			walkingPosition.Col++
		case "down":
			walkingPosition.Row++
		case "left":
			walkingPosition.Col--
		}

		if !m.isOnMap(walkingPosition) {
			return visitedLocations, true
		}

		if m.isObstruction(walkingPosition) {
			break
		} else {
			visitedLocations = append(visitedLocations, walkingPosition)
		}
	}

	return visitedLocations, false
}

var inputPath6 string = "inputs/day6.txt"

func inputDay6() (*Map, *Position, error) {
	file, err := os.Open(inputPath6)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	obstructionsMap := Map{Obstructions: []Position{}}
	var startingLocation Position

	rowI := -1

	for scanner.Scan() {
		rowI++
		for colI, value := range strings.Split(scanner.Text(), "") {
			if value == "#" {
				obstructionsMap.Obstructions = append(obstructionsMap.Obstructions, Position{Row: rowI, Col: colI})
			} else if value == "^" {
				startingLocation = Position{Row: rowI, Col: colI}
			}

			obstructionsMap.MaxCol = colI

		}

		obstructionsMap.MaxRow = rowI

	}

	return &obstructionsMap, &startingLocation, nil
}

func getNextDirection(direction string) (string, error) {
	switch direction {
	case "up":
		return "right", nil
	case "right":
		return "down", nil
	case "down":
		return "left", nil
	case "left":
		return "up", nil
	default:
		return "", fmt.Errorf("incorrect direction: %s", direction)
	}
}

func Day6a() error {
	obstructionsMap, guardPosition, err := inputDay6()
	if err != nil {
		return err
	}

	direction := "up"
	uniquePositions := map[Position]bool{}
	uniquePositions[Position{Row: guardPosition.Row, Col: guardPosition.Col}] = true

	for true {
		newPositions, finished := obstructionsMap.findWalkOver(guardPosition, direction)

		for _, position := range newPositions {
			uniquePositions[Position{Row: position.Row, Col: position.Col}] = true
		}

		guardPosition = &newPositions[len(newPositions)-1]

		if finished {
			break
		}

		direction, err = getNextDirection(direction)
		if err != nil {
			return err
		}
	}

	fmt.Println("Day 6a:", len(uniquePositions))

	return nil
}
