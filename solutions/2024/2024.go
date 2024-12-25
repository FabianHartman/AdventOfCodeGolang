package _2024

import (
	"fmt"
	"time"

	"adventOfCode/solutions/2024/day1"
	"adventOfCode/solutions/2024/day10"
	"adventOfCode/solutions/2024/day11"
	"adventOfCode/solutions/2024/day12"
	"adventOfCode/solutions/2024/day13"
	"adventOfCode/solutions/2024/day14"
	"adventOfCode/solutions/2024/day15"
	"adventOfCode/solutions/2024/day16"
	"adventOfCode/solutions/2024/day17"
	"adventOfCode/solutions/2024/day18"
	"adventOfCode/solutions/2024/day19"
	"adventOfCode/solutions/2024/day2"
	"adventOfCode/solutions/2024/day20"
	"adventOfCode/solutions/2024/day21"
	"adventOfCode/solutions/2024/day22"
	"adventOfCode/solutions/2024/day23"
	"adventOfCode/solutions/2024/day24"
	"adventOfCode/solutions/2024/day25"
	"adventOfCode/solutions/2024/day3"
	"adventOfCode/solutions/2024/day4"
	"adventOfCode/solutions/2024/day5"
	"adventOfCode/solutions/2024/day6"
	"adventOfCode/solutions/2024/day7"
	"adventOfCode/solutions/2024/day8"
	"adventOfCode/solutions/2024/day9"
)

func measureTime(day, part string, fn func() error) {
	start := time.Now()
	err := fn()
	duration := time.Since(start)
	if err != nil {
		panic(fmt.Sprintf("%s %s failed: %v", day, part, err))
	}
	fmt.Printf("%s %s took %v\n\n", day, part, duration)
}

func RunAll() {
	start := time.Now()
	// Day 1
	measureTime("Day 1", "Part 1", day1.Part1)
	measureTime("Day 1", "Part 2", day1.Part2)

	// Day 2
	measureTime("Day 2", "Part 1", day2.Part1)
	measureTime("Day 2", "Part 2", day2.Part2)

	// Day 3
	measureTime("Day 3", "Part 1", day3.Part1)
	measureTime("Day 3", "Part 2", day3.Part2)

	// Day 4
	measureTime("Day 4", "Part 1", day4.Part1)
	measureTime("Day 4", "Part 2", day4.Part2)

	// Day 5
	measureTime("Day 5", "Part 1", day5.Part1)
	measureTime("Day 5", "Part 2", day5.Part2)

	// Day 6
	measureTime("Day 6", "Part 1", day6.Part1)
	measureTime("Day 6", "Part 2", day6.Part2)

	// Day 7
	measureTime("Day 7", "Part 1", day7.Part1)
	measureTime("Day 7", "Part 2", day7.Part2)

	// Day 8
	measureTime("Day 8", "Part 1", day8.Part1)
	measureTime("Day 8", "Part 2", day8.Part2)

	// Day 9
	measureTime("Day 9", "Part 1", day9.Part1)
	measureTime("Day 9", "Part 2", day9.Part2)

	// Day 10
	measureTime("Day 10", "Part 1", day10.Part1)
	measureTime("Day 10", "Part 2", day10.Part2)

	// Day 11
	measureTime("Day 11", "Part 1", day11.Part1)
	measureTime("Day 11", "Part 2", day11.Part2)

	// Day 12
	measureTime("Day 12", "Part 1", day12.Part1)
	measureTime("Day 12", "Part 2", day12.Part2)

	// Day 13
	measureTime("Day 13", "Part 1", day13.Part1)
	measureTime("Day 13", "Part 2", day13.Part2)

	// Day 14
	measureTime("Day 14", "Part 1", day14.Part1)
	measureTime("Day 14", "Part 2", day14.Part2)

	// Day 15
	measureTime("Day 15", "Part 1", day15.Part1)
	measureTime("Day 15", "Part 2", day15.Part2)

	// Day 16
	measureTime("Day 16", "Part 1", day16.Part1)
	measureTime("Day 16", "Part 2", day16.Part2)

	// Day 17
	measureTime("Day 17", "Part 1", day17.Part1)
	measureTime("Day 17", "Part 2", day17.Part2)

	// Day 18
	measureTime("Day 18", "Part 1", day18.Part1)
	measureTime("Day 18", "Part 2", day18.Part2)

	// Day 19
	measureTime("Day 19", "Part 1", day19.Part1)
	measureTime("Day 19", "Part 2", day19.Part2)

	// Day 20
	measureTime("Day 20", "Part 1", day20.Part1)
	measureTime("Day 20", "Part 2", day20.Part2)

	// Day 21
	measureTime("Day 21", "Part 1", day21.Part1)
	measureTime("Day 21", "Part 2", day21.Part2)

	// Day 22
	measureTime("Day 22", "Part 1", day22.Part1)
	measureTime("Day 22", "Part 2", day22.Part2)

	// Day 23
	measureTime("Day 23", "Part 1", day23.Part1)
	measureTime("Day 23", "Part 2", day23.Part2)

	// Day 24
	measureTime("Day 24", "Part 1", day24.Part1)
	measureTime("Day 24", "Part 2", day24.Part2)

	// Day 25
	measureTime("Day 25", "Solution", day25.Solution)

	fmt.Println("\n2024 took", time.Since(start))
}
