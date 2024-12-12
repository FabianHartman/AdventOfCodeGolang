package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var inputPath string = "inputs/2024/day4.txt"

func input() ([][]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rows := [][]string{}

	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), ""))
	}

	return rows, nil
}

type Coord struct {
	Col int
	Row int
}

func (c *Coord) getLeftDiagonalCoords() []Coord {
	coords := []Coord{}

	coords = append(coords, Coord{c.Col - 1, c.Row - 1})
	coords = append(coords, Coord{c.Col + 1, c.Row + 1})

	return coords
}

func (c *Coord) getRightDiagonalCoords() []Coord {
	coords := []Coord{}

	coords = append(coords, Coord{c.Col + 1, c.Row - 1})
	coords = append(coords, Coord{c.Col - 1, c.Row + 1})

	return coords
}

func findACoords(input [][]string) []Coord {
	ACoords := []Coord{}

	for row := 1; row < len(input)-1; row++ {
		for col := 1; col < len(input[row])-1; col++ {
			if input[row][col] == "A" {
				ACoords = append(ACoords, Coord{col, row})
			}
		}
	}

	return ACoords
}

func reversedString(input []string) string {
	output := ""
	for _, s := range input {
		output = s + output
	}

	return output
}

func generateStrings(rows [][]string) []string {
	n := len(rows)
	var result []string
	for _, row := range rows {
		result = append(result, strings.Join(row, ""))
		result = append(result, reversedString(row))
	}

	for i := range rows[0] {
		col := ""
		for _, row := range rows {
			col += row[i]
		}

		result = append(result, col)
		result = append(result, reversedString(strings.Split(col, "")))
	}

	for colStart := 0; colStart < n; colStart++ {
		diagonal := ""
		for i := 0; colStart+i < n && i < n; i++ {
			diagonal = diagonal + rows[i][colStart+i]
		}

		result = append(result, diagonal)
		result = append(result, reversedString(strings.Split(diagonal, "")))
	}

	for rowStart := 1; rowStart < n; rowStart++ {
		diagonal := ""
		for i := 0; rowStart+i < n && i < n; i++ {
			diagonal = diagonal + rows[rowStart+i][i]
		}

		result = append(result, diagonal)
		result = append(result, reversedString(strings.Split(diagonal, "")))
	}

	for colStart := n - 1; colStart >= 0; colStart-- {
		diagonal := ""
		for i := 0; colStart-i >= 0 && i < n; i++ {
			diagonal = diagonal + rows[i][colStart-i]
		}

		result = append(result, diagonal)
		result = append(result, reversedString(strings.Split(diagonal, "")))
	}

	for rowStart := 1; rowStart < n; rowStart++ {
		diagonal := ""
		for i := 0; rowStart+i < n && n-i-1 >= 0; i++ {
			diagonal = diagonal + rows[rowStart+i][n-i-1]
		}

		result = append(result, diagonal)
		result = append(result, reversedString(strings.Split(diagonal, "")))
	}

	return result
}

func Part1() error {
	input, err := input()
	if err != nil {
		return err
	}

	options := generateStrings(input)
	xmasPattern := regexp.MustCompile(`XMAS`)
	total := 0

	for _, option := range options {
		total += len(xmasPattern.FindAllString(option, -1))
	}

	fmt.Println("day 4a:", total)

	return nil
}

func Part2() error {
	input, err := input()
	if err != nil {
		return err
	}

	total := 0

	ACoords := findACoords(input)

	for _, aCoord := range ACoords {
		leftCorrect, rightCorrect := false, false

		leftString := ""
		for _, coord := range aCoord.getLeftDiagonalCoords() {
			leftString += input[coord.Row][coord.Col]
		}
		if strings.Contains(leftString, "M") {
			if strings.Contains(leftString, "S") {
				leftCorrect = true
			}
		}

		rightString := ""
		for _, coord := range aCoord.getRightDiagonalCoords() {
			rightString += input[coord.Row][coord.Col]
		}
		if strings.Contains(rightString, "M") {
			if strings.Contains(rightString, "S") {
				rightCorrect = true
			}
		}

		if leftCorrect && rightCorrect {
			total++
		}
	}

	fmt.Println("day 4b:", total)

	return nil
}
