package dec8

// https://adventofcode.com/2022/day/8

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func makeTreeGrid(scanner *bufio.Scanner) treeGrid {
	grid := [][]*tree{}
	row := 0
	for scanner.Scan() {
		treeLine := scanner.Text()
		grid = append(grid, []*tree{})
		for _, c := range treeLine {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			grid[row] = append(grid[row], &tree{
				height:  height,
				visible: false,
			})
		}
		row += 1
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return treeGrid(grid)
}

func printGrid(grid treeGrid) {
	for _, row := range grid {
		for _, elem := range row {
			fmt.Printf("%v", elem.height)
		}
		fmt.Printf("\n")
	}
}

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	grid := makeTreeGrid(scanner)
	grid.assignVisible()
	answer := grid.visibleTrees()
	return answer, nil
}

func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	grid := makeTreeGrid(scanner)
	grid.assignVisible()
	answer := grid.computeHighestScenicScore()
	return answer, nil
}

// Create 2D grid of trees
// write function to set visibility

type tree struct {
	height  int
	visible bool
}

type treeGrid [][]*tree

func (this treeGrid) assignVisible() {
	totalRows := len(this)
	totalColumns := len(this[0])

	// iterate through from left side
	for curRow := 0; curRow < totalRows; curRow += 1 {
		lowestHeight := -1
		for curCol := 0; curCol < totalColumns; curCol += 1 {
			if lowestHeight < this[curRow][curCol].height {
				this[curRow][curCol].visible = true
				lowestHeight = this[curRow][curCol].height
			}
		}
	}

	// iterate through from right side
	for curRow := 0; curRow < totalRows; curRow += 1 {
		lowestHeight := -1
		for curCol := totalColumns - 1; curCol >= 0; curCol -= 1 {
			if lowestHeight < this[curRow][curCol].height {
				this[curRow][curCol].visible = true
				lowestHeight = this[curRow][curCol].height
			}
		}
	}

	// iterate through from top
	for curCol := 0; curCol < totalColumns; curCol += 1 {
		lowestHeight := -1
		for curRow := 0; curRow < totalRows; curRow += 1 {
			if lowestHeight < this[curRow][curCol].height {
				this[curRow][curCol].visible = true
				lowestHeight = this[curRow][curCol].height
			}
		}
	}

	// iterate through from bottom
	for curCol := 0; curCol < totalColumns; curCol += 1 {
		lowestHeight := -1
		for curRow := totalRows - 1; curRow >= 0; curRow -= 1 {
			if lowestHeight < this[curRow][curCol].height {
				this[curRow][curCol].visible = true
				lowestHeight = this[curRow][curCol].height
			}
		}
	}
}

func (this treeGrid) visibleTrees() int {
	visible := 0
	for _, treeLine := range this {
		for _, tree := range treeLine {
			if tree.visible {
				visible += 1
			}
		}
	}
	return visible
}

// Compute scenicScore of the tree as specified row, col
// Panics if its out of bounds
func (this treeGrid) scenicScore(row, col int) int {
	totalRows := len(this)
	totalColumns := len(this[0])
	treeHeight := this[row][col].height

	visibleRight, visibleLeft, visibleUp, visibleDown := 0, 0, 0, 0
	// look to the right
	for curCol := col + 1; curCol < totalColumns; curCol += 1 {
		if treeHeight > this[row][curCol].height {
			visibleRight += 1
		} else {
			visibleRight += 1
			break
		}
	}

	// look to the left
	for curCol := col - 1; curCol >= 0; curCol -= 1 {
		if treeHeight > this[row][curCol].height {
			visibleLeft += 1
		} else {
			visibleLeft += 1
			break
		}
	}

	// look up
	for curRow := row + 1; curRow < totalRows; curRow += 1 {
		if treeHeight > this[curRow][col].height {
			visibleUp += 1
		} else {
			visibleUp += 1
			break
		}
	}

	// look down
	for curRow := row - 1; curRow >= 0; curRow -= 1 {
		if treeHeight > this[curRow][col].height {
			visibleDown += 1
		} else {
			visibleDown += 1
			break
		}
	}

	//fmt.Printf("scenic(%v, %v) => %v * %v * %v * %v\n", row, col, visibleLeft, visibleUp, visibleRight, visibleDown)
	return visibleRight * visibleLeft * visibleUp * visibleDown
}

func (this treeGrid) computeHighestScenicScore() int {
	max := 0
	for row := range this {
		for col := range this[row] {
			scenicScore := this.scenicScore(row, col)
			//fmt.Printf("height(%v, %v) = %v\n", row, col, scenicScore)
			if max < scenicScore {
				max = scenicScore
			}
		}
	}
	return max
}
