package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "inputs/2025/day6.txt"

type MathProblemPart1 struct {
	Numbers  []int
	Operator string
}

func Part1() error {
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

	var problems []MathProblemPart1

	for inputRowIndex := 0; scanner.Scan(); inputRowIndex++ {
		inputRow := scanner.Text()

		for rowPartIndex, rowPart := range strings.Fields(inputRow) {
			if inputRowIndex == 0 {
				number, err := strconv.Atoi(rowPart)
				if err != nil {
					return fmt.Errorf("error converting %s to integer: %s", rowPart, err)
				}

				problems = append(problems, MathProblemPart1{
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

	result := 0

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

type MathProblemPart2 struct {
	Numbers  []string
	Operator string
}

func Part2() error {
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

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	width := 0
	length := len(input)

	for _, line := range input {
		if len(line) > width {
			width = len(line)
		}
	}

	allCellsEmpty := true
	problems := []MathProblemPart2{{
		Numbers: make([]string, length-1),
	}}
	problemsIndex := 0

	for colI := range width {
		allCellsEmpty = true

		for rowI := range input {
			if len(input[rowI]) <= colI {
				continue
			}

			if string(input[rowI][colI]) != " " {
				allCellsEmpty = false

				break
			}
		}

		if allCellsEmpty {
			problems = append(problems, MathProblemPart2{
				Numbers: make([]string, length-1),
			})
			problemsIndex++

			continue
		}

		for rowI := range input {
			if len(input[rowI]) <= colI {
				continue
			}

			char := string(input[rowI][colI])

			if char != " " {
				if rowI == length-1 {
					problems[problemsIndex].Operator = char
				}
			}

			if rowI < length-1 {
				problems[problemsIndex].Numbers[rowI] += char
			}
		}
	}

	result := 0

	for _, problem := range problems {
		numbersWidth := 0
		for _, number := range problem.Numbers {
			if len(number) > numbersWidth {
				numbersWidth = len(number)
			}
		}

		switch problem.Operator {
		case "*":
			problemResult := 1

			for i := 0; i < numbersWidth; i++ {
				verticalStringNumber := ""
				for _, number := range problem.Numbers {
					if i < len(number) {
						char := string(number[i])
						if char != " " {
							verticalStringNumber += char
						}
					}
				}

				verticalNumber, err := strconv.Atoi(verticalStringNumber)
				if err != nil {
					return fmt.Errorf("error converting %s to integer: %s", verticalStringNumber, err)
				}

				problemResult *= verticalNumber
			}

			result += problemResult

		case "+":
			problemResult := 0

			for i := 0; i < numbersWidth; i++ {
				verticalStringNumber := ""
				for _, number := range problem.Numbers {
					if i < len(number) {
						char := string(number[i])
						if char != " " {
							verticalStringNumber += char
						}
					}
				}

				verticalNumber, err := strconv.Atoi(verticalStringNumber)
				if err != nil {
					return fmt.Errorf("error converting %s to integer: %s", verticalStringNumber, err)
				}

				problemResult += verticalNumber
			}

			result += problemResult
		}
	}

	fmt.Println("Day 6b:", result)

	return nil
}
