package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputPath string = "inputs/2024/day5.txt"

type Rule struct {
	First  int
	Second int
}

type Update []int

func (u *Update) findIndex(num int) (int, error) {
	for i, value := range *u {
		if value == num {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no index found for %d", num)
}

func (u *Update) middleValue() int {
	return (*u)[len(*u)/2]
}

func (u *Update) isCorrect(rules []Rule) (bool, *Rule) {
	for _, rule := range rules {
		firstindex, err := u.findIndex(rule.First)
		if err != nil {
			continue
		}

		secondIndex, err := u.findIndex(rule.Second)
		if err != nil {
			continue
		}

		if firstindex > secondIndex {
			return false, &rule
		}
	}

	return true, nil
}

func (u *Update) fixFailingRule(rule *Rule) (Update, error) {
	result := make(Update, len(*u))
	copy(result, *u)

	firstIndex, err := result.findIndex(rule.First)
	if err != nil {
		return nil, err
	}

	secondIndex, err := result.findIndex(rule.Second)
	if err != nil {
		return nil, err
	}

	if firstIndex > secondIndex {
		result = append(result[:secondIndex], result[secondIndex+1:]...)

		if firstIndex > secondIndex {
			firstIndex--
		}

		result = append(result[:firstIndex+1], append([]int{rule.Second}, result[firstIndex+1:]...)...)
	}

	return result, nil
}

func input() ([]Rule, []Update, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := []Rule{}
	updates := []Update{}

	rulesFinished := false

	for scanner.Scan() {
		if scanner.Text() == "" {
			rulesFinished = true
			continue
		}

		if !rulesFinished {
			ruleValues := strings.Split(scanner.Text(), "|")
			first, err := strconv.Atoi(ruleValues[0])
			if err != nil {
				return nil, nil, err
			}

			second, err := strconv.Atoi(ruleValues[1])
			if err != nil {
				return nil, nil, err
			}

			rules = append(rules, Rule{First: first, Second: second})
		} else {
			updateValues := strings.Split(scanner.Text(), ",")
			update := Update{}
			for _, updateValue := range updateValues {
				updateValueInt, err := strconv.Atoi(updateValue)
				if err != nil {
					return nil, nil, err
				}

				update = append(update, updateValueInt)
			}

			updates = append(updates, update)
		}
	}

	return rules, updates, nil
}

func Part1() error {
	rules, updates, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, update := range updates {
		if correct, _ := update.isCorrect(rules); correct {
			total += update.middleValue()
		}
	}

	fmt.Println("day 5a:", total)

	return nil
}

func Part2() error {
	rules, updates, err := input()
	if err != nil {
		return err
	}

	total := 0

	for _, update := range updates {
		modified := false

		for true {
			correct, failingRule := update.isCorrect(rules)
			if correct {
				if modified {
					total += update.middleValue()
				}
				break
			} else {
				modified = true
				update, err = update.fixFailingRule(failingRule)
			}
		}
	}

	fmt.Println("day 5b:", total)

	return nil
}
