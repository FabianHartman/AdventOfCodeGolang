package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputPath = "inputs/2025/day1.txt"

type dial struct {
	Position int
}

func Part1() error {
	var result int

	dial := &dial{
		Position: 50,
	}

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

	for scanner.Scan() {
		instruction := scanner.Text()
		movement, _ := strconv.Atoi(instruction[1:])

		if instruction[0] == 'L' {
			dial.Position -= movement
		} else {
			dial.Position += movement
		}

		dial.Position = (dial.Position%100 + 100) % 100

		if dial.Position == 0 {
			result++
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 1a:", result)

	return nil
}

func Part2() error {
	var result int

	dial := &dial{
		Position: 50,
	}

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

	for scanner.Scan() {
		line := scanner.Text()
		movement, _ := strconv.Atoi(line[1:])

		result = result + movement/100
		movement = movement % 100

		if line[0] == 'L' {
			dial.Position -= movement
			if dial.Position < 0 {
				dial.Position += 100
				if dial.Position+movement != 100 {
					result++
				}
			} else if dial.Position == 0 {
				result++
			}
		} else {
			dial.Position += movement
			if dial.Position > 99 {
				dial.Position -= 100
				if dial.Position-movement != 0 {
					result++
				}
			} else if dial.Position == 0 {
				result++
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 1b:", result)

	return nil
}
