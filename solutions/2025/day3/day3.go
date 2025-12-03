package day3

import (
	"bufio"
	"fmt"
	"math"
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

		bankJoltages := make([]int, 12)
		lastBatteryWithoutSkip := len(inputRow) - 12

		for batteryIndex, thisBatteryJoltage := range inputRow {
			firstValidIndex := max(batteryIndex-lastBatteryWithoutSkip, 0)

			n, _ := strconv.Atoi(string(thisBatteryJoltage))

			for j := firstValidIndex; j < 12; j++ {
				if n > bankJoltages[j] {
					bankJoltages[j] = n

					for k := j + 1; k < 12; k++ {
						if bankJoltages[k] == 0 {
							break
						}

						bankJoltages[k] = 0
					}

					break
				}
			}
		}

		bankJoltage := 0
		for j := range 12 {
			bankJoltage += bankJoltages[j] * int(math.Pow10(12-1-j))
		}
		result += bankJoltage
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 3b:", result)

	return nil
}
