package dec2

// https://adventofcode.com/2022/day/2

import (
	"bufio"
	"io"
	"strings"
)

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)

	totalPoints := 0
	for scanner.Scan() {
		totalPoints += computePoints(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return totalPoints, nil
}

func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)

	totalPoints := 0
	for scanner.Scan() {
		totalPoints += computePoints2(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return totalPoints, nil
}

// ROCK     = A, X, 1
// PAPER    = B, Y, 2
// SCISSORS = C, Z, 3
func computePoints(rpsMatch string) int {
	split := strings.Split(rpsMatch, " ")
	points := 0
	if split[1] == "X" { // I PLAYED ROCK
		points += 1
		if split[0] == "A" { // draw
			points += 3
		}
		if split[0] == "B" { // lose
			points += 0
		}
		if split[0] == "C" { // win
			points += 6
		}
	}
	if split[1] == "Y" { // I PLAYED PAPER
		points += 2
		if split[0] == "A" { // win
			points += 6
		}
		if split[0] == "B" { // draw
			points += 3
		}
		if split[0] == "C" { // lose
			points += 0
		}
	}
	if split[1] == "Z" { // I PLAYED SCISSORS
		points += 3
		if split[0] == "A" { // lose
			points += 0
		}
		if split[0] == "B" { // win
			points += 6
		}
		if split[0] == "C" { // draw
			points += 3
		}
	}
	return points
}

const rock = "A"
const paper = "B"
const scissors = "C"

// ROCK     = A, 1
// PAPER    = B, 2
// SCISSORS = C, 3
// X = LOSE
// Y = DRAW
// Z = WIN
func computePoints2(rpsMatch string) int {
	split := strings.Split(rpsMatch, " ")
	points := 0
	if split[1] == "X" { // I SHOULD LOSE
		points += 0
		if split[0] == rock {
			points += 3
		}
		if split[0] == paper {
			points += 1
		}
		if split[0] == scissors {
			points += 2
		}
	}
	if split[1] == "Y" { // I SHOULD DRAW
		points += 3
		if split[0] == rock {
			points += 1
		}
		if split[0] == paper {
			points += 2
		}
		if split[0] == scissors {
			points += 3
		}
	}
	if split[1] == "Z" { // I SHOULD WIN
		points += 6
		if split[0] == rock {
			points += 2
		}
		if split[0] == paper {
			points += 3
		}
		if split[0] == scissors {
			points += 1
		}
	}
	return points
}
