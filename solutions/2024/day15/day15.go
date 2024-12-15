package day15

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputPath string = "inputs/2024/day15.txt"

type Position struct {
	X, Y int
}

type Warehouse struct {
	Robot          *Position
	Boxes, Walls   map[Position]bool
	MaxRow, MaxCol int
}

func (this *Warehouse) visualize() {
	output := [][]string{}

	for y := 0; y < this.MaxRow; y++ {
		output = append(output, []string{})
		for x := 0; x < this.MaxCol; x++ {
			output[y] = append(output[y], ".")
		}
	}

	for box := range this.Boxes {
		output[box.Y][box.X] = "O"
	}

	for wall := range this.Walls {
		output[wall.Y][wall.X] = "#"
	}

	output[this.Robot.Y][this.Robot.X] = "@"

	for _, line := range output {
		fmt.Println(strings.Join(line, ""))
	}
}

func (this *Warehouse) isEmpty(point *Position) bool {
	if _, exists := this.Boxes[*point]; exists {
		return false
	}

	_, exists := this.Walls[*point]

	return !exists
}

func (this *Warehouse) getTouchingWallsAndBoxes(position *Position, instruction string) (*Position, []Position) {
	boxesInBetween := []Position{}

	switch instruction {
	case "<":
		for i := position.X - 1; i >= 0; i-- {
			itteratedPosition := Position{X: i, Y: position.Y}
			if _, exists := this.Walls[itteratedPosition]; exists {
				return &itteratedPosition, boxesInBetween
			}

			if _, exists := this.Boxes[itteratedPosition]; exists {
				boxesInBetween = append(boxesInBetween, itteratedPosition)
				continue
			}

			return nil, boxesInBetween
		}
	case "^":
		for i := position.Y - 1; i >= 0; i-- {
			itteratedPosition := Position{X: position.X, Y: i}
			if _, exists := this.Walls[itteratedPosition]; exists {
				return &itteratedPosition, boxesInBetween
			}

			if _, exists := this.Boxes[itteratedPosition]; exists {
				boxesInBetween = append(boxesInBetween, itteratedPosition)
				continue
			}

			return nil, boxesInBetween
		}
	case ">":
		for i := position.X + 1; i < this.MaxCol; i++ {
			itteratedPosition := Position{X: i, Y: position.Y}

			if _, exists := this.Walls[itteratedPosition]; exists {
				return &itteratedPosition, boxesInBetween
			}

			if _, exists := this.Boxes[itteratedPosition]; exists {
				boxesInBetween = append(boxesInBetween, itteratedPosition)
				continue
			}

			return nil, boxesInBetween
		}
	case "v":
		for i := position.Y + 1; i < this.MaxRow; i++ {
			itteratedPosition := Position{X: position.X, Y: i}
			if _, exists := this.Walls[itteratedPosition]; exists {
				return &itteratedPosition, boxesInBetween
			}

			if _, exists := this.Boxes[itteratedPosition]; exists {
				boxesInBetween = append(boxesInBetween, itteratedPosition)
				continue
			}

			return nil, boxesInBetween
		}
	}

	return nil, nil
}

func (this *Warehouse) executeInstruction(instruction string) {
	var newRobotPos Position

	switch instruction {
	case "<":
		newRobotPos = Position{X: this.Robot.X - 1, Y: this.Robot.Y}
	case "^":
		newRobotPos = Position{X: this.Robot.X, Y: this.Robot.Y - 1}
	case ">":
		newRobotPos = Position{X: this.Robot.X + 1, Y: this.Robot.Y}
	case "v":
		newRobotPos = Position{X: this.Robot.X, Y: this.Robot.Y + 1}
	}

	if this.isEmpty(&newRobotPos) {
		this.Robot = &newRobotPos
	} else {
		touchingWall, touchingBoxes := this.getTouchingWallsAndBoxes(this.Robot, instruction)

		if touchingWall != nil {
			return
		}

		for i := len(touchingBoxes) - 1; i >= 0; i-- {
			box := touchingBoxes[i]
			delete(this.Boxes, box)

			switch instruction {
			case "<":
				this.Boxes[Position{X: box.X - 1, Y: box.Y}] = true
			case "^":
				this.Boxes[Position{X: box.X, Y: box.Y - 1}] = true
			case ">":
				this.Boxes[Position{X: box.X + 1, Y: box.Y}] = true
			case "v":
				this.Boxes[Position{X: box.X, Y: box.Y + 1}] = true
			}
		}

		this.Robot = &newRobotPos
	}
}

func (this *Warehouse) calculateBoxCoordinateValues() int {
	total := 0

	for box, _ := range this.Boxes {
		total += box.X
		total += box.Y * 100
	}

	return total
}

func input() (*Warehouse, string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, "", fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	warehouseFinished := false
	warehouseRow := -1
	warehouse := Warehouse{Robot: nil, Boxes: map[Position]bool{}, Walls: map[Position]bool{}, MaxRow: -1, MaxCol: -1}
	instructions := ""

	for scanner.Scan() {
		warehouse.MaxRow++

		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			warehouseFinished = true
			continue
		}

		if !warehouseFinished {
			warehouseRow++

			gridcells := strings.Split(line, "")

			if warehouse.MaxCol == -1 {
				warehouse.MaxCol = len(gridcells)
			}

			for warehouseColumn, value := range gridcells {
				switch value {
				case "#":
					warehouse.Walls[Position{X: warehouseColumn, Y: warehouseRow}] = true
				case "O":
					warehouse.Boxes[Position{X: warehouseColumn, Y: warehouseRow}] = true
				case "@":
					warehouse.Robot = &Position{X: warehouseColumn, Y: warehouseRow}
				}
			}
		} else {
			instructions += strings.TrimSpace(line)
		}
	}

	return &warehouse, instructions, nil
}

func Part1() error {
	warehouse, instructions, err := input()
	if err != nil {
		return err
	}

	for _, instruction := range strings.Split(instructions, "") {
		warehouse.executeInstruction(instruction)
		//warehouse.visualize()
	}

	fmt.Println("Day 15a:", warehouse.calculateBoxCoordinateValues())

	return nil
}
