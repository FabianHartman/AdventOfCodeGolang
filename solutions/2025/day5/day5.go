package day5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func minIngredientID(a, b IngredientID) IngredientID {
	if a < b {
		return a
	}

	return b
}

func maxIngredientID(a, b IngredientID) IngredientID {
	if a > b {
		return a
	}

	return b
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

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Begin < ranges[j].Begin
	})

	var mergedRanges []Range

	for _, ingredientRange := range ranges {
		merged := false
		for i := range mergedRanges {
			if ingredientRange.Begin <= mergedRanges[i].End+1 && ingredientRange.End >= mergedRanges[i].Begin-1 {
				mergedRanges[i].Begin = minIngredientID(ingredientRange.Begin, mergedRanges[i].Begin)
				mergedRanges[i].End = maxIngredientID(ingredientRange.End, mergedRanges[i].End)

				merged = true

				break
			}
		}

		if !merged {
			mergedRanges = append(mergedRanges, ingredientRange)
		}
	}

	result := 0
	for _, ingredientRange := range mergedRanges {
		result += int(ingredientRange.End) - int(ingredientRange.Begin) + 1
	}

	fmt.Println("Day 5b:", result)

	return nil
}
