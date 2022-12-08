package dec6

// https://adventofcode.com/2022/day/6

import (
	"bufio"
	"io"
	"sort"

	"golang.org/x/exp/slices"
)

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	answer := 0
	for scanner.Scan() {
		answer = charsToProcess(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return answer, nil
}

func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	answer := 0
	for scanner.Scan() {
		answer = charsToProcess14(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return answer, nil
}

func charsToProcess(s string) int {
	for i := range s {
		if fourDiff(s, i) {
			return i + 4
		}
	}
	return -1
}

func charsToProcess14(s string) int {
	for i := range s {
		if fourteenDiff(s, i) {
			return i + 14
		}
	}
	return -1
}

func fourDiff(s string, start int) bool {
	a, b, c, d := s[start], s[start+1], s[start+2], s[start+3]
	return a != b && a != c && a != d && b != c && b != d && c != d
}

func fourteenDiff(s string, start int) bool {
	runes := []rune(s)[start : start+14]
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return len(runes) == len(slices.Compact(runes))
}
