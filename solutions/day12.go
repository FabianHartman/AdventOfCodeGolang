package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputPathDay12 string = "inputs/day12.txt"

func inputDay12() ([][]string, error) {
	file, err := os.Open(inputPathDay12)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	gardenMap := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gardenMap = append(gardenMap, strings.Split(scanner.Text(), ""))
	}

	return gardenMap, nil
}

var directions = [4][2]int{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

type PlantCoord struct {
	X, Y int
}

func calculateRegion(gardenMap [][]string, visited [][]bool, coord *PlantCoord, plantType string) (int, int) {
	stack := []PlantCoord{*coord}
	visited[coord.X][coord.Y] = true
	area := 0
	perimeter := 0

	for len(stack) > 0 {
		currentCoord := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area++

		for _, direction := range directions {
			adjacentCoord := PlantCoord{X: currentCoord.X + direction[0], Y: currentCoord.Y + direction[1]}

			if adjacentCoord.X < 0 || adjacentCoord.X >= len(gardenMap) || adjacentCoord.Y < 0 || adjacentCoord.Y >= len(gardenMap[0]) || gardenMap[adjacentCoord.X][adjacentCoord.Y] != plantType {
				perimeter++
			} else if !visited[adjacentCoord.X][adjacentCoord.Y] {
				visited[adjacentCoord.X][adjacentCoord.Y] = true
				stack = append(stack, adjacentCoord)
			}
		}
	}

	return area, perimeter
}

func solveDay12Part1(input [][]string) (int, error) {
	visited := make([][]bool, len(input))
	for i := range visited {
		visited[i] = make([]bool, len(input[i]))
	}

	totalPrice := 0

	for x := 0; x < len(input); x++ {
		for y := 0; y < len(input[x]); y++ {
			if !visited[x][y] {
				plantType := input[x][y]
				area, perimeter := calculateRegion(input, visited, &PlantCoord{X: x, Y: y}, plantType)
				price := area * perimeter
				totalPrice += price
			}
		}
	}

	return totalPrice, nil
}

func Day12a() error {
	input, err := inputDay12()
	if err != nil {
		return err
	}

	total, err := solveDay12Part1(input)
	if err != nil {
		return err
	}

	fmt.Println("Day12a:", total)

	return nil
}
