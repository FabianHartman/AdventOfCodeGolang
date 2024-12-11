package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPathDay11 string = "inputs/day11.txt"

func inputDay11() (map[int]int, error) {
	file, err := os.Open(inputPathDay11)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	numbers := map[int]int{}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	textNumbers := strings.Split(scanner.Text(), " ")
	for _, textNumber := range textNumbers {
		number, err := strconv.Atoi(textNumber)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to int: %s", textNumber, err)
		}

		numbers[number]++
	}

	return numbers, nil
}

func solve(n int) (int, error) {
	numbers, err := inputDay11()
	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {
		newNumbers := map[int]int{}
		for number, amount := range numbers {
			if number == 0 {
				newNumbers[1] += amount
				continue
			}

			stringNumber := strconv.Itoa(number)

			if len(stringNumber)%2 == 0 {
				if stone, err := strconv.Atoi(stringNumber[:len(stringNumber)/2]); err == nil {
					newNumbers[stone] += amount
				} else {
					return 0, err
				}

				if stone, err := strconv.Atoi(stringNumber[len(stringNumber)/2:]); err == nil {
					newNumbers[stone] += amount
				} else {
					return 0, err
				}
			} else {
				newNumbers[number*2024] += amount
			}
		}

		numbers = newNumbers
	}

	total := 0

	for _, amount := range numbers {
		total += amount
	}

	return total, nil
}

func Day11a() error {
	total, err := solve(25)
	if err != nil {
		return err
	}

	fmt.Println("Day11a:", total)

	return nil
}

func Day11b() error {
	total, err := solve(75)
	if err != nil {
		return err
	}

	fmt.Println("Day11a:", total)

	return nil
}
