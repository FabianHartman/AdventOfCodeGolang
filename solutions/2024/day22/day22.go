package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputPath = "inputs/2024/day22.txt"

type BananaDeltaPair struct {
	Delta  int
	Amount int
}

type BananaWindow struct {
	FirstValue, SecondValue, ThirdValue, FourthValue int
}

func input() ([]int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	initials := []int{}

	for scanner.Scan() {
		initial, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Error converting %s to int", scanner.Text())
		}

		initials = append(initials, initial)
	}

	return initials, nil
}

func getNthSecretNumber(secret int, iterations int) int {
	for ; iterations > 0; iterations-- {
		secret = ((secret << 6) ^ secret) & (16777216 - 1)
		secret = ((secret >> 5) ^ secret) & (16777216 - 1)
		secret = ((secret << 11) ^ secret) & (16777216 - 1)
	}

	return secret
}

func addInstructionsAndResults(secret int, iterations int, instructionsAndResults *map[BananaWindow]int) []BananaDeltaPair {
	visited := make(map[BananaWindow]bool)
	deltas := make([]BananaDeltaPair, 0, iterations-1)
	currentWindow := make([]int, 0, 8)
	prev := secret % 10

	for ; iterations > 1; iterations-- {
		secret = ((secret << 6) ^ secret) & (16777216 - 1)
		secret = ((secret >> 5) ^ secret) & (16777216 - 1)
		secret = ((secret << 11) ^ secret) & (16777216 - 1)

		delta, bananas := secret%10-prev, secret%10

		deltas = append(deltas, BananaDeltaPair{Delta: delta, Amount: bananas})
		currentWindow = append(currentWindow, delta)

		if len(currentWindow) == 4 {
			key := BananaWindow{FirstValue: currentWindow[0], SecondValue: currentWindow[1], ThirdValue: currentWindow[2], FourthValue: currentWindow[3]}
			if !visited[key] {
				(*instructionsAndResults)[key] += secret % 10
				visited[key] = true
			}

			currentWindow = currentWindow[1:]
		}

		prev = secret % 10
	}

	return deltas
}

func Part1() error {
	initialSecrets, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, num := range initialSecrets {
		total += getNthSecretNumber(num, 2000)
	}

	fmt.Println("Day22a:", total)

	return nil
}

func Part2() error {
	initialSecrets, err := input()
	if err != nil {
		return err
	}

	instructionsAndResults := make(map[BananaWindow]int)

	for _, num := range initialSecrets {
		addInstructionsAndResults(num, 2000, &instructionsAndResults)
	}

	result := 0

	for _, instructionResult := range instructionsAndResults {
		if instructionResult > result {
			result = instructionResult
		}
	}

	fmt.Println("Day22b:", result)

	return nil
}
