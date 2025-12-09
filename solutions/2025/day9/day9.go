package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const inputPath = "inputs/2025/day9.txt"

type Polygon struct {
	Tiles []Tile
}

type Tile struct {
	X, Y int
}

func (this *Polygon) Contains(tile *Tile) bool {
	edgeCrossCount := 0
	n := len(this.Tiles)

	for i := 0; i < n; i++ {
		start := this.Tiles[i]
		end := this.Tiles[(i+1)%n]

		if (start.Y == end.Y && tile.Y == start.Y && valueBetween(tile.X, start.X, end.X)) ||
			(start.X == end.X && tile.X == start.X && valueBetween(tile.Y, start.Y, end.Y)) {
			return true
		}

		if (start.Y > tile.Y) != (end.Y > tile.Y) {
			if float64(tile.X) < float64(end.X-start.X)/float64(end.Y-start.Y)*float64(tile.Y-start.Y)+float64(start.X) {
				edgeCrossCount++
			}
		}
	}

	return edgeCrossCount%2 == 1
}

func (this *Tile) Area(that *Tile) int {
	if this.X == that.X {
		return absValue(this.Y - that.Y)
	}

	if this.Y == that.Y {
		return absValue(this.X - that.X)
	}

	return (1 + absValue(this.X-that.X)) * (1 + absValue(this.Y-that.Y))
}

type Link struct {
	t1, t2 *Tile
	area   int
}

func NewLink(t1, t2 *Tile) Link {
	return Link{t1: t1, t2: t2, area: t1.Area(t2)}
}

func absValue(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func valueBetween(val, a, b int) bool {
	return val >= minValue(a, b) && val <= maxValue(a, b)
}

func polygonEdgeCrossesRectangle(rectStart, rectEnd, segStart, segEnd *Tile) bool {
	minX, maxX := minValue(rectStart.X, rectEnd.X), maxValue(rectStart.X, rectEnd.X)
	minY, maxY := minValue(rectStart.Y, rectEnd.Y), maxValue(rectStart.Y, rectEnd.Y)

	if (segStart.X < minX && segEnd.X < minX) || (segStart.X > maxX && segEnd.X > maxX) ||
		(segStart.Y < minY && segEnd.Y < minY) || (segStart.Y > maxY && segEnd.Y > maxY) {
		return false
	}

	if onBoundary(segStart, minX, maxX, minY, maxY) && onBoundary(segEnd, minX, maxX, minY, maxY) {
		if (segStart.X == segEnd.X && (segStart.X == minX || segStart.X == maxX)) ||
			(segStart.Y == segEnd.Y && (segStart.Y == minY || segStart.Y == maxY)) {
			return false
		}
	}

	rectEdges := [][2]*Tile{
		{{X: minX, Y: minY}, {X: maxX, Y: minY}},
		{{X: maxX, Y: minY}, {X: maxX, Y: maxY}},
		{{X: maxX, Y: maxY}, {X: minX, Y: maxY}},
		{{X: minX, Y: maxY}, {X: minX, Y: minY}},
	}

	for _, edge := range rectEdges {
		if crosses(segStart, segEnd, edge[0], edge[1]) {
			return true
		}
	}

	return false
}

func onBoundary(p *Tile, minX, maxX, minY, maxY int) bool {
	return (p.X == minX || p.X == maxX) && valueBetween(p.Y, minY, maxY) ||
		(p.Y == minY || p.Y == maxY) && valueBetween(p.X, minX, maxX)
}

func crosses(a1, a2, b1, b2 *Tile) bool {
	d1 := orientation(b1, b2, a1)
	d2 := orientation(b1, b2, a2)
	d3 := orientation(a1, a2, b1)
	d4 := orientation(a1, a2, b2)

	return ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0))
}

func orientation(t1, t2, t3 *Tile) int {
	val := (t2.X-t1.X)*(t3.Y-t1.Y) - (t2.Y-t1.Y)*(t3.X-t1.X)
	if val > 0 {
		return 1
	}

	if val < 0 {
		return -1
	}

	return 0
}

func All[T any](input []T, predicate func(T) bool) bool {
	for _, v := range input {
		if !predicate(v) {
			return false
		}
	}

	return true
}

func parseTiles() ([]Tile, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var tiles []Tile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing X: %w", err)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing Y: %w", err)
		}

		tiles = append(tiles, Tile{X: x, Y: y})
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return tiles, nil
}

func Part1() error {
	tiles, err := parseTiles()
	if err != nil {
		return err
	}

	maxArea := 0
	for i, t1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := t1.Area(&tiles[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Println("Day 9a:", maxArea)

	return nil
}

func Part2() error {
	tiles, err := parseTiles()
	if err != nil {
		return err
	}

	polygon := Polygon{Tiles: tiles}

	var links []Link
	for i, t1 := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			links = append(links, NewLink(&t1, &tiles[j]))
		}
	}

	slices.SortFunc(links, func(a, b Link) int { return b.area - a.area })

	for _, link := range links {
		t1, t2 := link.t1, link.t2

		if t1.X == t2.X || t1.Y == t2.Y {
			continue
		}

		corners := []*Tile{
			{X: t1.X, Y: t1.Y},
			{X: t1.X, Y: t2.Y},
			{X: t2.X, Y: t1.Y},
			{X: t2.X, Y: t2.Y},
		}

		if !All(corners, polygon.Contains) {
			continue
		}

		valid := true
		for i := 0; i < len(polygon.Tiles); i++ {
			next := (i + 1) % len(polygon.Tiles)
			if polygonEdgeCrossesRectangle(t1, t2, &polygon.Tiles[i], &polygon.Tiles[next]) {
				valid = false

				break
			}
		}

		if valid {
			fmt.Println("Day 9b:", link.area)

			break
		}
	}

	return nil
}
