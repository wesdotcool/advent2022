package dec1

// https://adventofcode.com/2022/day/1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run(test bool) error {
	var filename string
	if test {
		filename = "dec1/test_input.txt"
	} else {
		filename = "dec1/input.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

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
		return err
	}

	var maxCalories int
	for _, curElf := range elves {
		if curElf.calories > maxCalories {
			maxCalories = curElf.calories
		}
	}
	fmt.Println(maxCalories)

	return nil
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
