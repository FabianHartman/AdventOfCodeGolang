package day14

import (
	"bufio"
	"fmt"
	"os"
)

var inputPath string = "inputs/2024/day14.txt"

type Position struct {
	X, Y int
}

type Velocity struct {
	X, Y int
}

type Robot struct {
	Position *Position
	Velocity *Velocity
}

func input() ([]Robot, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	robots := []Robot{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		robot := Robot{Position: &Position{}, Velocity: &Velocity{}}

		_, err = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.Position.X, &robot.Position.Y, &robot.Velocity.X, &robot.Velocity.Y)
		if err != nil {
			return nil, fmt.Errorf("error scanning line: %s", err)
		}

		robots = append(robots, robot)
	}

	return robots, nil
}

func (this *Robot) getPositionAfterNSeconds(seconds int, width int, tall int) *Position {
	position := Position{X: this.Position.X, Y: this.Position.Y}

	position.X += this.Velocity.X * seconds
	position.Y += this.Velocity.Y * seconds

	position.X = (position.X + 1000*width) % width
	position.Y = (position.Y + 1000*tall) % tall

	fmt.Println(position)

	return &position
}

// Get the Quadrant of the position, Quadrants are
// 1  2
// 3  4
//
// 0 Indicates that it is somewhere in the middle
func (this *Position) getQuadrant(width int, tall int) int {
	if float64(this.X) < float64(width-1)/float64(2) {
		if float64(this.Y) < float64(tall-1)/float64(2) {
			return 1
		} else if float64(this.Y) > float64(tall-1)/float64(2) {
			return 3
		}
	} else if float64(this.X) > float64(width-1)/float64(2) {
		if float64(this.Y) < float64(tall-1)/float64(2) {
			return 2
		} else if float64(this.Y) > float64(tall-1)/float64(2) {
			return 4
		}
	}

	return 0
}

func Part1() error {
	width, height := 101, 103
	robots, err := input()
	if err != nil {
		return err
	}

	quadrantRobotCounts := map[int]int{}

	for _, robot := range robots {
		quadrantRobotCounts[robot.getPositionAfterNSeconds(100, width, height).getQuadrant(width, height)]++
	}

	delete(quadrantRobotCounts, 0)

	total := 1

	for _, count := range quadrantRobotCounts {
		total *= count
	}

	fmt.Println("Day 14a:", total)

	return nil
}

func Part2() error {
	return nil
}
