package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inputPath = "inputs/2024/day10.txt"

type Trail struct {
	Map   [][]int64
	Heads [][]int
}

type VisitedState struct {
	Map          [][]bool
	HeadsVisited [][]bool
}

type Traversal struct {
	TrailMap     [][]int64
	VisitedState *VisitedState
}

func input() ([]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func solvePart1(lines []string) int {
	solution := 0
	trailData := parseTrailData(lines)

	for _, trailHead := range trailData.Heads {
		visitedState := NewVisitedState(len(trailData.Map))
		traversal := NewTraversal(trailData.Map, visitedState)

		traversal.traversePartOne(trailHead[0], trailHead[1])

		solution += visitedState.countVisitedHeads()
	}

	return solution
}

func solvePart2(lines []string) int {
	solution := 0
	trailData := parseTrailData(lines)

	for _, trailHead := range trailData.Heads {
		rank := 0

		traversal := NewTraversal(trailData.Map, nil)
		traversal.traversePartTwo(trailHead[0], trailHead[1], &rank)

		solution += rank
	}

	return solution
}

func parseTrailData(lines []string) *Trail {
	var trailMap [][]int64
	var trailHeads [][]int

	for i, line := range lines {
		if len(line) > 0 {
			row := make([]int64, len(line))
			for j, char := range line {
				if char == '.' {
					row[j] = 100
				} else {
					value, _ := strconv.ParseInt(string(char), 32, 10)
					row[j] = value

					if value == 0 {
						trailHeads = append(trailHeads, []int{i, j})
					}
				}
			}
			trailMap = append(trailMap, row)
		}
	}

	return &Trail{
		Map:   trailMap,
		Heads: trailHeads,
	}
}

func NewVisitedState(rows int) *VisitedState {
	visited := make([][]bool, rows)
	visitedHeads := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, rows)
		visitedHeads[i] = make([]bool, rows)
	}

	return &VisitedState{
		Map:          visited,
		HeadsVisited: visitedHeads,
	}
}

func NewTraversal(trailMap [][]int64, visitedState *VisitedState) *Traversal {
	if visitedState == nil {
		visitedState = NewVisitedState(len(trailMap))
	}

	return &Traversal{
		TrailMap:     trailMap,
		VisitedState: visitedState,
	}
}

func (t *Traversal) traversePartOne(row, col int) {
	if row < 0 || col < 0 || row >= len(t.TrailMap) || col >= len(t.TrailMap[0]) || t.VisitedState.Map[row][col] {
		return
	}

	if t.TrailMap[row][col] == 9 {
		t.VisitedState.HeadsVisited[row][col] = true
		return
	}

	t.VisitedState.Map[row][col] = true

	if row+1 < len(t.TrailMap) && t.TrailMap[row+1][col]-t.TrailMap[row][col] == 1 {
		t.traversePartOne(row+1, col)
	}

	if row-1 >= 0 && t.TrailMap[row-1][col]-t.TrailMap[row][col] == 1 {
		t.traversePartOne(row-1, col)
	}

	if col+1 < len(t.TrailMap[0]) && t.TrailMap[row][col+1]-t.TrailMap[row][col] == 1 {
		t.traversePartOne(row, col+1)
	}

	if col-1 >= 0 && t.TrailMap[row][col-1]-t.TrailMap[row][col] == 1 {
		t.traversePartOne(row, col-1)
	}
}

func (t *Traversal) traversePartTwo(row, col int, rank *int) {
	if row < 0 || col < 0 || row >= len(t.TrailMap) || col >= len(t.TrailMap[0]) {
		return
	}

	if t.TrailMap[row][col] == 9 {
		(*rank)++
	}

	if row+1 < len(t.TrailMap) && t.TrailMap[row+1][col]-t.TrailMap[row][col] == 1 {
		t.traversePartTwo(row+1, col, rank)
	}

	if row-1 >= 0 && t.TrailMap[row-1][col]-t.TrailMap[row][col] == 1 {
		t.traversePartTwo(row-1, col, rank)
	}

	if col+1 < len(t.TrailMap[0]) && t.TrailMap[row][col+1]-t.TrailMap[row][col] == 1 {
		t.traversePartTwo(row, col+1, rank)
	}

	if col-1 >= 0 && t.TrailMap[row][col-1]-t.TrailMap[row][col] == 1 {
		t.traversePartTwo(row, col-1, rank)
	}
}

func (v *VisitedState) countVisitedHeads() int {
	count := 0

	for _, row := range v.HeadsVisited {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}

	return count
}

func Part1() error {
	lines, err := input()
	if err != nil {
		return err
	}

	fmt.Println("day10a:", solvePart1(lines))

	return nil
}

func Part2() error {
	lines, err := input()
	if err != nil {
		return err
	}

	fmt.Println("day10b:", solvePart2(lines))

	return nil
}
