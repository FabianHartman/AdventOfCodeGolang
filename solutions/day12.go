package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var inputFilePath string = "inputs/day12.txt"

func readInput() ([][]rune, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	var garden [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		garden = append(garden, row)
	}
	return garden, nil
}

type Coordinate struct {
	X, Y int
}

var movementDirections = []Coordinate{
	{-1, 0}, // Up
	{0, -1}, // Left
	{1, 0},  // Down
	{0, 1},  // Right
}

var cornerOffsets = map[Coordinate][]Coordinate{
	{1, 1}:   {{1, 0}, {0, 1}},
	{-1, 1}:  {{-1, 0}, {0, 1}},
	{1, -1}:  {{1, 0}, {0, -1}},
	{-1, -1}: {{-1, 0}, {0, -1}},
}

func exploreGardenWithPerimeter(currentPosition Coordinate, garden [][]rune, visited map[Coordinate]bool, currentPlant rune, areaSize, perimeterSize *int) {
	maxRow, maxCol := len(garden), len(garden[0])
	neighborCount := 0

	visited[currentPosition] = true

	for _, direction := range movementDirections {
		nextPosition := Coordinate{X: currentPosition.X + direction.X, Y: currentPosition.Y + direction.Y}
		if isInsideMap(nextPosition, maxRow, maxCol) && garden[nextPosition.X][nextPosition.Y] == currentPlant {
			neighborCount++

			if !visited[nextPosition] {
				*areaSize++

				exploreGardenWithPerimeter(nextPosition, garden, visited, currentPlant, areaSize, perimeterSize)
			}
		}
	}

	*perimeterSize += 4 - neighborCount
}

func calculatePartTwo(garden [][]rune) (totalCost int) {
	visited := make(map[Coordinate]bool)

	for row := range garden {
		for col := range garden[row] {
			if !visited[Coordinate{X: row, Y: col}] {
				var areaSize, cornerCount int

				exploreGardenWithCorners(Coordinate{X: row, Y: col}, garden, visited, garden[row][col], &areaSize, &cornerCount)

				totalCost += (areaSize + 1) * cornerCount
			}
		}
	}
	return
}

func exploreGardenWithCorners(currentPosition Coordinate, garden [][]rune, visited map[Coordinate]bool, currentPlant rune, areaSize, cornerCount *int) {
	maxRow, maxCol := len(garden), len(garden[0])
	visited[currentPosition] = true

	for _, direction := range movementDirections {
		nextPosition := Coordinate{X: currentPosition.X + direction.X, Y: currentPosition.Y + direction.Y}
		if isInsideMap(nextPosition, maxRow, maxCol) && garden[nextPosition.X][nextPosition.Y] == currentPlant {
			if !visited[nextPosition] {
				*areaSize++

				exploreGardenWithCorners(nextPosition, garden, visited, currentPlant, areaSize, cornerCount)
			}
		}
	}

	for corner, pair := range cornerOffsets {
		cornerPosition := Coordinate{X: currentPosition.X + corner.X, Y: currentPosition.Y + corner.Y}
		adjacent1 := Coordinate{X: currentPosition.X + pair[0].X, Y: currentPosition.Y + pair[0].Y}
		adjacent2 := Coordinate{X: currentPosition.X + pair[1].X, Y: currentPosition.Y + pair[1].Y}

		if !isMatching(adjacent1, currentPosition, garden) && !isMatching(adjacent2, currentPosition, garden) {
			*cornerCount++
		}

		if isMatching(adjacent1, currentPosition, garden) && isMatching(adjacent2, currentPosition, garden) && !isMatching(currentPosition, cornerPosition, garden) {
			*cornerCount++
		}
	}
}

func isMatching(position1, position2 Coordinate, garden [][]rune) bool {
	maxRow, maxCol := len(garden), len(garden[0])

	position1Valid, position2Valid := isInsideMap(position1, maxRow, maxCol), isInsideMap(position2, maxRow, maxCol)

	if !position1Valid && !position2Valid {
		return true
	} else if position1Valid && position2Valid {
		plants := []rune{garden[position1.X][position1.Y], garden[position2.X][position2.Y]}
		return plants[0] == plants[1]
	} else {
		return false
	}
}

func isInsideMap(c Coordinate, maxRow, maxCol int) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < maxRow && c.Y < maxCol
}

func calculatePartOne(garden [][]rune) (totalCost int) {
	visited := map[Coordinate]bool{}

	for row := range garden {
		for col := range garden[row] {
			if !visited[Coordinate{X: row, Y: col}] {
				var areaSize, perimeterSize int

				exploreGardenWithPerimeter(Coordinate{X: row, Y: col}, garden, visited, garden[row][col], &areaSize, &perimeterSize)

				totalCost += (areaSize + 1) * perimeterSize
			}
		}
	}
	return
}

func Day12a() error {
	garden, err := readInput()
	if err != nil {
		return err
	}

	fmt.Println("Day12a:", calculatePartOne(garden))

	return nil
}

func Day12b() {
	garden, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day12b: ", calculatePartTwo(garden))
}
