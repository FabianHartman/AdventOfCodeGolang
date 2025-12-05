package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "inputs/2025/day5.txt"

type IngredientID int

type Range struct {
	Begin IngredientID
	End   IngredientID
}

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

	var ranges []Range

	for scanner.Scan() {
		inputRow := scanner.Text()

		if inputRow == "" {
			break
		}

		rangeParts := strings.Split(inputRow, "-")
		if len(rangeParts) != 2 {
			return fmt.Errorf("invalid input for range: %s", inputRow)
		}

		begin, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			return fmt.Errorf("invalid input for range: %s", rangeParts[0])
		}

		end, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			return fmt.Errorf("invalid input for range: %s", rangeParts[1])
		}

		ranges = append(ranges, Range{
			Begin: IngredientID(begin),
			End:   IngredientID(end),
		})
	}

	var ingredients []IngredientID

	for scanner.Scan() {
		inputRow := scanner.Text()

		id, err := strconv.Atoi(inputRow)
		if err != nil {
			return fmt.Errorf("invalid ingredient ID: %s", inputRow)
		}

		ingredients = append(ingredients, IngredientID(id))
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	for _, ingredient := range ingredients {
		for _, ingredientRange := range ranges {
			if ingredient >= ingredientRange.Begin && ingredient <= ingredientRange.End {
				result++

				break
			}
		}
	}

	fmt.Println("Day 5a:", result)

	return nil
}

func Part2() error {
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

	freshIngredients := make(map[IngredientID]bool)

	for scanner.Scan() {
		inputRow := scanner.Text()

		if inputRow == "" {
			break
		}

		rangeParts := strings.Split(inputRow, "-")
		if len(rangeParts) != 2 {
			return fmt.Errorf("invalid input for range: %s", inputRow)
		}

		begin, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			return fmt.Errorf("invalid input for range: %s", rangeParts[0])
		}

		end, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			return fmt.Errorf("invalid input for range: %s", rangeParts[1])
		}

		for i := begin; i <= end; i++ {
			freshIngredients[IngredientID(i)] = true
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 5b:", len(freshIngredients))

	return nil
}
