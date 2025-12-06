package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "inputs/2025/day6.txt"

type MathProblem struct {
	Numbers  []int
	Operator string
}

func Part1() error {
	var result int
	var problems []MathProblem

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

	for inputRowIndex := 0; scanner.Scan(); inputRowIndex++ {
		inputRow := scanner.Text()

		for rowPartIndex, rowPart := range strings.Fields(inputRow) {
			if inputRowIndex == 0 {
				number, err := strconv.Atoi(rowPart)
				if err != nil {
					return fmt.Errorf("error converting %s to integer: %s", rowPart, err)
				}

				problems = append(problems, MathProblem{
					Numbers: []int{number},
				})
			} else {
				number, err := strconv.Atoi(rowPart)
				if err == nil {
					problems[rowPartIndex].Numbers = append(problems[rowPartIndex].Numbers, number)
				} else {
					problems[rowPartIndex].Operator = rowPart
				}
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	for _, problem := range problems {
		switch problem.Operator {
		case "*":
			problemResult := 1
			for _, number := range problem.Numbers {
				problemResult *= number
			}

			result += problemResult

		case "+":
			problemResult := 0
			for _, number := range problem.Numbers {
				problemResult += number
			}

			result += problemResult
		}
	}

	fmt.Println("Day 6a:", result)

	return nil
}
