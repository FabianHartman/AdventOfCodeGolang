package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "inputs/2025/day10.txt"

func toBitmask(arr []int) int {
	mask := 0
	for _, v := range arr {
		mask |= 1 << v
	}

	return mask
}

func Part1() error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner error: %s", err)
	}

	var result int
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			return fmt.Errorf("invalid line format: %s", line)
		}

		lights := strings.Trim(parts[0], "[]")
		wiring := parts[1 : len(parts)-1]

		var lightPositions []int
		for i, ch := range lights {
			if ch == '#' {
				lightPositions = append(lightPositions, i)
			}
		}

		var buttonMasks []int
		for _, wire := range wiring {
			wire = strings.Trim(wire, "()")
			if wire == "" {
				continue
			}

			nums := strings.Split(wire, ",")
			var button []int
			for _, numStr := range nums {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return fmt.Errorf("error converting number: %v", err)
				}

				button = append(button, num)
			}

			buttonMasks = append(buttonMasks, toBitmask(button))
		}

		startMask := toBitmask(lightPositions)
		endMask := 0

		current := map[int]bool{startMask: true}
		iterations := 0

		for {
			_, exists := current[endMask]
			if exists {
				break
			}

			nextSet := make(map[int]bool)
			for currentMask := range current {
				for _, buttonMask := range buttonMasks {
					nextSet[currentMask^buttonMask] = true
				}
			}

			current = nextSet
			iterations++
		}

		result += iterations
	}

	fmt.Println("Day10a:", result)

	return nil
}

type buttonCombinaison struct {
	counter          []int
	nbPressedButtons int
}

func solveMinPresses(buttons [][]int, targets []int) int {
	numCounters := len(targets)
	combinaisons := allCombinaisons(buttons, numCounters)

	minPresses, _ := solveRecursive(targets, combinaisons)

	return minPresses
}

func allCombinaisons(buttons [][]int, numCounters int) []buttonCombinaison {
	numButtons := len(buttons)
	if numButtons == 0 {
		return []buttonCombinaison{{counter: make([]int, numCounters), nbPressedButtons: 0}}
	}

	res := make([]buttonCombinaison, 0, 1<<numButtons)
	for n := 0; n < (1 << numButtons); n++ {
		counter := make([]int, numCounters)
		nbPressedButtons := 0
		for j := 0; j < numButtons; j++ {
			if (n & (1 << j)) != 0 {
				nbPressedButtons++
				for _, idx := range buttons[j] {
					counter[idx]++
				}
			}
		}

		res = append(res, buttonCombinaison{counter, nbPressedButtons})
	}
	return res
}

func solveRecursive(counter []int, combinaisons []buttonCombinaison) (int, bool) {
	if isZero(counter) {
		return 0, true
	}

	res := int(^uint(0) >> 1)
	for _, comb := range combinaisons {
		if !smallerOrEqual(comb.counter, counter) {
			continue
		}
		if !equalsModulo2(comb.counter, counter) {
			continue
		}

		nextCounter := make([]int, len(counter))
		for i := 0; i < len(counter); i++ {
			nextCounter[i] = (counter[i] - comb.counter[i]) / 2
		}

		rec, ok := solveRecursive(nextCounter, combinaisons)
		if !ok {
			continue
		}

		if n := 2*rec + comb.nbPressedButtons; n < res {
			res = n
		}
	}

	if res < int(^uint(0)>>1) {
		return res, true
	}

	return 0, false
}

func isZero(counter []int) bool {
	for i := 0; i < len(counter); i++ {
		if counter[i] != 0 {
			return false
		}
	}
	return true
}

func smallerOrEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return false
		}
	}
	return true
}

func equalsModulo2(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i]%2 != b[i]%2 {
			return false
		}
	}
	return true
}

func Part2() error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("scanner error: %s", err)
	}

	var result int
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			return fmt.Errorf("invalid line format: %s", line)
		}

		var buttons [][]int
		startIdx := 1
		endIdx := len(parts) - 1

		for i := startIdx; i < endIdx; i++ {
			wire := strings.Trim(parts[i], "()")
			if wire == "" {
				continue
			}

			nums := strings.Split(wire, ",")
			var button []int
			for _, numStr := range nums {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return fmt.Errorf("error converting number: %v", err)
				}

				button = append(button, num)
			}

			buttons = append(buttons, button)
		}

		joltageStr := strings.Trim(parts[len(parts)-1], "{}")
		joltageNums := strings.Split(joltageStr, ",")
		var targets []int
		for _, numStr := range joltageNums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return fmt.Errorf("error converting joltage: %v", err)
			}

			targets = append(targets, num)
		}

		minPresses := solveMinPresses(buttons, targets)
		result += minPresses
	}

	fmt.Println("Day10b:", result)

	return nil
}
