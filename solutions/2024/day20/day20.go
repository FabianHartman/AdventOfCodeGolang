package day20

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var inputPath = "inputs/2024/day20.txt"

type Position struct {
	Row, Col int
}

type RaceTrack struct {
	Grid  [][]string
	Start *Position
	End   *Position
}

func (this *RaceTrack) findPosition(char string) (*Position, error) {
	for rowI, row := range this.Grid {
		for colI, gridCell := range row {
			if gridCell == char {
				return &Position{rowI, colI}, nil
			}
		}
	}

	return nil, fmt.Errorf("position not found")
}

func input() (*RaceTrack, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err)
	}

	defer file.Close()

	raceTrack := RaceTrack{Grid: [][]string{}}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		raceTrack.Grid = append(raceTrack.Grid, strings.Split(scanner.Text(), ""))
	}

	raceTrack.Start, err = raceTrack.findPosition("S")
	if err != nil {
		return nil, err
	}

	raceTrack.End, err = raceTrack.findPosition("E")
	if err != nil {
		return nil, err
	}

	return &raceTrack, nil
}

func isValidPosition(c string) bool {
	return c == "S" || c == "E" || c == "."
}

func isInTrack(track map[Position]int, pos *Position) bool {
	_, exists := track[*pos]
	return exists
}

func (this *RaceTrack) picoSecondsCheats() int {
	track := make(map[Position]int)
	track[*this.Start] = 0
	currentPosition := *this.Start
	stepI := 0

	for currentPosition != *this.End {
		stepI++
		found := false

		for _, dir := range []Position{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
			newPosition := Position{Row: currentPosition.Row + dir.Row, Col: currentPosition.Col + dir.Col}

			if !isInTrack(track, &newPosition) && isValidPosition(this.Grid[newPosition.Row][newPosition.Col]) {
				currentPosition = newPosition
				track[currentPosition] = stepI
				found = true

				break
			}
		}

		if !found {
			break
		}
	}

	count := 0

	for trackPosition := range track {
		for _, dir := range []Position{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
			newRow, newCol := trackPosition.Row+dir.Row, trackPosition.Col+dir.Col

			if isInTrack(track, &Position{newRow, newCol}) {
				continue
			}

			otherPos := Position{newRow + dir.Row, newCol + dir.Col}

			if isInTrack(track, &otherPos) && track[otherPos]-track[trackPosition] >= 102 {
				count++
			}
		}
	}

	return count
}

func (this *Position) cheatEndpoints(track map[Position]int) map[Position]bool {
	output := make(map[Position]bool)
	for deltaRow := -20; deltaRow <= 20; deltaRow++ {
		maxDeltaCol := 20 - int(math.Abs(float64(deltaRow)))

		for deltaCol := -maxDeltaCol; deltaCol <= maxDeltaCol; deltaCol++ {
			newPos := Position{Row: this.Row + deltaRow, Col: this.Col + deltaCol}

			if _, exists := track[newPos]; exists {
				output[newPos] = true
			}
		}
	}

	return output
}

func (this *Position) manhattanDistance(position2 *Position) int {
	return int(math.Abs(float64(this.Row-position2.Row))) + int(math.Abs(float64(this.Col-position2.Col)))
}

func (this *RaceTrack) picoSecondsBigCheats() int {
	track := make(map[Position]int)
	track[*this.Start] = 0
	currentPosition := *this.Start
	stepI := 0

	for currentPosition != *this.End {
		stepI++
		found := false

		for _, direction := range [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
			newPosition := Position{Row: currentPosition.Row + direction[0], Col: currentPosition.Col + direction[1]}

			if !isInTrack(track, &newPosition) && isValidPosition(this.Grid[newPosition.Row][newPosition.Col]) {
				currentPosition = newPosition
				track[currentPosition] = stepI
				found = true

				break
			}
		}

		if !found {
			break
		}
	}

	count := 0
	for trackPosition := range track {
		potentialCheatEndpoints := trackPosition.cheatEndpoints(track)

		for potentialCheatEndpoint := range potentialCheatEndpoints {
			if track[potentialCheatEndpoint]-track[trackPosition]-trackPosition.manhattanDistance(&potentialCheatEndpoint) >= 100 {
				count++
			}
		}
	}

	return count
}

func Part1() error {
	raceTrack, err := input()
	if err != nil {
		return err
	}

	goodAvailableCheatOptions := raceTrack.picoSecondsCheats()

	fmt.Println("Day 20a:", goodAvailableCheatOptions)

	return nil
}

func Part2() error {
	raceTrack, err := input()
	if err != nil {
		return err
	}

	goodAvailableCheatOptions := raceTrack.picoSecondsBigCheats()

	fmt.Println("Day 20b:", goodAvailableCheatOptions)

	return nil
}
