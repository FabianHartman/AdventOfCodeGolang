package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "inputs/2025/day2.txt"

func Part1() error {
	var result int

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
	scanner.Scan()

	for _, instructionPart := range strings.Split(scanner.Text(), ",") {
		dashIndex := strings.Index(instructionPart, "-")

		begin, _ := strconv.Atoi(instructionPart[:dashIndex])
		end, _ := strconv.Atoi(instructionPart[dashIndex+1:])

		for id := begin; id <= end; id++ {
			stringValue := strconv.Itoa(id)

			midpoint := len(stringValue) / 2

			part1 := stringValue[:midpoint]
			part2 := stringValue[midpoint:]

			if part1 == part2 {
				result += id
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 2a:", result)

	return nil
}

func sequenceIsRepeating(stringValue string, sequenceLength int, sequence string) bool {
	nOfSequenceInstances := len(stringValue) / sequenceLength

	for i := 1; i < nOfSequenceInstances; i++ {
		if stringValue[sequenceLength*i:sequenceLength*i+sequenceLength] != sequence {
			return false
		}
	}

	return true
}

func hasRepeatingSequence(id int) bool {
	stringValue := strconv.Itoa(id)

	for sequenceLength := 1; sequenceLength <= len(stringValue)/2; sequenceLength++ {
		if len(stringValue)%sequenceLength != 0 {
			continue
		}

		if sequenceIsRepeating(stringValue, sequenceLength, stringValue[:sequenceLength]) {
			return true
		}
	}

	return false
}

func Part2() error {
	var result int

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
	scanner.Scan()

	for _, instructionPart := range strings.Split(scanner.Text(), ",") {
		dashIndex := strings.Index(instructionPart, "-")

		begin, _ := strconv.Atoi(instructionPart[:dashIndex])
		end, _ := strconv.Atoi(instructionPart[dashIndex+1:])

		for id := begin; id <= end; id++ {
			if hasRepeatingSequence(id) {
				result += id
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 2b:", result)

	return nil
}
