package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPath = "inputs/2024/day7.txt"

type Equation struct {
	Result  int
	Numbers []int
}

type Operators []string

func (e *Equation) possiblyValid(part int) (bool, error) {
	var operatorCombinations []Operators

	switch part {
	case 1:
		operatorCombinations = e.generateOperatorCombinationsPart1()
	case 2:
		operatorCombinations = e.generateOperatorCombinationsPart2()
	default:
		return false, fmt.Errorf("incorrect part " + strconv.Itoa(part))
	}

	for _, operatorCombination := range operatorCombinations {
		valid, err := e.validOperators(operatorCombination)
		if err != nil {
			return false, err
		}

		if valid {
			return true, nil
		}
	}

	return false, nil
}

func (e *Equation) validOperators(ops Operators) (bool, error) {
	numbers := make([]int, len(e.Numbers))
	copy(numbers, e.Numbers)

	result := numbers[0]

	for i, op := range ops {
		switch op {
		case "*":
			result *= numbers[i+1]
		case "+":
			result += numbers[i+1]
		case "||":
			var err error
			result, err = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(numbers[i+1]))
			if err != nil {
				return false, err
			}
		default:
			return false, fmt.Errorf("invalid operator %s", op)
		}
	}

	return result == e.Result, nil
}

func (e *Equation) generateOperatorCombinationsPart1() []Operators {
	totalCombinations := 1 << (len(e.Numbers) - 1)
	combinations := make([]Operators, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		opCombo := Operators{}
		for j := 0; j < len(e.Numbers)-1; j++ {
			if (i & (1 << j)) != 0 {
				opCombo = append(opCombo, "+")
			} else {
				opCombo = append(opCombo, "*")
			}
		}

		combinations[i] = opCombo
	}

	return combinations
}

func (e *Equation) generateOperatorCombinationsPart2() []Operators {
	totalCombinations := 1
	for i := 0; i < len(e.Numbers)-1; i++ {
		totalCombinations *= 3
	}

	combinations := make([]Operators, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		opCombo := Operators{}
		for j := 0; j < len(e.Numbers)-1; j++ {
			powerOfThree := 1
			for k := 0; k < j; k++ {
				powerOfThree *= 3
			}

			switch (i / powerOfThree) % 3 {
			case 0:
				opCombo = append(opCombo, "*")
			case 1:
				opCombo = append(opCombo, "+")
			case 2:
				opCombo = append(opCombo, "||")
			}
		}

		combinations[i] = opCombo
	}

	return combinations
}

func input() ([]Equation, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	equations := []Equation{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowParts := strings.Split(scanner.Text(), ": ")
		result, err := strconv.Atoi(rowParts[0])
		if err != nil {
			return nil, fmt.Errorf("error converting row part: %s", err)
		}

		equation := Equation{Result: result}

		stringNumbers := strings.Split(rowParts[1], " ")
		for _, stringNumber := range stringNumbers {
			number, err := strconv.Atoi(stringNumber)
			if err != nil {
				return nil, fmt.Errorf("error converting number: %s", err)
			}

			equation.Numbers = append(equation.Numbers, number)
		}

		equations = append(equations, equation)
	}

	return equations, nil
}

func Part1() error {
	equations, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, equation := range equations {
		valid, err := equation.possiblyValid(1)
		if err != nil {
			return err
		}

		if valid {
			total += equation.Result
		}
	}

	fmt.Println("Day 7a:", total)

	return nil
}

func Part2() error {
	equations, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, equation := range equations {
		valid, err := equation.possiblyValid(2)
		if err != nil {
			return err
		}

		if valid {
			total += equation.Result
		}
	}

	fmt.Println("Day 7b:", total)

	return nil
}
