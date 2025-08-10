package day25

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputPath = "inputs/2024/day25.txt"

func input() ([][]int, [][]int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, fmt.Errorf("open file error: %s", err)
	}

	defer file.Close()

	var data [][]string
	var group []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			data = append(data, group)
			group = []string{}
		} else {
			group = append(group, line)
		}
	}

	data = append(data, group)

	locks, keys := [][]int{}, [][]int{}

	for _, lockOrKey := range data {
		if strings.Count(lockOrKey[0], ".") == 0 {
			locks = append(locks, countHashtags(lockOrKey))
		}

		if strings.Count(lockOrKey[len(lockOrKey)-1], ".") == 0 {
			keys = append(keys, countHashtags(lockOrKey))
		}
	}

	return locks, keys, nil
}

func countHashtags(grid []string) []int {
	hashtagCountsPerCol := make([]int, len(grid[0]))

	for _, row := range grid {
		for col := range row {
			if row[col] == '#' {
				hashtagCountsPerCol[col]++
			}
		}
	}

	return hashtagCountsPerCol
}

func keyLockPairFits(lock []int, key []int) bool {
	for i := range lock {
		if lock[i]+key[i] > 7 {
			return false
		}
	}

	return true
}

func Solution() error {
	locks, keys, err := input()
	if err != nil {
		return err
	}

	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if keyLockPairFits(lock, key) {
				count++
			}
		}
	}

	fmt.Println("Day 25:", count)

	return nil
}
