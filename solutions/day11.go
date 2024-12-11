package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPathDay11 string = "inputs/day11.txt"

func inputDay11() ([]int, error) {
	file, err := os.Open(inputPathDay11)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	textNumbers := strings.Split(scanner.Text(), " ")
	for _, textNumber := range textNumbers {
		number, err := strconv.Atoi(textNumber)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to int: %s", textNumber, err)
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func Day11a() error {
	numbers, err := inputDay11()
	if err != nil {
		return err
	}

	for i := 0; i < 25; i++ {
		newNumbers := []int{}
		for _, number := range numbers {
			if number == 0 {
				newNumbers = append(newNumbers, 1)
				continue
			}

			stringNumber := strconv.Itoa(number)

			if len(stringNumber)%2 == 0 {
				if stone, err := strconv.Atoi(stringNumber[:len(stringNumber)/2]); err == nil {
					newNumbers = append(newNumbers, stone)
				} else {
					return err
				}

				if stone, err := strconv.Atoi(stringNumber[len(stringNumber)/2:]); err == nil {
					newNumbers = append(newNumbers, stone)
				} else {
					return err
				}
			} else {
				newNumbers = append(newNumbers, number*2024)
			}
		}

		numbers = newNumbers
	}

	fmt.Println("day11a:", len(numbers))

	return nil
}
