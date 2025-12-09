package main

import (
	"adventOfCode/solutions/2025/day9"
	"fmt"

	"adventOfCode/helper/runner"

	"adventOfCode/solutions/2025/day1"
	"adventOfCode/solutions/2025/day2"
	"adventOfCode/solutions/2025/day3"
	"adventOfCode/solutions/2025/day4"
	"adventOfCode/solutions/2025/day5"
	"adventOfCode/solutions/2025/day6"
	"adventOfCode/solutions/2025/day7"
	"adventOfCode/solutions/2025/day8"
)

func main() {
	runner.Run(runAll)
}

func runAll() error {
	runner.Run(day1.Part1)
	runner.Run(day1.Part2)
	runner.Run(day2.Part1)
	runner.Run(day2.Part2)
	runner.Run(day3.Part1)
	runner.Run(day3.Part2)
	runner.Run(day4.Part1)
	runner.Run(day4.Part2)
	runner.Run(day5.Part1)
	runner.Run(day5.Part2)
	runner.Run(day6.Part1)
	runner.Run(day6.Part2)
	runner.Run(day7.Part1)
	runner.Run(day7.Part2)
	runner.Run(day8.Part1)
	runner.Run(day8.Part2)
	runner.Run(day9.Part1)
	runner.Run(day9.Part2)

	fmt.Print("\n\n2025 ")

	return nil
}
