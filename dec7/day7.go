package dec7

// https://adventofcode.com/2022/day/7

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// read a cd command -> execute immediately

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	answer := 0
	rootDir := &directory{
		name:    "/",
		subdirs: make([]*directory, 0),
		files:   make([]int, 0),
	}
	currentDir := rootDir
	lineToProcess := ""
	for true {
		if lineToProcess == "" {
			notEOF := scanner.Scan()
			if !notEOF {
				break
			}
			lineToProcess = scanner.Text()
		}
		cmd := splitCommand(lineToProcess)
		if cmd.cd() {
			switch cmd.dirName() {
			case "..":
				currentDir = currentDir.parent
			case "/":
				currentDir = rootDir
			default:
				currentDir = currentDir.getSubDir(cmd.dirName())
			}
			lineToProcess = ""
		} else { // It's an ls command
			lineToProcess = ""
			for scanner.Scan() {
				line := scanner.Text()
				if line[0] != '$' {
					toProcess := makeLsOutput(line)
					if toProcess.isDir() {
						currentDir.addDir(toProcess.name())
					} else {
						currentDir.addFile(toProcess.file())
					}
				} else {
					lineToProcess = line
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	answer = sumValuesUnder(rootDir.traverseTree(), 100000)
	return answer, nil
}

func (this *directory) addDir(name string) {
	newDir := &directory{
		name:    name,
		parent:  this,
		subdirs: make([]*directory, 0),
		files:   make([]int, 0),
	}
	this.subdirs = append(this.subdirs, newDir)
}

func FooBar() {}

func (this *directory) addFile(fileSize int) {
	this.files = append(this.files, fileSize)
}

type lsOutput []string
type commandSplit []string

func makeLsOutput(line string) lsOutput {
	return lsOutput(strings.Split(line, " "))
}
func (this lsOutput) isDir() bool {
	return this[0] == "dir"
}

func (this lsOutput) name() string {
	return this[1]
}

// just returns the files size. files are just ints
func (this lsOutput) file() int {
	s, _ := strconv.Atoi(this[0])
	return s
}

func splitCommand(commandText string) commandSplit {
	return strings.Split(commandText, " ")
}

// panics if called on an ls command
func (this commandSplit) dirName() string {
	return this[2]
}
func (this commandSplit) cd() bool {
	return this[1] == "cd"
}
func (this commandSplit) ls() bool {
	return this[1] == "ls"
}

func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	answer := 0
	rootDir := &directory{
		name:    "/",
		subdirs: make([]*directory, 0),
		files:   make([]int, 0),
	}
	currentDir := rootDir
	lineToProcess := ""
	for true {
		if lineToProcess == "" {
			notEOF := scanner.Scan()
			if !notEOF {
				break
			}
			lineToProcess = scanner.Text()
		}
		cmd := splitCommand(lineToProcess)
		if cmd.cd() {
			switch cmd.dirName() {
			case "..":
				currentDir = currentDir.parent
			case "/":
				currentDir = rootDir
			default:
				currentDir = currentDir.getSubDir(cmd.dirName())
			}
			lineToProcess = ""
		} else { // It's an ls command
			lineToProcess = ""
			for scanner.Scan() {
				line := scanner.Text()
				if line[0] != '$' {
					toProcess := makeLsOutput(line)
					if toProcess.isDir() {
						currentDir.addDir(toProcess.name())
					} else {
						currentDir.addFile(toProcess.file())
					}
				} else {
					lineToProcess = line
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	spaceNeededToClear := 30000000 - (70000000 - rootDir.totalSize())
	answer = findLowestOver(rootDir.traverseTree(), spaceNeededToClear)
	return answer, nil
}

// read commands
// Read a line
//   if a $ cd line, change the directory
//   if a $ ls line, read the following lines until you find another $ line

// construct tree of directories
// cd => change the currentDir to a directory, either the parent or a subdir
// ls => populate subdirs and files for the currentDir

// write function to sum up file sizes given 1 node
// traverse the tree and calculate each file size
type directory struct {
	parent  *directory
	name    string
	subdirs []*directory
	files   []int
}

func (this *directory) directFileSize() int {
	total := 0
	for _, v := range this.files {
		total += v
	}
	return total
}

func (this *directory) totalSize() int {
	total := 0
	for _, dir := range this.subdirs {
		total += dir.totalSize()
	}
	total += this.directFileSize()
	return total
}

// Returns a slice of the total size of every directory below this, including this directory
func (this *directory) traverseTree() []int {
	result := []int{this.totalSize()}
	for _, dir := range this.subdirs {
		result = append(result, dir.traverseTree()...)
	}
	return result
}

// sums all values in the slice that are <= maxValue
func sumValuesUnder(values []int, maxValue int) int {
	total := 0
	for _, v := range values {
		if v <= maxValue {
			total += v
		}
	}
	return total
}

func findLowestOver(values []int, minValue int) int {
	lowestOver := 3000000000
	for _, v := range values {
		if v >= minValue && v < lowestOver {
			lowestOver = v
		}
	}
	return lowestOver
}

func (this *directory) getSubDir(name string) *directory {
	for _, v := range this.subdirs {
		if v.name == name {
			return v
		}
	}
	panic(fmt.Sprintf("Failed to find subdirectory %v in %v", name, this.name))
}
