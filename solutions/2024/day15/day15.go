package day15

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputPath string = "inputs/2024/day15.txt"

type Position struct {
	Y int
	X int
}

type WideBox struct {
	LeftPart  *Position
	RightPart *Position
}

type BigWarehouse struct {
	Boxes         map[WideBox]bool
	BoxParts      map[Position]WideBox
	Walls         map[Position]bool
	Robot         *Position
	Width, Height int
}

type Warehouse struct {
	moveSeq       string
	Boxes, Walls  map[Position]bool
	robot         *Position
	Width, Height int
}

func (this *Warehouse) convertToBigWarehouse() *BigWarehouse {
	var bigWarehouse BigWarehouse

	bigWarehouse.Height = this.Height
	bigWarehouse.Width = this.Width * 2
	bigWarehouse.Robot = &Position{this.robot.Y, this.robot.X * 2}

	bigWarehouse.Walls = make(map[Position]bool)

	for wall := range this.Walls {
		leftWallPart := Position{wall.Y, wall.X * 2}
		rightWallPart := Position{leftWallPart.Y, leftWallPart.X + 1}

		bigWarehouse.Walls[leftWallPart] = true
		bigWarehouse.Walls[rightWallPart] = true
	}

	bigWarehouse.Boxes = make(map[WideBox]bool)
	bigWarehouse.BoxParts = make(map[Position]WideBox)

	for box := range this.Boxes {
		leftBoxPart := Position{box.Y, box.X * 2}
		rightBoxPart := Position{box.Y, leftBoxPart.X + 1}

		wideBox := WideBox{&leftBoxPart, &rightBoxPart}

		bigWarehouse.BoxParts[leftBoxPart] = wideBox
		bigWarehouse.BoxParts[rightBoxPart] = wideBox

		bigWarehouse.Boxes[wideBox] = true
	}

	return &bigWarehouse
}

func input() (*Warehouse, string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, "", fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var warehouse Warehouse

	warehouse.Boxes = make(map[Position]bool)
	warehouse.Walls = make(map[Position]bool)

	scanner.Scan()
	topWalls := scanner.Text()

	for i := range len(topWalls) {
		warehouse.Walls[Position{0, i}] = true
	}

	rowIndex := 1

	for scanner.Scan() {
		line := scanner.Text()

		if line == topWalls {
			for columnIndex := range len(topWalls) {
				warehouse.Walls[Position{rowIndex, columnIndex}] = true
			}

			break
		}

		for col, char := range strings.Split(line, "") {
			switch char {
			case "#":
				warehouse.Walls[Position{rowIndex, col}] = true
				break
			case "O":
				warehouse.Boxes[Position{rowIndex, col}] = true
				break
			case "@":
				warehouse.robot = &Position{rowIndex, col}
				break
			}
		}

		rowIndex++
	}

	warehouse.Height = rowIndex + 1
	warehouse.Width = len(topWalls)

	instructions := ""

	scanner.Scan()

	for scanner.Scan() {
		instructions += scanner.Text()
	}

	return &warehouse, instructions, nil
}

func canBoxMove(house *Warehouse, box *Position, dir string) bool {
	nextCoords := box.getNextPosition(dir)
	if _, exists := house.Walls[*nextCoords]; exists {
		return false
	}

	if _, exists := house.Boxes[*nextCoords]; exists {
		return canBoxMove(house, nextCoords, dir)
	}

	return true
}

func (this *Warehouse) moveBoxes(box *Position, dir string) {
	nextCoords := box.getNextPosition(dir)

	if _, exists := this.Walls[*nextCoords]; exists {
		return
	}

	if _, exists := this.Boxes[*nextCoords]; exists {
		this.moveBoxes(nextCoords, dir)
	}

	delete(this.Boxes, *box)
	this.Boxes[*nextCoords] = true
}

func (this *Warehouse) move(dir string) {
	nextCoords := this.robot.getNextPosition(dir)

	if _, exists := this.Walls[*nextCoords]; exists {
		return
	}

	if _, exists := this.Boxes[*nextCoords]; exists {
		if canBoxMove(this, nextCoords, dir) {
			this.moveBoxes(nextCoords, dir)
		} else {
			return
		}
	}

	this.robot = nextCoords
}

func (this *Warehouse) solve(instructions string) int {
	for _, instruction := range strings.Split(instructions, "") {
		this.move(instruction)
	}

	total := 0

	for box := range this.Boxes {
		total += 100*box.Y + box.X
	}

	return total
}

func (this *Position) getNextPosition(dir string) *Position {
	nextPosition := Position{this.Y, this.X}

	switch dir {
	case "^":
		nextPosition.Y--
	case ">":
		nextPosition.X++
	case "v":
		nextPosition.Y++
	default:
		nextPosition.X--
	}

	return &nextPosition
}

func (this *BigWarehouse) canWideBoxMove(side *Position, dir string) bool {
	canMove := true
	wideBox := this.BoxParts[*side]

	leftNext := wideBox.LeftPart.getNextPosition(dir)
	rightNext := wideBox.RightPart.getNextPosition(dir)

	_, leftExists := this.Walls[*leftNext]
	_, rightExists := this.Walls[*rightNext]

	if leftExists || rightExists {
		return false
	}

	if dir == "<" {
		if _, leftExists := this.BoxParts[*leftNext]; leftExists {
			canMove = this.canWideBoxMove(leftNext, dir)
		}

		return canMove
	}

	if dir == ">" {
		if _, rightExists := this.BoxParts[*rightNext]; rightExists {
			canMove = this.canWideBoxMove(rightNext, dir)
		}

		return canMove
	}

	wideBoxLeft, LeftOK := this.BoxParts[*leftNext]
	wideBoxRight, RightOK := this.BoxParts[*rightNext]

	if LeftOK {
		canMove = canMove && this.canWideBoxMove(leftNext, dir)
	}
	if RightOK && wideBoxLeft != wideBoxRight {
		canMove = canMove && this.canWideBoxMove(rightNext, dir)
	}

	return canMove
}

func (this *BigWarehouse) wideBox(side *Position, dir string) {
	wideBox := this.BoxParts[*side]
	left, right := wideBox.LeftPart, wideBox.RightPart
	leftNext, rightNext := left.getNextPosition(dir), right.getNextPosition(dir)

	if dir == "<" {
		if _, leftExists := this.BoxParts[*leftNext]; leftExists {
			this.wideBox(leftNext, dir)
		}

		delete(this.Boxes, wideBox)
		delete(this.BoxParts, *left)
		delete(this.BoxParts, *right)

		wideBox.RightPart = wideBox.LeftPart
		wideBox.LeftPart = leftNext

		this.Boxes[wideBox] = true
		this.BoxParts[*left] = wideBox
		this.BoxParts[*leftNext] = wideBox

		return
	}

	if dir == ">" {
		if _, rightExists := this.BoxParts[*rightNext]; rightExists {
			this.wideBox(rightNext, dir)
		}

		delete(this.Boxes, wideBox)
		delete(this.BoxParts, *left)
		delete(this.BoxParts, *right)

		wideBox.LeftPart = wideBox.RightPart
		wideBox.RightPart = rightNext

		this.Boxes[wideBox] = true
		this.BoxParts[*right] = wideBox
		this.BoxParts[*rightNext] = wideBox

		return
	}

	leftPart, leftExists := this.BoxParts[*leftNext]
	rightPart, rightExists := this.BoxParts[*rightNext]

	if leftExists {
		this.wideBox(leftNext, dir)
	}
	if rightExists && leftPart != rightPart {
		this.wideBox(rightNext, dir)
	}

	delete(this.Boxes, wideBox)
	delete(this.BoxParts, *left)
	delete(this.BoxParts, *right)

	wideBox.LeftPart = leftNext
	wideBox.RightPart = rightNext

	this.Boxes[wideBox] = true
	this.BoxParts[*leftNext] = wideBox
	this.BoxParts[*rightNext] = wideBox
}

func (this *BigWarehouse) bigMove(dir string) {
	nextCoords := this.Robot.getNextPosition(dir)

	if _, exists := this.Walls[*nextCoords]; exists {
		return
	}

	if _, exists := this.BoxParts[*nextCoords]; exists {
		if this.canWideBoxMove(nextCoords, dir) {
			this.wideBox(nextCoords, dir)
		} else {
			return
		}
	}

	this.Robot = nextCoords
}

func (this *BigWarehouse) solve(instructions string) int {
	for _, instruction := range strings.Split(instructions, "") {
		this.bigMove(instruction)
	}

	total := 0

	for wideBox := range this.Boxes {
		total += 100*wideBox.LeftPart.Y + wideBox.LeftPart.X
	}

	return total
}

func Part1() error {
	warehouse, instructions, err := input()
	if err != nil {
		return err
	}

	fmt.Println("Day 15a:", warehouse.solve(instructions))

	return nil
}

func Part2() error {
	warehouse, instructions, err := input()
	if err != nil {
		return err
	}

	bigWarehouse := warehouse.convertToBigWarehouse()

	fmt.Println("Day 15b:", bigWarehouse.solve(instructions))

	return nil
}
