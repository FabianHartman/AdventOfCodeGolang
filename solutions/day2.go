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
