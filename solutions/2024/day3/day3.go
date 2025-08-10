package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputPath = "inputs/2024/day3.txt"

type Status struct {
	MulProgress          string
	FirstNumber          string
	SecondNumber         string
	FirstNumberCompleted bool
	MulOpened            bool
}

func (s *Status) Reset() {
	s.MulOpened = false
	s.MulProgress = ""
	s.FirstNumberCompleted = false
	s.FirstNumber = ""
	s.SecondNumber = ""
}

func (s *Status) String() string {
	return fmt.Sprintf("MulProgress: %s, FirstNumber: %s, SecondNumber: %s, FirstNumberCompleted: %v, MulOpened: %v",
		s.MulProgress, s.FirstNumber, s.SecondNumber, s.FirstNumberCompleted, s.MulOpened)
}

func containsString(slice []string, string string) bool {
	for _, s := range slice {
		if s == string {
			return true
		}
	}

	return false
}

func input() ([]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rows := []string{}

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows, nil
}

func Part1() error {
	input, err := input()
	if err != nil {
		return err
	}

	var total int64

	for _, row := range input {
		status := Status{
			MulProgress:          "",
			FirstNumber:          "",
			SecondNumber:         "",
			FirstNumberCompleted: false,
			MulOpened:            false,
		}

		numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
		mulString := []string{"m", "u", "l"}

		for _, char := range row {
			if containsString(numbers, string(char)) {
				if !status.FirstNumberCompleted && status.MulOpened {
					if status.MulProgress != "mul" {
						status.Reset()
					} else {
						status.FirstNumber += string(char)
					}
				} else if status.MulOpened {
					status.SecondNumber += string(char)
				} else {
					status.Reset()
				}
			} else if char == '(' {
				if status.MulOpened {
					status.Reset()
				} else {
					if status.MulProgress == "mul" {
						status.MulOpened = true
					} else {
						status.Reset()
					}
				}
			} else if char == ',' {
				if 1 <= len(status.FirstNumber) && len(status.FirstNumber) <= 3 {
					status.FirstNumberCompleted = true
				} else {
					status.Reset()
				}
			} else if char == ')' {
				if status.MulOpened {
					if 1 <= len(status.FirstNumber) && len(status.FirstNumber) <= 3 && 1 <= len(status.SecondNumber) && len(status.SecondNumber) <= 3 {
						first, _ := strconv.ParseInt(status.FirstNumber, 10, 64)
						second, _ := strconv.ParseInt(status.SecondNumber, 10, 64)

						total += first * second

						status.Reset()
					}
				}
			} else if containsString(mulString, string(char)) {
				if !containsString(strings.Split(status.MulProgress, ""), string(char)) {
					status.MulProgress += string(char)
				} else {
					status.Reset()
					if char == 'm' {
						status.MulProgress += string(char)
					}
				}
			} else {
				status.Reset()
			}
		}
		status.Reset()
	}

	fmt.Println("Day 3a:", total)

	return nil
}

func Part2() error {
	input, err := input()
	if err != nil {
		return err
	}

	var total int64

	lastDo := -1
	lastDont := -2

	for _, row := range input {
		status := Status{
			MulProgress:          "",
			FirstNumber:          "",
			SecondNumber:         "",
			FirstNumberCompleted: false,
			MulOpened:            false,
		}

		numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
		mulString := []string{"m", "u", "l"}

		for i, char := range row {
			if containsString(numbers, string(char)) {
				if !status.FirstNumberCompleted && status.MulOpened {
					if status.MulProgress != "mul" {
						status.Reset()
					} else {
						status.FirstNumber += string(char)
					}
				} else if status.MulOpened {
					status.SecondNumber += string(char)
				} else {
					status.Reset()
				}
			} else if char == '(' {
				if status.MulOpened {
					status.Reset()
				} else {
					if status.MulProgress == "mul" {
						status.MulOpened = true
					} else {
						status.Reset()
					}
				}
			} else if char == ',' {
				if 1 <= len(status.FirstNumber) && len(status.FirstNumber) <= 3 {
					status.FirstNumberCompleted = true
				} else {
					status.Reset()
				}
			} else if char == ')' {
				if status.MulOpened {
					if 1 <= len(status.FirstNumber) && len(status.FirstNumber) <= 3 && 1 <= len(status.SecondNumber) && len(status.SecondNumber) <= 3 {
						first, _ := strconv.ParseInt(status.FirstNumber, 10, 64)
						second, _ := strconv.ParseInt(status.SecondNumber, 10, 64)

						doPattern := regexp.MustCompile(`do\(\)`)
						dontPattern := regexp.MustCompile(`don't\(\)`)

						beforeInput := row[:i]

						dos := []int{}
						for _, i := range doPattern.FindAllStringIndex(beforeInput, len(beforeInput)-1) {
							dos = append(dos, i[1])
						}

						donts := []int{}
						for _, i := range dontPattern.FindAllStringIndex(beforeInput, len(beforeInput)-1) {
							donts = append(donts, i[1])
						}

						do := -1
						dont := -2

						if len(dos) > 0 {
							do = dos[len(dos)-1]
						}
						if len(donts) > 0 {
							dont = donts[len(donts)-1]
						}

						if do == -1 && dont == -2 {
							do = lastDo
							dont = lastDont
						}

						lastDo = do
						lastDont = dont

						if do > dont {
							total += first * second
						}

						status.Reset()
					}
				}
			} else if containsString(mulString, string(char)) {
				if !containsString(strings.Split(status.MulProgress, ""), string(char)) {
					status.MulProgress += string(char)
				} else {
					status.Reset()
					if char == 'm' {
						status.MulProgress += string(char)
					}
				}
			} else {
				status.Reset()
			}
		}
	}

	fmt.Println("Day 3b:", total)

	return nil
}
