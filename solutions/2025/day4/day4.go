package day4

import (
	"bufio"
	"fmt"
	"os"
)

const inputPath = "inputs/2025/day4.txt"

type Position struct {
	Row int
	Col int
}

func (this Position) GetAdjacentPositions() []Position {
	positions := []Position{}

	positions = append(positions, Position{Row: this.Row - 1, Col: this.Col - 1})
	positions = append(positions, Position{Row: this.Row - 1, Col: this.Col})
	positions = append(positions, Position{Row: this.Row - 1, Col: this.Col + 1})
	positions = append(positions, Position{Row: this.Row, Col: this.Col - 1})
	positions = append(positions, Position{Row: this.Row, Col: this.Col + 1})
	positions = append(positions, Position{Row: this.Row + 1, Col: this.Col - 1})
	positions = append(positions, Position{Row: this.Row + 1, Col: this.Col})
	positions = append(positions, Position{Row: this.Row + 1, Col: this.Col + 1})

	return positions
}

func Part1() error {
	var result int
	paperPositions := map[Position]bool{}

	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for row := 0; scanner.Scan(); row++ {
		inputRow := scanner.Text()

		for col, char := range inputRow {
			if char == '@' {
				paperPositions[Position{Row: row, Col: col}] = true
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	for paper := range paperPositions {
		adjacentsAmount := 0
		for _, adjacentPosition := range paper.GetAdjacentPositions() {
			if paperPositions[adjacentPosition] {
				adjacentsAmount++

				if adjacentsAmount >= 4 {
					break
				}
			}
		}

		if adjacentsAmount < 4 {
			result++
		}
	}

	fmt.Println("Day 4a:", result)

	return nil
}
