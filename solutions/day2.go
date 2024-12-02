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

func input() ([][]int, error) {
	file, err := os.Open(inputPath2)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		nums := []int{}
		for _, str := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("error converting %s to int: %s", str, err)
			}

			nums = append(nums, num)
		}

		reports = append(reports, nums)
	}

	return reports, nil
}

func isValid(report []int) bool {
	value := report[1]
	ascending := report[1] > report[0]

	startingDifference := math.Abs(float64(report[1] - report[0]))

	if startingDifference > 3 || startingDifference < 1 {
		return false
	}

	for i, num := range report[2:] {
		if ascending && num > value && num <= value+3 {
			value = num
		} else if (!ascending) && (num < value && num >= value-3) {
			value = num
		} else {
			return false
		}
		if i == len(report)-1-2 {
			return true
		}
	}

	return false
}

func generateMutations(report []int) [][]int {
	mutations := [][]int{}

	for i := range report {
		mutatedReport := []int{}
		if i == 0 {
			mutatedReport = report[1:]
		} else {
			mutatedReport = append(mutatedReport, report[:i]...)
			mutatedReport = append(mutatedReport, report[i+1:]...)
		}

		mutations = append(mutations, mutatedReport)
	}

	return mutations
}

func Day2a() error {
	reports, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, report := range reports {
		if isValid(report) {
			total++
		}
	}

	fmt.Println("Day 2a:", total)

	return nil
}

func Day2b() error {
	reports, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, report := range reports {
		if isValid(report) {
			total++
			continue
		}
		for _, mutatedReport := range generateMutations(report) {
			if isValid(mutatedReport) {
				total++
				break
			}
		}

	}

	fmt.Println("Day 2b:", total)

	return nil
}
