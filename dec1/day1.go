package dec1

// https://adventofcode.com/2022/day/1

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)

	var currentText string
	elves := make([]elf, 1)
	currentElfIndex := 0
	for scanner.Scan() {
		currentText = scanner.Text()
		if currentText == "" {
			currentElfIndex += 1
			elves = append(elves, elf{})
		} else {
			elves[currentElfIndex].addCalories(currentText)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var maxCalories int
	for _, curElf := range elves {
		if curElf.calories > maxCalories {
			maxCalories = curElf.calories
		}
	}

	return maxCalories, nil
}

type elf struct {
	calories int
}

func (this *elf) addCalories(calorieString string) error {
	calories, err := strconv.Atoi(calorieString)
	if err != nil {
		return err
	}
	this.calories += calories
	return nil
}

func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)

	var currentText string
	elves := make([]elf, 1)
	currentElfIndex := 0
	for scanner.Scan() {
		currentText = scanner.Text()
		if currentText == "" {
			currentElfIndex += 1
			elves = append(elves, elf{})
		} else {
			elves[currentElfIndex].addCalories(currentText)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
	answer := elves[0].calories + elves[1].calories + elves[2].calories

	return answer, nil
}
