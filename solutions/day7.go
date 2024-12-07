package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPath7 string = "inputs/day7.txt"

type Equation struct {
	Result  int
	Numbers []int
}

type Operators []string

func (e *Equation) possiblyValid() bool {
	operatorCombinations := e.generateOperatorCombinations()
	for _, operatorCombination := range operatorCombinations {
		if e.validOperators(operatorCombination) {
			return true
		}
	}

	return false
}

func (e *Equation) validOperators(ops Operators) bool {
	result := e.Numbers[0]
	for i, op := range ops {
		if op == "*" {
			result *= e.Numbers[i+1]
		} else if op == "+" {
			result += e.Numbers[i+1]
		}
	}

	return result == e.Result
}

func (e *Equation) generateOperatorCombinations() []Operators {
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

func inputDay7() ([]Equation, error) {
	file, err := os.Open(inputPath7)
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

func Day7a() error {
	equations, err := inputDay7()
	if err != nil {
		return err
	}

	total := 0

	for _, equation := range equations {
		if equation.possiblyValid() {
			total += equation.Result
		}
	}

	fmt.Println("Day 7a:", total)

	return nil
}
