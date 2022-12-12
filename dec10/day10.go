package dec10

// https://adventofcode.com/2022/day/10

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const noopCycles = 1
const addxCycles = 2

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	answer := 0
	cycle := 0
	x := 1
	for scanner.Scan() {
		instruction := parseLine(scanner.Text())
		if instruction.noop() {
			if crossesCycleBoundary(cycle, noopCycles) {
				answer += ((cycle + noopCycles) / 10 * 10) * x
			}
			cycle += noopCycles
		} else {
			if crossesCycleBoundary(cycle, addxCycles) {
				answer += ((cycle + addxCycles) / 10 * 10) * x
			}
			cycle += addxCycles
			x += instruction.arg
		}
	}
	return answer, nil
}

// start on a cycle
// read instruction
//
//	noop -> increment cycle by 1. print out pixel.
//	addx -> increment cycle by 1. print pixel. increment cycle by 1. add x, print
//
// printing a pixel:
// compare current x with current cycle % 40. if they are within 1, then print #
// if cycle % 40 == 1, print a newline
func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	answer := 0
	cycle := 0
	x := 1
	fmt.Printf("#")
	for scanner.Scan() {
		instruction := parseLine(scanner.Text())
		if instruction.noop() {
			cycle += noopCycles
			printPixel(x, cycle)
		} else {
			printPixel(x, cycle+1)
			cycle += addxCycles
			x += instruction.arg
			printPixel(x, cycle)
		}
	}
	return answer, nil
}

func printPixel(x, cycle int) {
	if cycle%40 == 0 {
		fmt.Printf("\n")
	}
	result := x - (cycle % 40)
	if result >= -1 && result <= 1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
}

func crossesCycleBoundary(cycle, cycleLength int) bool {
	return (cycle+cycleLength+20)%40 < (cycle+20)%40
}

type instruction struct {
	cmd string
	arg int
}

func (this *instruction) noop() bool {
	return this.cmd == "noop"
}

func parseLine(line string) *instruction {
	splits := strings.Split(line, " ")
	inst := &instruction{splits[0], 0}
	if len(splits) > 1 {
		inst.arg, _ = strconv.Atoi(splits[1])
	}
	return inst
}

// keep track of cycle
// look at instruction
//   if noop, increase cycle
//   if addx, increase cycle by 2 and increase x
// when (cycle + 20) % 40 == 0 then add signal strength to the answer
