package day12

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputPath = "inputs/2025/day12.txt"

type Coordinate struct {
	Row    int
	Column int
}

type PresentShape struct {
	Index         int
	OccupiedCells map[Coordinate]struct{}
}

type GiftRegion struct {
	Dimensions             Coordinate
	GiftShapeQuotasByIndex map[int]int
}

func (this *GiftRegion) canFitAllShapesInQuota(presentShapesByIndex map[int]*PresentShape) bool {
	type ShapeWithSize struct {
		Shape *PresentShape
		Size  int
	}

	var shapesWithSizes []ShapeWithSize
	totalShapeArea := 0
	for idx, count := range this.GiftShapeQuotasByIndex {
		shape := presentShapesByIndex[idx]
		shapeArea := len(shape.OccupiedCells)
		for i := 0; i < count; i++ {
			shapesWithSizes = append(shapesWithSizes, ShapeWithSize{shape, shapeArea})
			totalShapeArea += shapeArea
		}
	}

	gridArea := this.Dimensions.Row * this.Dimensions.Column
	if totalShapeArea > gridArea {
		return false
	}

	slices.SortFunc(shapesWithSizes, func(a, b ShapeWithSize) int {
		return b.Size - a.Size
	})

	var shapesToPlace []*PresentShape
	for _, shapeWithSize := range shapesWithSizes {
		shapesToPlace = append(shapesToPlace, shapeWithSize.Shape)
	}

	grid := make([][]bool, this.Dimensions.Row)
	for i := range grid {
		grid[i] = make([]bool, this.Dimensions.Column)
	}

	return tryPlaceShapes(grid, shapesToPlace, presentShapesByIndex, 0, gridArea-totalShapeArea)
}

func tryPlaceShapes(grid [][]bool, shapesToPlace []*PresentShape, allShapes map[int]*PresentShape, shapeIndex int, remainingFreeArea int) bool {
	if shapeIndex >= len(shapesToPlace) {
		return true
	}

	shape := shapesToPlace[shapeIndex]
	rotations := getUniqueRotations(shape)

	startRow, startCol := findFirstEmptyCell(grid)
	if startRow == -1 {
		return false
	}

	for row := startRow; row < len(grid); row++ {
		colStart := 0
		if row == startRow {
			colStart = startCol
		}

		for col := colStart; col < len(grid[0]); col++ {
			if !grid[row][col] {
				for _, rotated := range rotations {
					if canPlaceShape(grid, rotated, row, col) {
						placeShape(grid, rotated, row, col, true)
						if tryPlaceShapes(grid, shapesToPlace, allShapes, shapeIndex+1, remainingFreeArea) {
							return true
						}

						placeShape(grid, rotated, row, col, false)
					}
				}
			}
		}
	}

	return false
}

func findFirstEmptyCell(grid [][]bool) (int, int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if !grid[row][col] {
				return row, col
			}
		}
	}
	return -1, -1
}

func getUniqueRotations(shape *PresentShape) []*PresentShape {
	rotations := []*PresentShape{shape}
	seen := make(map[string]struct{})
	seen[shapeSignature(shape)] = struct{}{}

	current := shape
	for i := 0; i < 3; i++ {
		current = rotate90(current)
		sig := shapeSignature(current)
		if _, exists := seen[sig]; !exists {
			rotations = append(rotations, current)
			seen[sig] = struct{}{}
		}
	}

	flipped := flipShape(shape)
	sig := shapeSignature(flipped)
	if _, exists := seen[sig]; !exists {
		rotations = append(rotations, flipped)
		seen[sig] = struct{}{}
	}

	current = flipped
	for i := 0; i < 3; i++ {
		current = rotate90(current)
		signature := shapeSignature(current)
		if _, exists := seen[signature]; !exists {
			rotations = append(rotations, current)
			seen[signature] = struct{}{}
		}
	}

	return rotations
}

func flipShape(shape *PresentShape) *PresentShape {
	newCells := make(map[Coordinate]struct{})
	for coord := range shape.OccupiedCells {
		newCoord := Coordinate{Row: coord.Row, Column: -coord.Column}
		newCells[newCoord] = struct{}{}
	}

	minRow, minCol := 1000000, 1000000
	for coord := range newCells {
		if coord.Row < minRow {
			minRow = coord.Row
		}

		if coord.Column < minCol {
			minCol = coord.Column
		}
	}

	normalized := make(map[Coordinate]struct{})
	for coord := range newCells {
		normalized[Coordinate{
			Row:    coord.Row - minRow,
			Column: coord.Column - minCol,
		}] = struct{}{}
	}

	return &PresentShape{
		Index:         shape.Index,
		OccupiedCells: normalized,
	}
}

func shapeSignature(shape *PresentShape) string {
	var coords []Coordinate
	for coord := range shape.OccupiedCells {
		coords = append(coords, coord)
	}

	slices.SortFunc(coords, func(a, b Coordinate) int {
		if a.Row != b.Row {
			return a.Row - b.Row
		}
		return a.Column - b.Column
	})

	var sig strings.Builder
	for _, coord := range coords {
		sig.WriteString(strconv.Itoa(coord.Row))
		sig.WriteString(",")
		sig.WriteString(strconv.Itoa(coord.Column))
		sig.WriteString(";")
	}

	return sig.String()
}

func rotate90(shape *PresentShape) *PresentShape {
	newCells := make(map[Coordinate]struct{})
	for coord := range shape.OccupiedCells {
		newCoord := Coordinate{Row: coord.Column, Column: -coord.Row}
		newCells[newCoord] = struct{}{}
	}

	minRow, minCol := 1000000, 1000000
	for coord := range newCells {
		if coord.Row < minRow {
			minRow = coord.Row
		}
		if coord.Column < minCol {
			minCol = coord.Column
		}
	}

	normalized := make(map[Coordinate]struct{})
	for coord := range newCells {
		normalized[Coordinate{
			Row:    coord.Row - minRow,
			Column: coord.Column - minCol,
		}] = struct{}{}
	}

	return &PresentShape{
		Index:         shape.Index,
		OccupiedCells: normalized,
	}
}

func canPlaceShape(grid [][]bool, shape *PresentShape, startRow, startCol int) bool {
	for coord := range shape.OccupiedCells {
		row := startRow + coord.Row
		col := startCol + coord.Column

		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return false
		}

		if grid[row][col] {
			return false
		}
	}
	return true
}

func placeShape(grid [][]bool, shape *PresentShape, startRow, startCol int, place bool) {
	for coord := range shape.OccupiedCells {
		row := startRow + coord.Row
		col := startCol + coord.Column
		grid[row][col] = place
	}
}

func parseShape(index int, lines []string) *PresentShape {
	occupiedCells := make(map[Coordinate]struct{})

	for row, line := range lines {
		for col, ch := range line {
			if ch == '#' {
				occupiedCells[Coordinate{Row: row, Column: col}] = struct{}{}
			}
		}
	}

	return &PresentShape{
		Index:         index,
		OccupiedCells: occupiedCells,
	}
}

func parseInput(scanner *bufio.Scanner) (map[int]*PresentShape, []*GiftRegion) {
	presentShapesByIndex := make(map[int]*PresentShape)
	var giftRegions []*GiftRegion

	var currentShapeIndex *int
	var currentShapeLines []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentShapeIndex != nil {
				shape := parseShape(*currentShapeIndex, currentShapeLines)
				presentShapesByIndex[*currentShapeIndex] = shape
				currentShapeIndex = nil
				currentShapeLines = nil
			}
			continue
		}

		if strings.HasSuffix(line, ":") && len(line) <= 3 {
			if currentShapeIndex != nil {
				shape := parseShape(*currentShapeIndex, currentShapeLines)
				presentShapesByIndex[*currentShapeIndex] = shape
			}

			idx, _ := strconv.Atoi(strings.TrimSuffix(line, ":"))
			currentShapeIndex = &idx
			currentShapeLines = nil
			continue
		}

		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			dimParts := strings.Split(parts[0], "x")
			cols, _ := strconv.Atoi(dimParts[0])
			rows, _ := strconv.Atoi(dimParts[1])

			quotas := make(map[int]int)
			counts := strings.Fields(parts[1])
			for i, countStr := range counts {
				count, _ := strconv.Atoi(countStr)
				if count > 0 {
					quotas[i] = count
				}
			}

			giftRegions = append(giftRegions, &GiftRegion{
				Dimensions:             Coordinate{Row: rows, Column: cols},
				GiftShapeQuotasByIndex: quotas,
			})
			continue
		}

		if currentShapeIndex != nil {
			currentShapeLines = append(currentShapeLines, line)
		}
	}

	if currentShapeIndex != nil {
		shape := parseShape(*currentShapeIndex, currentShapeLines)
		presentShapesByIndex[*currentShapeIndex] = shape
	}

	return presentShapesByIndex, giftRegions
}

func Part1() error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	presentShapesByIndex, giftRegions := parseInput(scanner)

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fitCount := 0
	for _, region := range giftRegions {
		if region.canFitAllShapesInQuota(presentShapesByIndex) {
			fitCount++
		}
	}

	fmt.Println("Day 12:", fitCount)

	return nil
}
