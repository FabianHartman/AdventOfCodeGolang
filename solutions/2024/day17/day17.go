package day17

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var inputPath string = "inputs/2024/day17.txt"

type Registers struct {
	A, B, C int
}

type Instruction struct {
	Opcode, Operand int
}

func input() (*Registers, []Instruction, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	registers := Registers{}
	_, err = fmt.Sscanf(scanner.Text(), "Register A: %d", &registers.A)
	if err != nil {
		return nil, nil, err
	}

	scanner.Scan()
	_, err = fmt.Sscanf(scanner.Text(), "Register B: %d", &registers.B)
	if err != nil {
		return nil, nil, err
	}

	scanner.Scan()
	_, err = fmt.Sscanf(scanner.Text(), "Register C: %d", &registers.C)
	if err != nil {
		return nil, nil, err
	}

	scanner.Scan()
	scanner.Scan()

	var instructions []Instruction
	strInstructions := strings.Split(scanner.Text()[9:], ",")
	for idx := 0; idx < len(strInstructions); idx += 2 {
		opcode, _ := strconv.Atoi(strInstructions[idx])
		operand, _ := strconv.Atoi(strInstructions[idx+1])

		instructions = append(instructions, Instruction{
			Opcode:  opcode,
			Operand: operand,
		})
	}

	return &registers, instructions, nil
}

func run(instruction Instruction, registers *Registers, instructionsPointer *int) string {
	var combo int

	switch instruction.Operand {
	case 4:
		combo = registers.A
	case 5:
		combo = registers.B
	case 6:
		combo = registers.C
	default:
		combo = instruction.Operand
	}

	var output string
	switch instruction.Opcode {
	case 0:
		registers.A = registers.A / int(math.Pow(2, float64(combo)))
		*instructionsPointer++
	case 1:
		registers.B = registers.B ^ instruction.Operand
		*instructionsPointer++
	case 2:
		registers.B = combo % 8
		*instructionsPointer++
	case 3:
		if registers.A != 0 {
			*instructionsPointer = instruction.Operand / 2
		} else {
			*instructionsPointer++
		}
	case 4:
		registers.B = registers.B ^ registers.C
		*instructionsPointer++
	case 5:
		output = strconv.Itoa(combo % 8)
		*instructionsPointer++
	case 6:
		registers.B = registers.A / int(math.Pow(2, float64(combo)))
		*instructionsPointer++
	case 7:
		registers.C = registers.A / int(math.Pow(2, float64(combo)))
		*instructionsPointer++
	}

	return output
}

func runAndGetOutput(registers *Registers, instructions []Instruction) (string, error) {
	pc := 0
	outs := []string{}

	for pc < len(instructions) {
		out := run(instructions[pc], registers, &pc)
		if out != "" {
			outs = append(outs, out)
		}
	}

	return strings.Join(outs, ","), nil
}

func backPropagation(instructions []Instruction, position, initVal int) (int, error) {
	intIns := make([]int, len(instructions)*2)
	for idx, instruction := range instructions {
		intIns[idx*2] = instruction.Opcode
		intIns[idx*2+1] = instruction.Operand
	}

	for idx := 0; idx < 8; idx++ {
		registers := Registers{
			A: initVal*8 + idx,
			B: 0,
			C: 0,
		}

		pc := 0
		var output []int

		for pc < len(instructions) {
			out := run(instructions[pc], &registers, &pc)
			if out != "" {
				n, err := strconv.Atoi(out)
				if err != nil {
					return 0, err
				}

				output = append(output, n)
			}
		}

		ok := true

		for j := position; j < len(intIns); j++ {
			if intIns[j] != output[j-position] {
				ok = false
				break
			}
		}

		if ok {
			if position == 0 {
				return initVal*8 + idx, nil
			}

			val, err := backPropagation(instructions, position-1, initVal*8+idx)
			if err != nil {
				return 0, err
			}

			if val != -1 {
				return val, nil
			}
		}
	}
	return -1, nil
}

func findLowestAValue(instructions []Instruction) (int, error) {
	return backPropagation(instructions, len(instructions)*2-1, 0)
}

func Part1() error {
	registers, instructions, err := input()
	if err != nil {
		return err
	}

	output, err := runAndGetOutput(registers, instructions)
	if err != nil {
		return err
	}

	fmt.Println("Day 17a:", output)

	return nil
}

func Part2() error {
	_, instructions, err := input()
	if err != nil {
		return err
	}

	output, err := findLowestAValue(instructions)
	if err != nil {
		return err
	}

	fmt.Println("Day 17b:", output)

	return nil
}
