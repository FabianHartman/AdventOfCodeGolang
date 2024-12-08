package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputPath8 string = "inputs/day8.txt"

type AntennaMap struct {
	NRows, NCols int
	Antennas     map[string]FrequencyAntennaLocations
}

type FrequencyAntennaLocations []AntennaLocation

type AntennaLocation struct {
	Row, Col int
}

func (al *AntennaLocation) antinodeLocation(to *AntennaLocation) *AntennaLocation {
	rowIncrement := to.Row - al.Row
	colIncrement := to.Col - al.Col

	antinode := AntennaLocation{Row: al.Row + 2*rowIncrement, Col: al.Col + 2*colIncrement}

	return &antinode
}

func (am *AntennaMap) countUniqueAntinodeLocations() int {
	uniqueLocations := map[AntennaLocation]bool{}

	for _, frequency := range am.Antennas {
		for fromI, fromAntenna := range frequency {
			for toI, toAntenna := range frequency {
				if fromI == toI {
					continue
				}

				antinode := fromAntenna.antinodeLocation(&toAntenna)
				if antinode.Col < 0 || antinode.Col >= am.NCols || antinode.Row < 0 || antinode.Row >= am.NCols {
					continue
				}

				uniqueLocations[*antinode] = true
			}
		}
	}

	return len(uniqueLocations)
}

func inputDay8() (*AntennaMap, error) {
	file, err := os.Open(inputPath8)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	antennaMap := AntennaMap{Antennas: make(map[string]FrequencyAntennaLocations)}

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

	antennaMap.NRows = rowI

	return &antennaMap, nil
}

func Day8a() error {
	antennaMap, err := inputDay8()
	if err != nil {
		return err
	}

	fmt.Println("Day 8a:", antennaMap.countUniqueAntinodeLocations())

	return nil
}
