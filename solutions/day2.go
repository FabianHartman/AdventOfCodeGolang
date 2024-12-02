package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var inputPath2 string = "inputs/day2.txt"

func Day2a() error {
	file, err := os.Open(inputPath2)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		nums := []int{}
		for _, str := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				return fmt.Errorf("error converting %s to int: %s", str, err)
			}

			nums = append(nums, num)
		}

		reports = append(reports, nums)
	}

	total := 0

	for _, report := range reports {
		value := report[1]
		ascending := report[1] > report[0]

		startingDifference := math.Abs(float64(report[1] - report[0]))

		if startingDifference > 3 || startingDifference < 1 {
			continue
		}

		for i, num := range report[2:] {
			if ascending && num > value && num <= value+3 {
				value = num
			} else if (!ascending) && (num < value && num >= value-3) {
				value = num
			} else {
				break
			}
			if i == len(report)-1-2 {
				total++
				fmt.Println(report)
			}
		}
	}
	fmt.Println("Day 2a:", total)

	return nil
}

func Day2b() error {
	file, err := os.Open(inputPath2)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		nums := []int{}
		for _, str := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				return fmt.Errorf("error converting %s to int: %s", str, err)
			}

			nums = append(nums, num)
		}

		reports = append(reports, nums)
	}

	total := 0

	for _, report := range reports {
		unsafeCount := 0

		var ascending bool

		value := report[1]
		if report[1] != report[0] {
			ascending = report[1] > report[0]
		} else {
			ascending = report[2] > report[1]
			unsafeCount++
		}

		startingDifference := math.Abs(float64(report[1] - report[0]))

		if startingDifference > 3 || startingDifference < 1 {
			continue
		}

		for i, num := range report[1:] {
			if num > value {
				if !ascending {
					unsafeCount++
				} else {
					if !(num > value && num <= value+3) {
						unsafeCount++
					}
				}
				value = num
				ascending = true
			} else if num == value {
				unsafeCount++
			} else if num < value {
				if ascending {
					unsafeCount++
				} else {
					if !(num < value && num >= value-3) {
						unsafeCount++
					}
				}
				value = num
				ascending = false
			}
			if unsafeCount > 2 {
				continue
			}
			if i == len(report)-1-1 {
				total++
				fmt.Println(report)
			}
		}
	}
	fmt.Println("Day 2b:", total)

	return nil
}
