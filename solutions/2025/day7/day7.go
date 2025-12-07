package day7

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const inputPath = "inputs/2025/day7.txt"

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

	heads := make(map[int]bool)

	heads[regexp.MustCompile("S").FindStringIndex(scanner.Text())[1]] = true

	for scanner.Scan() {
		splitters := regexp.MustCompile("\\^").FindAllStringIndex(scanner.Text(), -1)

		for _, splitter := range splitters {
			if heads[splitter[1]] {
				result++

				heads[splitter[1]-1] = true
				heads[splitter[1]+1] = true

				delete(heads, splitter[1])
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 7a:", result)

	return nil
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

	heads := make(map[int]int)

	heads[regexp.MustCompile("S").FindStringIndex(scanner.Text())[1]] = 1

	for scanner.Scan() {
		splitters := regexp.MustCompile("\\^").FindAllStringIndex(scanner.Text(), -1)

		for _, splitter := range splitters {
			numberOfPaths, ok := heads[splitter[1]]
			if ok {
				heads[splitter[1]-1] += numberOfPaths
				heads[splitter[1]+1] += numberOfPaths

				delete(heads, splitter[1])
			}
		}
	}

	for _, numberOfPaths := range heads {
		result += numberOfPaths
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 7b:", result)

	return nil
}
