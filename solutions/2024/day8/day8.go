package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputPath string = "inputs/2024/day8.txt"

type Map struct {
	NRows, NCols int
	Antennas     map[string]FrequencyAntennaLocations
}

type FrequencyAntennaLocations []AntennaLocation

type AntennaLocation struct {
	Row, Col int
}

func (al *AntennaLocation) antinodeLocationPart1(to *AntennaLocation) *AntennaLocation {
	rowIncrement := to.Row - al.Row
	colIncrement := to.Col - al.Col

	antinode := AntennaLocation{Row: al.Row + 2*rowIncrement, Col: al.Col + 2*colIncrement}

	return &antinode
}

func (al *AntennaLocation) antinodeLocationPart2(to *AntennaLocation, NRows int, NCols int) []AntennaLocation {
	rowIncrement := to.Row - al.Row
	colIncrement := to.Col - al.Col

	row := to.Row
	col := to.Col

	antinodes := []AntennaLocation{}
	antinodes = append(antinodes, AntennaLocation{Row: row, Col: col})

	for {
		row += rowIncrement
		col += colIncrement

		if !(row < 0 || row >= NRows || col < 0 || col >= NCols) {
			antinodes = append(antinodes, AntennaLocation{Row: row, Col: col})
		} else {
			break
		}

	}

	return antinodes
}

func (am *Map) countUniqueAntinodeLocationsPart1() int {
	uniqueLocations := map[AntennaLocation]bool{}

	for _, frequency := range am.Antennas {
		for fromI, fromAntenna := range frequency {
			for toI, toAntenna := range frequency {
				if fromI == toI {
					continue
				}

				antinode := fromAntenna.antinodeLocationPart1(&toAntenna)
				if antinode.Col < 0 || antinode.Col >= am.NCols || antinode.Row < 0 || antinode.Row >= am.NCols {
					continue
				}

				uniqueLocations[*antinode] = true
			}
		}
	}

	return len(uniqueLocations)
}

func (am *Map) countUniqueAntinodeLocationsPart2() int {
	uniqueLocations := map[AntennaLocation]bool{}

	for _, frequency := range am.Antennas {
		for fromI, fromAntenna := range frequency {
			for toI, toAntenna := range frequency {
				if fromI == toI {
					continue
				}

				antinodes := fromAntenna.antinodeLocationPart2(&toAntenna, am.NRows, am.NCols)
				for _, antinode := range antinodes {
					if !(antinode.Col < 0 || antinode.Col >= am.NCols || antinode.Row < 0 || antinode.Row >= am.NCols) {
						uniqueLocations[antinode] = true
					}
				}
			}
		}
	}

	return len(uniqueLocations)
}

func input() (*Map, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	antennaMap := Map{Antennas: make(map[string]FrequencyAntennaLocations)}

	rowI := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowI++
		rowParts := strings.Split(scanner.Text(), "")

		if rowI == 0 {
			antennaMap.NCols = len(rowParts)
		}

		for colI, rowPart := range rowParts {
			if rowPart == "." {
				continue
			}

			antennaMap.Antennas[rowPart] = append(antennaMap.Antennas[rowPart], AntennaLocation{
				Row: rowI,
				Col: colI,
			})
		}
	}

	antennaMap.NRows = rowI + 1

	return &antennaMap, nil
}

func Part1() error {
	antennaMap, err := input()
	if err != nil {
		return err
	}

	fmt.Println("Day 8a:", antennaMap.countUniqueAntinodeLocationsPart1())

	return nil
}

func Part2() error {
	antennaMap, err := input()
	if err != nil {
		return err
	}

	fmt.Println("Day 8b:", antennaMap.countUniqueAntinodeLocationsPart2())

	return nil
}
