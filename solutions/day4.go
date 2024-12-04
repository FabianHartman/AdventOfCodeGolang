package solutions

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var inputPath4 string = "inputs/day4.txt"

func inputDay4() ([][]string, error) {
	file, err := os.Open(inputPath4)
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

func Day4a() error {
	input, err := inputDay4()
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
