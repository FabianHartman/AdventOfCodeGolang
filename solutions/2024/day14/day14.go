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

func MoveRobot(robot *Robot, maxX, maxY int) {
	robot.Position.X = (robot.Position.X + robot.Velocity.X) % maxX
	robot.Position.Y = (robot.Position.Y + robot.Velocity.Y) % maxY

	if robot.Position.X < 0 {
		robot.Position.X += maxX
	}
	if robot.Position.Y < 0 {
		robot.Position.Y += maxY
	}
}

func CalculateChristmasTreeSeconds(robots []Robot, maxX, maxY int) (int, error) {
	for counter := 0; counter >= 0; counter++ {
		for i := range robots {
			MoveRobot(&robots[i], maxX, maxY)
		}

		if isChristmasTree(robots, maxX, maxY) {
			return counter + 1, nil
		}
	}

	return 0, fmt.Errorf("no solution found")
}

// (I don't get it why we check that there are > 10 gridcells with robots in a row
// I used someone else's code as example for part 2
func isChristmasTree(robots []Robot, maxX, maxY int) bool {
	var grid [][]int
	grid = make([][]int, maxY)

	for y := 0; y < maxY; y++ {
		grid[y] = make([]int, maxX)
	}

	for _, robot := range robots {
		grid[robot.Position.Y][robot.Position.X]++
	}

	for y := 0; y < maxY; y++ {
		count := 0

		for x := 0; x < maxX; x++ {
			if grid[y][x] == 1 {
				count++
			}

			if count > 10 {
				return true
			}

			if grid[y][x] == 0 {
				count = 0
			}
		}
	}

	return false
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
	robots, err := input()
	if err != nil {
		return err
	}

	seconds, err := CalculateChristmasTreeSeconds(robots, 101, 103)
	if err != nil {
		return err
	}

	fmt.Println("Day14b:", seconds)

	return nil
}
