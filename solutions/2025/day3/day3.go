package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputPath = "inputs/2025/day3.txt"

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

	for scanner.Scan() {
		inputRow := scanner.Text()

		highest := 0

		for startI := 0; startI < len(inputRow); startI++ {
			for endI := startI + 1; endI < len(inputRow); endI++ {
				value, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(inputRow[startI]), string(inputRow[endI])))
				if value > highest {
					highest = value
				}
			}
		}

		result += highest
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 3a:", result)

	return nil
}

func Part2() error {
	var totalSum int

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

		highestDigits := make([]int, 12)
		skipOffset := len(line) - 12

		for charIndex, char := range line {
			digit, _ := strconv.Atoi(string(char))

			for slotIndex := max(charIndex-skipOffset, 0); slotIndex < 12; slotIndex++ {
				if digit > highestDigits[slotIndex] {
					highestDigits[slotIndex] = digit

					for nextSlot := slotIndex + 1; nextSlot < 12; nextSlot++ {
						if highestDigits[nextSlot] == 0 {
							break
						}

						highestDigits[nextSlot] = 0
					}

					break
				}
			}
		}

		numberStr := ""
		for i := 0; i < 12; i++ {
			numberStr += strconv.Itoa(highestDigits[i])
		}

		number, _ := strconv.Atoi(numberStr)
		totalSum += number
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 3b:", totalSum)
	return nil
}
