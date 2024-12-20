package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputPath string = "inputs/2024/day19.txt"

func input() ([]string, []string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	availablePatterns := strings.Split(scanner.Text(), ", ")
	scanner.Scan()

	desiredDesigns := []string{}

	for scanner.Scan() {
		desiredDesigns = append(desiredDesigns, scanner.Text())
	}

	return availablePatterns, desiredDesigns, nil
}

func isDesignPossible(design string, availablePatterns []string) bool {
	canConstructTillIndex := make([]bool, len(design)+1)
	canConstructTillIndex[0] = true

	for i := 1; i <= len(design); i++ {
		for _, towel := range availablePatterns {
			if i-len(towel) >= 0 && canConstructTillIndex[i-len(towel)] && design[i-len(towel):i] == towel {
				canConstructTillIndex[i] = true
				break
			}
		}
	}

	return canConstructTillIndex[len(design)]
}

func calculateAmountOfPossibleDesigns(availablePatterns []string, desiredDesigns []string) int {
	possibleCount := 0

	for _, desiredDesign := range desiredDesigns {
		if isDesignPossible(desiredDesign, availablePatterns) {
			possibleCount++
		}
	}

	return possibleCount
}

func countWaysToFormDesigns(towelPatterns []string, designs []string) int {
	towelSet := make(map[string]bool)
	for _, towel := range towelPatterns {
		towelSet[towel] = true
	}

	totalWays := 0

	for _, design := range designs {
		designLength := len(design)
		amountOfWays := make([]int, designLength+1)
		amountOfWays[0] = 1

		for checkTillLength := 1; checkTillLength <= designLength; checkTillLength++ {
			for currentCheckLength := 1; currentCheckLength <= checkTillLength; currentCheckLength++ {
				if towelSet[design[checkTillLength-currentCheckLength:checkTillLength]] {
					amountOfWays[checkTillLength] += amountOfWays[checkTillLength-currentCheckLength]
				}
			}
		}

		totalWays += amountOfWays[designLength]
	}

	return totalWays
}

func Part1() error {
	availablePatterns, desiredDesigns, err := input()
	if err != nil {
		return err
	}

	possibleAmountOfDesigns := calculateAmountOfPossibleDesigns(availablePatterns, desiredDesigns)

	fmt.Println("Day 19a:", possibleAmountOfDesigns)

	return nil
}

func Part2() error {
	availablePatterns, desiredDesigns, err := input()
	if err != nil {
		return err
	}

	possibleAmountOfDesigns := countWaysToFormDesigns(availablePatterns, desiredDesigns)

	fmt.Println("Day 19b:", possibleAmountOfDesigns)

	return nil
}
