package dec9

// https://adventofcode.com/2022/day/9

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
)

// Given a command like R 4
// Update the head
// check tail/head position
// update tail if needed
// save new tail position to set of points

// Data structures
//
//	head/tail are points, and a set of visited tail points
//	points need a couple methods:
//	- update position given a direction
//	- compare head/tail and update tail accordingly
func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	tailPositions := hashset.New()
	head, tail := &point{0, 0}, &point{0, 0}
	tailPositions.Add(*tail)
	for scanner.Scan() {
		direction, times := parseLine(scanner.Text())
		for times > 0 {
			head.move(direction)
			tail.updateTailPosition(head)
			tailPositions.Add(*tail)
			times -= 1
		}
	}
	answer := tailPositions.Size()
	return answer, nil
}

// Now the tail is 10 knots long, not just 2 (head/tail)
// Count the unique positions of the 10th knot
func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	tailPositions := hashset.New()
	knots := []*point{}
	numKnots := 10
	for i := 0; i < numKnots; i += 1 {
		knots = append(knots, &point{0, 0})
	}
	tailPositions.Add(*knots[numKnots-1])
	for scanner.Scan() {
		direction, times := parseLine(scanner.Text())
		for times > 0 {
			knots[0].move(direction)
			for i := 1; i < numKnots; i += 1 {
				knots[i].updateTailPosition(knots[i-1])
			}
			tailPositions.Add(*knots[numKnots-1])
			times -= 1
		}
	}
	answer := tailPositions.Size()
	return answer, nil
}

func parseLine(line string) (rune, int) {
	splits := strings.Split(line, " ")
	dir := rune(splits[0][0])
	times, _ := strconv.Atoi(splits[1])
	return dir, times
}

type point struct {
	x int
	y int
}

// Receives an instruction in the form direction, number
func (this *point) move(direction rune) {
	switch direction {
	case 'R':
		this.x += 1
	case 'L':
		this.x -= 1
	case 'U':
		this.y += 1
	case 'D':
		this.y -= 1
	default:
		panic(fmt.Sprintf("Unknown direction: %v", direction))
	}
}

// If adjacent, then return
// If not, decide direction to move and move
// To decide:
//
//	if in the same row/column, move 1 closer in that same row/column
//	otherwise, move diagonally towards head
func (tail *point) updateTailPosition(head *point) {
	// The real code
	// head.y is 2 away
	if tail.y-head.y == 2 {
		tail.y -= 1
		if tail.x-head.x > 0 {
			tail.x -= 1
		}
		if head.x-tail.x > 0 {
			tail.x += 1
		}
		return
	}
	if head.y-tail.y == 2 {
		tail.y += 1
		if tail.x-head.x > 0 {
			tail.x -= 1
		}
		if head.x-tail.x > 0 {
			tail.x += 1
		}
		return
	}
	// head.x is 2 away
	if tail.x-head.x == 2 {
		tail.x -= 1
		if tail.y-head.y > 0 {
			tail.y -= 1
		}
		if head.y-tail.y > 0 {
			tail.y += 1
		}
		return
	}
	if head.x-tail.x == 2 {
		tail.x += 1
		if tail.y-head.y > 0 {
			tail.y -= 1
		}
		if head.y-tail.y > 0 {
			tail.y += 1
		}
		return
	}
}
