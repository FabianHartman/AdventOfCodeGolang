package day21

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var inputPath string = "inputs/2024/day21.txt"

type Position struct {
	X, Y int
}

func input() ([]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %s", err)
	}

	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output, nil
}

func makeMaps() (map[string]Position, map[string]Position) {
	numbers := make(map[string]Position)
	directions := make(map[string]Position)

	numbers["A"] = Position{2, 0}
	numbers["0"] = Position{1, 0}
	numbers["1"] = Position{0, 1}
	numbers["2"] = Position{1, 1}
	numbers["3"] = Position{2, 1}
	numbers["4"] = Position{0, 2}
	numbers["5"] = Position{1, 2}
	numbers["6"] = Position{2, 2}
	numbers["7"] = Position{0, 3}
	numbers["8"] = Position{1, 3}
	numbers["9"] = Position{2, 3}

	directions["A"] = Position{2, 1}
	directions["^"] = Position{1, 1}
	directions["<"] = Position{0, 0}
	directions["v"] = Position{1, 0}
	directions[">"] = Position{2, 0}

	return numbers, directions
}

func getNumericValue(line string) (int, error) {
	numString := ""

	for _, char := range line {
		if unicode.IsDigit(char) {
			numString += string(char)
		}
	}

	if len(numString) != 0 {
		localNum, err := strconv.ParseInt(numString, 10, 64)
		if err != nil {
			return 0, err
		}

		return int(localNum), nil
	}

	return 0, nil
}

func splitStepsByA(input []string) [][]string {
	output := [][]string{}
	current := []string{}

	for _, char := range input {
		current = append(current, char)

		if char == "A" {
			output = append(output, current)
			current = []string{}
		}
	}

	return output
}

func calculateNumericPresses(input []string, start string, numbers map[string]Position) []string {
	currentButton := numbers[start]
	buttonsToPress := []string{}

	for _, char := range input {
		horizontalInputs, verticalInputs := []string{}, []string{}
		destination := numbers[char]
		deltaX, deltaY := destination.X-currentButton.X, destination.Y-currentButton.Y

		for i := 0; float64(i) < math.Abs(float64(deltaX)); i++ {
			if deltaX >= 0 {
				horizontalInputs = append(horizontalInputs, ">")
			} else {
				horizontalInputs = append(horizontalInputs, "<")
			}
		}

		for i := 0; float64(i) < math.Abs(float64(deltaY)); i++ {
			if deltaY >= 0 {
				verticalInputs = append(verticalInputs, "^")
			} else {
				verticalInputs = append(verticalInputs, "v")
			}
		}

		if currentButton.Y == 0 && destination.X == 0 {
			buttonsToPress = append(buttonsToPress, verticalInputs...)
			buttonsToPress = append(buttonsToPress, horizontalInputs...)
		} else if currentButton.X == 0 && destination.Y == 0 {
			buttonsToPress = append(buttonsToPress, horizontalInputs...)
			buttonsToPress = append(buttonsToPress, verticalInputs...)
		} else if deltaX < 0 {
			buttonsToPress = append(buttonsToPress, horizontalInputs...)
			buttonsToPress = append(buttonsToPress, verticalInputs...)
		} else {
			buttonsToPress = append(buttonsToPress, verticalInputs...)
			buttonsToPress = append(buttonsToPress, horizontalInputs...)
		}

		currentButton = destination
		buttonsToPress = append(buttonsToPress, "A")
	}

	return buttonsToPress
}

func calculateDirectionalPresses(input []string, start string, directions map[string]Position) []string {
	currentButton := directions[start]
	actionsToPerform := []string{}

	for _, char := range input {
		horizontalInputs, verticalInputs := []string{}, []string{}
		destination := directions[char]
		deltaX, deltaY := destination.X-currentButton.X, destination.Y-currentButton.Y

		for i := 0; float64(i) < math.Abs(float64(deltaX)); i++ {
			if deltaX >= 0 {
				horizontalInputs = append(horizontalInputs, ">")
			} else {
				horizontalInputs = append(horizontalInputs, "<")
			}
		}

		for i := 0; float64(i) < math.Abs(float64(deltaY)); i++ {
			if deltaY >= 0 {
				verticalInputs = append(verticalInputs, "^")
			} else {
				verticalInputs = append(verticalInputs, "v")
			}
		}

		if currentButton.X == 0 && destination.Y == 1 {
			actionsToPerform = append(actionsToPerform, horizontalInputs...)
			actionsToPerform = append(actionsToPerform, verticalInputs...)
		} else if currentButton.Y == 1 && destination.X == 0 {
			actionsToPerform = append(actionsToPerform, verticalInputs...)
			actionsToPerform = append(actionsToPerform, horizontalInputs...)
		} else if deltaX < 0 {
			actionsToPerform = append(actionsToPerform, horizontalInputs...)
			actionsToPerform = append(actionsToPerform, verticalInputs...)
		} else {
			actionsToPerform = append(actionsToPerform, verticalInputs...)
			actionsToPerform = append(actionsToPerform, horizontalInputs...)
		}

		currentButton = destination
		actionsToPerform = append(actionsToPerform, "A")
	}

	return actionsToPerform
}

func complexity(input []string, numbers, directions map[string]Position, robots int) (int, error) {
	count := 0
	cache := make(map[string][]int)

	for _, line := range input {
		row := strings.Split(line, "")

		shortestLen := getShortest(calculateNumericPresses(row, "A", numbers), robots, 1, cache, directions)
		numericValue, err := getNumericValue(line)
		if err != nil {
			return 0, err
		}

		count += numericValue * shortestLen
	}

	return count, nil
}

func getShortest(input []string, maxRobots int, robot int, cache map[string][]int, directionalMap map[string]Position) int {
	if cachedValue, exists := cache[strings.Join(input, "")]; exists {
		if cachedValue[robot-1] != 0 {
			return cachedValue[robot-1]
		}
	} else {
		cache[strings.Join(input, "")] = make([]int, maxRobots)
	}

	directionalPresses := calculateDirectionalPresses(input, "A", directionalMap)
	cache[strings.Join(input, "")][0] = len(directionalPresses)

	if robot == maxRobots {
		return len(directionalPresses)
	}

	individualSteps := splitStepsByA(directionalPresses)

	count := 0

	for _, step := range individualSteps {
		shortest := getShortest(step, maxRobots, robot+1, cache, directionalMap)

		if _, exists := cache[strings.Join(step, "")]; !exists {
			cache[strings.Join(step, "")] = make([]int, maxRobots)
		}

		cache[strings.Join(step, "")][0] = shortest
		count += shortest
	}

	cache[strings.Join(input, "")][robot-1] = count

	return count
}

func Part1() error {
	codes, err := input()
	if err != nil {
		return err
	}

	numericalMap, directionalMap := makeMaps()

	complexity, err := complexity(codes, numericalMap, directionalMap, 2)
	if err != nil {
		return err
	}

	fmt.Println("Day21a:", complexity)

	return nil
}

func Part2() error {
	codes, err := input()
	if err != nil {
		return err
	}

	numericalMap, directionalMap := makeMaps()

	complexity, err := complexity(codes, numericalMap, directionalMap, 25)
	if err != nil {
		return err
	}

	fmt.Println("Day21a:", complexity)

	return nil
}
