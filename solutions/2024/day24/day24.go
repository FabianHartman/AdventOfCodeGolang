package day24

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var inputPath = "inputs/2024/day24.txt"

func input() (map[string]int, map[string][]string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't open %s: %v", inputPath, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	initialValues := map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		lineParts := strings.Split(line, ": ")
		if len(lineParts) != 2 {
			return nil, nil, fmt.Errorf("invalid line: %s", lineParts)
		}

		initialValues[lineParts[0]], err = strconv.Atoi(lineParts[1])
	}

	pattern := `([a-zA-Z0-9]+)\s+(AND|XOR|OR)\s+([a-zA-Z0-9]+)\s*->\s*([a-zA-Z0-9]+)`
	re := regexp.MustCompile(pattern)

	gates := map[string][]string{}

	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		if match == nil {
			return nil, nil, fmt.Errorf("invalid line: %s", scanner.Text())
		}

		gates[match[4]] = append(gates[match[4]], match[1])
		gates[match[4]] = append(gates[match[4]], match[3])
		gates[match[4]] = append(gates[match[4]], match[2])
	}

	return initialValues, gates, scanner.Err()
}

func calculatePortValues(values map[string]int, gates map[string][]string) map[string]int {
	for {
		gatesToCalculate := []string{}
		for gate := range gates {
			if _, exists := values[gate]; !exists {
				gatesToCalculate = append(gatesToCalculate, gate)
			}
		}

		if len(gatesToCalculate) == 0 {
			break
		}

		for _, gate := range gatesToCalculate {
			value1, value2 := 0, 0
			exists := false

			currentGate := gates[gate]
			if value1, exists = values[currentGate[0]]; !exists {
				continue
			}

			if value2, exists = values[currentGate[1]]; !exists {
				continue
			}

			switch currentGate[2] {
			case "AND":
				if value1 == 1 && value2 == 1 {
					values[gate] = 1
				} else {
					values[gate] = 0
				}
			case "OR":
				if value1 == 1 || value2 == 1 {
					values[gate] = 1
				} else {
					values[gate] = 0
				}
			case "XOR":
				if value1 != value2 {
					values[gate] = 1
				} else {
					values[gate] = 0
				}
			}
		}
	}

	return values
}

func getZOutput(values map[string]int) string {
	zGates := []string{}

	for gate := range values {
		if gate[0] == 'z' {
			zGates = append(zGates, gate)
		}
	}

	sort.Strings(zGates)

	binaryOutputString := ""
	for _, gate := range zGates {
		gateValue := strconv.Itoa(values[gate])
		binaryOutputString = gateValue + binaryOutputString
	}

	return binaryOutputString
}

func convertBinaryToInt(binary string) (int, error) {
	num, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return 0, err
	}

	return int(num), nil
}

func findGate(inputWire1, inputWire2, gateOperator string, gates map[string][]string) string {
	for gate, parts := range gates {
		if (parts[0] == inputWire1 && parts[1] == inputWire2 && parts[2] == gateOperator) ||
			(parts[0] == inputWire2 && parts[1] == inputWire1 && parts[2] == gateOperator) {
			return gate
		}
	}

	return ""
}

func swapWires(gates map[string][]string) (string, error) {
	var swapped []string
	var carry string

	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var xorCarryXor, nextCarry string

		xor := findGate("x"+n, "y"+n, "XOR", gates)
		and := findGate("x"+n, "y"+n, "AND", gates)

		if carry != "" {
			andCarryXor := findGate(carry, xor, "AND", gates)
			if andCarryXor == "" {
				xor, and = and, xor
				swapped = append(swapped, xor, and)
				andCarryXor = findGate(carry, xor, "AND", gates)
			}

			xorCarryXor = findGate(carry, xor, "XOR", gates)

			if strings.HasPrefix(xor, "z") {
				xor, xorCarryXor = xorCarryXor, xor
				swapped = append(swapped, xor, xorCarryXor)
			}

			if strings.HasPrefix(and, "z") {
				and, xorCarryXor = xorCarryXor, and
				swapped = append(swapped, and, xorCarryXor)
			}

			if strings.HasPrefix(andCarryXor, "z") {
				andCarryXor, xorCarryXor = xorCarryXor, andCarryXor
				swapped = append(swapped, andCarryXor, xorCarryXor)
			}

			nextCarry = findGate(andCarryXor, and, "OR", gates)
		}

		if strings.HasPrefix(nextCarry, "z") && nextCarry != "z45" {
			nextCarry, xorCarryXor = xorCarryXor, nextCarry
			swapped = append(swapped, nextCarry, xorCarryXor)
		}

		if carry == "" {
			carry = and
		} else {
			carry = nextCarry
		}
	}

	sort.Strings(swapped)

	return strings.Join(swapped, ","), nil
}

func Part1() error {
	initialValues, gates, err := input()
	if err != nil {
		return err
	}

	output, err := convertBinaryToInt(getZOutput(calculatePortValues(initialValues, gates)))
	if err != nil {
		return err
	}

	fmt.Println("Day24a:", output)

	return nil
}

func Part2() error {
	_, gates, err := input()
	if err != nil {
		return err
	}

	swapped, err := swapWires(gates)
	if err != nil {
		return err
	}

	fmt.Println("Day24b:", swapped)
	return nil
}
