package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputPath1 string = "inputs/day1.txt"

func Day1a() error {
	file, err := os.Open(inputPath1)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstList := []int{}
	secondList := []int{}

	for scanner.Scan() {
		cols := strings.Split(scanner.Text(), "   ")

		first, _ := strconv.Atoi(cols[0])
		second, _ := strconv.Atoi(cols[1])

		firstList = append(firstList, first)
		secondList = append(secondList, second)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	total := 0

	for i := 0; i < len(firstList); i++ {
		total += int(math.Abs(float64(firstList[i] - secondList[i])))
	}

	fmt.Println("Day 1a:", total)

	return nil
}

func Day1b() error {
	file, err := os.Open(inputPath1)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstList := []int{}
	secondCounts := map[int]int{}

	for scanner.Scan() {
		cols := strings.Split(scanner.Text(), "   ")

		first, _ := strconv.Atoi(cols[0])
		second, _ := strconv.Atoi(cols[1])

		firstList = append(firstList, first)
		secondCounts[second]++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	total := 0

	for i := 0; i < len(firstList); i++ {
		number := firstList[i]
		total += secondCounts[number] * number
	}

	fmt.Println("Day 1b:", total)

	return nil
}
