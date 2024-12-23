package day13

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputPath = "inputs/2024/day13.txt"

type Pair struct {
	One, Two int
}
type Equation struct {
	a, b, c *Pair
}

func extractCoordinates(input string) (*Pair, error) {
	re := regexp.MustCompile(`X[+=](\d+), Y[+=](\d+)`)

	matches := re.FindStringSubmatch(input)

	if len(matches) < 3 {
		return nil, fmt.Errorf("no match found")
	}

	x, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	return &Pair{One: x, Two: y}, nil
}

func input() ([]Equation, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	equations := []Equation{}
	scanner := bufio.NewScanner(file)

	equation := Equation{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "A") {
			equation.a, err = extractCoordinates(line)
		} else if strings.Contains(line, "B") {
			equation.b, err = extractCoordinates(line)
		} else if strings.Contains(line, "=") {
			equation.c, err = extractCoordinates(line)

			equations = append(equations, equation)

			equation = Equation{}
		}

		if err != nil {
			return nil, err
		}
	}

	return equations, nil
}

func calcMachine(eq *Equation) int {
	ap := (eq.b.Two*eq.c.One - eq.b.One*eq.c.Two) / (eq.a.One*eq.b.Two - eq.a.Two*eq.b.One)
	bp := (eq.a.Two*eq.c.One - eq.a.One*eq.c.Two) / (eq.a.Two*eq.b.One - eq.a.One*eq.b.Two)

	if (eq.a.One*ap+eq.b.One*bp == eq.c.One) || (eq.a.Two*bp == eq.c.Two) {
		if ap < 0 || bp < 0 {
			return 0
		}

		return ap*3 + bp
	}

	return 0
}

func Part1() error {
	equations, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, equation := range equations {
		total += calcMachine(&equation)
		fmt.Println(total)
	}

	fmt.Println("Day 13a:", total)

	return nil
}

func Part2() error {
	equations, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, equation := range equations {
		equation.c.One += 10000000000000
		equation.c.Two += 10000000000000

		total += calcMachine(&equation)
	}

	fmt.Println("Day 13b:", total)

	return nil
}
