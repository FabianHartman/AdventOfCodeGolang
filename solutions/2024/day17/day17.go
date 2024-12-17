package day17

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPath string = "inputs/2024/day17.txt"

func input() (*Computer, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %v", inputPath, err)
	}

	scanner := bufio.NewScanner(file)
	registers := map[string]int{}

	scanner.Scan()
	registers["A"], err = strconv.Atoi(scanner.Text()[12:])
	if err != nil {
		return nil, fmt.Errorf("error parsing register A: %v", err)
	}

	scanner.Scan()
	registers["B"], err = strconv.Atoi(scanner.Text()[12:])
	if err != nil {
		return nil, fmt.Errorf("error parsing register B: %v", err)
	}

	scanner.Scan()
	registers["C"], err = strconv.Atoi(scanner.Text()[12:])
	if err != nil {
		return nil, fmt.Errorf("error parsing register C: %v", err)
	}

	scanner.Scan()
	scanner.Scan()

	stringProgram := strings.Split(scanner.Text()[9:], ",")

	program := make([]int, len(stringProgram))
	for i, v := range stringProgram {
		program[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("error parsing program instruction: %v", err)
		}
	}

	return &Computer{
		Registers: registers,
		Program:   program,
	}, nil
}

type Computer struct {
	Registers map[string]int
	Program   []int
	currentI  int
	Output    []int
}

func (this *Computer) copyInitial() *Computer {
	registers := map[string]int{}
	for k, v := range this.Registers {
		registers[k] = v
	}

	program := make([]int, len(this.Program))
	copy(program, this.Program)

	return &Computer{
		Registers: registers,
		Program:   program,
	}
}

func (this *Computer) adv(operand int) {
	numerator := this.Registers["A"]
	denominator := 1 << this.getComboValue(operand)

	this.Registers["A"] = numerator / denominator
	this.currentI += 2
}

func (this *Computer) bdv(operand int) {
	numerator := this.Registers["A"]
	denominator := 1 << this.getComboValue(operand)

	this.Registers["B"] = numerator / denominator
	this.currentI += 2
}

func (this *Computer) cdv(operand int) {
	numerator := this.Registers["A"]
	denominator := 1 << this.getComboValue(operand)

	this.Registers["C"] = numerator / denominator
	this.currentI += 2
}

func (this *Computer) bxl(operand int) {
	this.Registers["B"] ^= operand
	this.currentI += 2
}

func (this *Computer) bst(operand int) {
	this.Registers["B"] = this.getComboValue(operand) % 8
	this.currentI += 2
}

func (this *Computer) jnz(operand int) {
	if this.Registers["A"] == 0 {
		this.currentI += 2
		return
	}
	this.currentI = operand
}

func (this *Computer) bxc() {
	this.Registers["B"] ^= this.Registers["C"]
	this.currentI += 2
}

func (this *Computer) out(operand int) {
	this.Output = append(this.Output, this.getComboValue(operand)%8)
	this.currentI += 2
}

func (this *Computer) getComboValue(operand int) int {
	switch operand {
	case 4:
		return this.Registers["A"]
	case 5:
		return this.Registers["B"]
	case 6:
		return this.Registers["C"]
	default:
		return operand
	}
}

func (this *Computer) runCorrectInstruction(opcode int, operand int) error {
	switch opcode {
	case 0:
		this.adv(operand)
	case 1:
		this.bxl(operand)
	case 2:
		this.bst(operand)
	case 3:
		this.jnz(operand)
	case 4:
		this.bxc()
	case 5:
		this.out(operand)
	case 6:
		this.bdv(operand)
	case 7:
		this.cdv(operand)
	default:
		return fmt.Errorf("unknown opcode: %d", opcode)
	}

	return nil
}

func (this *Computer) getJoinedOutput() string {
	output := []string{}

	for _, value := range this.Output {
		output = append(output, strconv.Itoa(value))
	}

	return strings.Join(output, ",")
}

func (this *Computer) runProgram() (string, error) {
	for this.currentI < len(this.Program)-1 {
		opcode, operand := this.Program[this.currentI], this.Program[this.currentI+1]

		err := this.runCorrectInstruction(opcode, operand)
		if err != nil {
			return "", err
		}
	}

	return this.getJoinedOutput(), nil
}

func (c *Computer) findQuine() (int, error) {
	for i := 0; ; i++ {
		computer := c.copyInitial()
		computer.Registers["A"] = i
		output, err := computer.runProgram()
		if err != nil {
			return 0, err
		}

		if isMatchingOutput(output, computer.Program) {
			return i, nil
		}
	}
}

func isMatchingOutput(output string, expectedProgram []int) bool {
	expectedOutput := []string{}
	for _, v := range expectedProgram {
		expectedOutput = append(expectedOutput, strconv.Itoa(v))
	}

	return output == strings.Join(expectedOutput, ",")
}

func Part1() error {
	computer, err := input()
	if err != nil {
		return err
	}

	output, err := computer.runProgram()
	if err != nil {
		return err
	}

	fmt.Println("Day17a:", output)

	return nil
}

func Part2() error {
	computer, err := input()
	if err != nil {
		return err
	}

	initialValue, err := computer.findQuine()
	if err != nil {
		return err
	}

	fmt.Println("Day17b:", initialValue)
	return nil
}
