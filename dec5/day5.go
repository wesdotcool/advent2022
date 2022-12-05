package dec5

// https://adventofcode.com/2022/day/4

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/stacks/arraystack"
)

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	as := arraystack.New()
	_ = as
	var stacks []*arraystack.Stack
	var finalStacks []*arraystack.Stack
	answer := ""
	firstLineProcessed, stacksProcessed := false, false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("dumping blank line")
			continue
		}
		fmt.Println("processing line:", line)
		if !firstLineProcessed {
			stacks = allocateStacks(line)
			firstLineProcessed = true
			fmt.Println("proc first line")
			continue
		}
		if !stacksProcessed {
			stackLine := line
			if string(stackLine[1]) == "1" { // Last line of stacks
				fmt.Println("finished processing stack lines")
				fmt.Println("reversing stacks")
				for _, stack := range stacks {
					finalStacks = append(finalStacks, reverseStack(stack))
				}
				stacksProcessed = true
			} else {
				processStackLine(stacks, stackLine)
			}
			continue
		}
		processInstruction(finalStacks, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for _, stack := range finalStacks {
		if !stack.Empty() {
			crate, _ := stack.Pop()
			answer += crate.(string)
		}
	}
	return answer, nil
}

func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	as := arraystack.New()
	_ = as
	var stacks []*arraystack.Stack
	var finalStacks []*arraystack.Stack
	answer := ""
	firstLineProcessed, stacksProcessed := false, false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("dumping blank line")
			continue
		}
		fmt.Println("processing line:", line)
		if !firstLineProcessed {
			stacks = allocateStacks(line)
			firstLineProcessed = true
			fmt.Println("proc first line")
			continue
		}
		if !stacksProcessed {
			stackLine := line
			if string(stackLine[1]) == "1" { // Last line of stacks
				fmt.Println("finished processing stack lines")
				fmt.Println("reversing stacks")
				for _, stack := range stacks {
					finalStacks = append(finalStacks, reverseStack(stack))
				}
				stacksProcessed = true
			} else {
				processStackLine(stacks, stackLine)
			}
			continue
		}
		processInstruction9001(finalStacks, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for _, stack := range finalStacks {
		if !stack.Empty() {
			crate, _ := stack.Pop()
			answer += crate.(string)
		}
	}
	return answer, nil
}

func allocateStacks(firstLine string) []*arraystack.Stack {
	stacks := make([]*arraystack.Stack, 0)
	iter := &stringIterator{firstLine, 0}
	for {
		token := iter.nextToken()
		if token == "" {
			return stacks
		}
		stack := arraystack.New()
		if token != " " {
			stack.Push(token)
		}
		stacks = append(stacks, stack)
	}
}

type stringIterator struct {
	value    string
	position int
}

func (this *stringIterator) nextToken() string {
	if this.position >= len(this.value) {
		return "" // string is completely processed
	}
	result := string(this.value[this.position+1])
	this.position += 4
	return result
}

func reverseStack(stack *arraystack.Stack) *arraystack.Stack {
	result := arraystack.New()
	for !stack.Empty() {
		item, _ := stack.Pop()
		result.Push(item)
	}
	return result
}

// stackLine is a string of the form: "[N] [C]    "
// ((   )|(\[.]) ?)+
func processStackLine(stacks []*arraystack.Stack, stackLine string) {
	iter := &stringIterator{stackLine, 0}
	for i := range stacks {
		token := iter.nextToken()
		if token == "" {
			panic(fmt.Sprintf("stack line: %v was shorter than len(stacks): %v", stackLine, len(stacks)))
		}
		if token != " " {
			stacks[i].Push(token)
		}
	}
}

// instruction is a string of the form: "move 1 from 2 to 1"
func processInstruction(stacks []*arraystack.Stack, instruction string) {
	fmt.Println("processing instruction:", instruction)
	splits := strings.Split(instruction, " ")
	toMove, _ := strconv.Atoi(splits[1])
	source, _ := strconv.Atoi(splits[3])
	dest, _ := strconv.Atoi(splits[5])
	source -= 1 // indexing from 0
	dest -= 1   // indexing from 0
	for toMove != 0 {
		crate, _ := stacks[source].Pop()
		stacks[dest].Push(crate)
		toMove -= 1
	}
}

// instruction is a string of the form: "move 1 from 2 to 1"
// in the second part of the problem, the crane operates slightly differently
func processInstruction9001(stacks []*arraystack.Stack, instruction string) {
	fmt.Println("processing instruction:", instruction)
	splits := strings.Split(instruction, " ")
	toMove, _ := strconv.Atoi(splits[1])
	source, _ := strconv.Atoi(splits[3])
	dest, _ := strconv.Atoi(splits[5])
	source -= 1 // indexing from 0
	dest -= 1   // indexing from 0
	tempStack := arraystack.New()
	for toMove != 0 {
		crate, _ := stacks[source].Pop()
		tempStack.Push(crate)
		toMove -= 1
	}
	for !tempStack.Empty() {
		crate, _ := tempStack.Pop()
		stacks[dest].Push(crate)
	}
}
