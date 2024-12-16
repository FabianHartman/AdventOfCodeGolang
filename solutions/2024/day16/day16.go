package day16

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

var inputPath string = "inputs/2024/day16.txt"

type Position struct {
	Col, Row int
}

type Direction struct {
	DeltaCol, DeltaRow int
}

type Head struct {
	Pos       *Position
	Direction int
	Score     int
	Path      []Position
}

type Maze struct {
	Grid  [][]string
	Start *Position
	End   *Position
}

var directions = []Direction{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type PriorityQueue []Head

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Score < pq[j].Score
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Head))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]

	return item
}

func input() (*Maze, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	grid := [][]string{}
	var start, end Position
	row := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
		for col, char := range grid[row] {
			if char == "S" {
				start = Position{col, row}
			} else if char == "E" {
				end = Position{col, row}
			}
		}
		row++
	}

	return &Maze{Grid: grid, Start: &start, End: &end}, nil
}

func (this *Position) newPosition(direction *Direction) *Position {
	return &Position{this.Col + direction.DeltaCol, this.Row + direction.DeltaRow}
}

func (this *Position) generateMapKey(direction int) string {
	return fmt.Sprintf("%d,%d,%d", this.Col, this.Row, direction)
}

func (this *Position) generateMapKeyWithoutDirection() string {
	return fmt.Sprintf("%d,%d", this.Col, this.Row)
}

func (this *Maze) isInGridAndNotInWall(position *Position) bool {
	return position.Row >= 0 && position.Row < len(this.Grid) && position.Col >= 0 && position.Col < len(this.Grid[0]) && this.Grid[position.Row][position.Col] != "#"
}

func (this *Maze) isFinished(p *Position) bool {
	return *p == *this.End
}

func (this *Maze) calculateBestRoute() (int, error) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	initialHead := Head{Pos: &Position{Col: this.Start.Col, Row: this.Start.Row}, Direction: 1, Score: 0}
	heap.Push(&pq, initialHead)

	visited := map[string]bool{}

	for pq.Len() > 0 {
		currentHead := heap.Pop(&pq).(Head)

		if this.isFinished(currentHead.Pos) {
			return currentHead.Score, nil
		}

		key := currentHead.Pos.generateMapKey(currentHead.Direction)
		if visited[key] {
			continue
		}

		visited[key] = true

		nextPos := currentHead.Pos.newPosition(&directions[currentHead.Direction])
		if this.isInGridAndNotInWall(nextPos) {
			heap.Push(&pq, Head{
				Pos:       nextPos,
				Direction: currentHead.Direction,
				Score:     currentHead.Score + 1,
			})
		}

		heap.Push(&pq, Head{Pos: currentHead.Pos, Direction: (currentHead.Direction + 1) % 4, Score: currentHead.Score + 1000})
		heap.Push(&pq, Head{Pos: currentHead.Pos, Direction: (currentHead.Direction + 3) % 4, Score: currentHead.Score + 1000})
	}

	return 0, fmt.Errorf("no route found")
}

func (this *Maze) countOptimalPathTiles(score int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	initialHead := Head{
		Pos:       &Position{Col: this.Start.Col, Row: this.Start.Row},
		Direction: 1,
		Score:     0,
		Path:      []Position{{Col: this.Start.Col, Row: this.Start.Row}},
	}

	heap.Push(&pq, initialHead)

	visited := map[string]int{}
	paths := [][]Position{}

	for pq.Len() > 0 {
		currentHead := heap.Pop(&pq).(Head)

		if currentHead.Score > score {
			continue
		}

		key := currentHead.Pos.generateMapKey(currentHead.Direction)
		if previousScore, exists := visited[key]; exists && previousScore < currentHead.Score {
			continue
		}

		visited[key] = currentHead.Score

		if this.isFinished(currentHead.Pos) {
			paths = append(paths, currentHead.Path)
			continue
		}

		nextPos := currentHead.Pos.newPosition(&directions[currentHead.Direction])
		if this.isInGridAndNotInWall(nextPos) {
			newPath := append([]Position{}, currentHead.Path...)
			newPath = append(newPath, *nextPos)

			heap.Push(&pq, Head{
				Pos:       nextPos,
				Direction: currentHead.Direction,
				Score:     currentHead.Score + 1,
				Path:      newPath,
			})
		}

		heap.Push(&pq, Head{Pos: currentHead.Pos, Direction: (currentHead.Direction + 1) % 4, Score: currentHead.Score + 1000, Path: append([]Position{}, currentHead.Path...)})
		heap.Push(&pq, Head{Pos: currentHead.Pos, Direction: (currentHead.Direction + 3) % 4, Score: currentHead.Score + 1000, Path: append([]Position{}, currentHead.Path...)})
	}

	uniqueTiles := map[string]bool{}
	for _, path := range paths {
		for _, tile := range path {
			uniqueTiles[tile.generateMapKeyWithoutDirection()] = true
		}
	}

	return len(uniqueTiles)
}

func Part1() error {
	maze, err := input()
	if err != nil {
		return err
	}

	lowestScore, err := maze.calculateBestRoute()
	if err != nil {
		return err
	}

	fmt.Println("Day 16a:", lowestScore)

	return nil
}

func Part2() error {
	maze, err := input()
	if err != nil {
		return err
	}

	lowestScore, err := maze.calculateBestRoute()
	if err != nil {
		return err
	}

	fmt.Println("Day 16b:", maze.countOptimalPathTiles(lowestScore))

	return nil
}
