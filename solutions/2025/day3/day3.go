package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputPath = "inputs/2025/day3.txt"

func Part1() error {
	var result int

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

	for scanner.Scan() {
		inputRow := scanner.Text()

		highest := 0

		for startI := 0; startI < len(inputRow); startI++ {
			for endI := startI + 1; endI < len(inputRow); endI++ {
				value, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(inputRow[startI]), string(inputRow[endI])))
				if value > highest {
					highest = value
				}
			}
		}

		result += highest
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 3a:", result)

	return nil
}

func Part2() error {
	var result int

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

	for scanner.Scan() {
		inputRow := scanner.Text()

		highest := 0

		for index1 := 0; index1 < len(inputRow); index1++ {
			for index2 := index1 + 1; index2 < len(inputRow); index2++ {
				for index3 := index2 + 1; index3 < len(inputRow); index3++ {
					for index4 := index3 + 1; index4 < len(inputRow); index4++ {
						for index5 := index4 + 1; index5 < len(inputRow); index5++ {
							for index6 := index5 + 1; index6 < len(inputRow); index6++ {
								for index7 := index6 + 1; index7 < len(inputRow); index7++ {
									for index8 := index7 + 1; index8 < len(inputRow); index8++ {
										for index9 := index8 + 1; index9 < len(inputRow); index9++ {
											for index10 := index9 + 1; index10 < len(inputRow); index10++ {
												for index11 := index10 + 1; index11 < len(inputRow); index11++ {
													for index12 := index11 + 1; index12 < len(inputRow); index12++ {
														value, _ := strconv.Atoi(fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s", string(inputRow[index1]), string(inputRow[index2]), string(inputRow[index3]), string(inputRow[index4]), string(inputRow[index5]), string(inputRow[index6]), string(inputRow[index7]), string(inputRow[index8]), string(inputRow[index9]), string(inputRow[index10]), string(inputRow[index11]), string(inputRow[index12])))
														if value > highest {
															highest = value
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		result += highest
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("error reading the file: %s", err)
	}

	fmt.Println("Day 3b:", result)

	return nil
}
