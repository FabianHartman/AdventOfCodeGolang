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

func Day1() error {
	file, err := os.Open("inputs/day1.txt")
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

	fmt.Println(total)

	return nil
}
