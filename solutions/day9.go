package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputPath9 string = "inputs/day9.txt"

type Block struct {
	Length int
	ID     int
}

func inputDay9() (map[int]*Block, error) {
	file, err := os.Open(inputPath9)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	blockMap := map[int]*Block{}
	diskI := -1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "")
		id := 0
		for i, number := range numbers {
			intNumber, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}

			if i%2 == 0 {
				blockMap[diskI] = &Block{Length: intNumber, ID: id}
				diskI += intNumber
				id++
			} else {
				diskI += intNumber
			}

		}
	}

	return blockMap, nil
}

func createIdList(blockMap map[int]*Block) ([]int, int) {
	startingPositions := []int{}
	length := 0

	for k := range blockMap {
		startingPositions = append(startingPositions, k)
	}

	sort.Ints(startingPositions)

	idList := []int{}

	for _, startingPosition := range startingPositions {
		for len(idList) <= startingPosition {
			idList = append(idList, -1)
		}

		for i := 0; i < blockMap[startingPosition].Length; i++ {
			idList = append(idList, blockMap[startingPosition].ID)
			length++
		}
	}

	return idList, length
}

func generateChecksumPart1(idList []int, length int) []int {
	result := []int{}
	i := -1
	reverseI := 1

	for {
		i++
		if idList[i] != -1 {
			result = append(result, idList[i])
		} else {
			newValue := -1
			for newValue == -1 {
				newValue = idList[len(idList)-reverseI]
				reverseI++

			}

			result = append(result, newValue)
		}

		if i+1 >= length {
			break
		}
	}

	return result
}

func generateChecksumPart2(idList []int, length int) []int {
	result := make([]int, len(idList))
	copy(result, idList)
	index := len(idList) - 1

	for index > 0 {
		if result[index] != -1 {
			number := result[index]

			block := []int{}
			block = append(block, index)

			for {
				index--
				if index <= 0 {
					break
				}

				if number != result[index] {
					break
				}

				block = append(block, index)
			}

			freePlaces := []int{}
			for i := 0; i <= block[len(block)-1]; i++ {
				if result[i] == -1 {
					freePlaces = append(freePlaces, i)
					if len(freePlaces) == len(block) {
						for _, place := range freePlaces {
							result[place] = number
						}

						for _, previousPlace := range block {
							result[previousPlace] = -1
						}

						break
					}
				} else {
					freePlaces = []int{}
				}
			}

		} else {
			index--
		}
	}

	return result
}

func sumUpChecksum(checksum []int) int {
	total := 0
	for i, v := range checksum {
		if v != -1 {
			total += i * v
		}
	}

	return total
}

func Day9a() error {
	blockMap, err := inputDay9()
	if err != nil {
		return err
	}

	fmt.Println("day9a:", sumUpChecksum(generateChecksumPart1(createIdList(blockMap))))

	return nil
}

func Day9b() error {
	blockMap, err := inputDay9()
	if err != nil {
		return err
	}

	fmt.Println("day9b:", sumUpChecksum(generateChecksumPart2(createIdList(blockMap))))

	return nil
}
