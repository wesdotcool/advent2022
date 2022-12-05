package dec4

// https://adventofcode.com/2022/day/4

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)

	answer := 0
	for scanner.Scan() {
		ranges := makeRanges(scanner.Text())
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i].len() > ranges[j].len()
		})
		if ranges[0].fullyCovers(ranges[1]) {
			answer += 1
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return answer, nil
}

func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)
	answer := 0
	var ranges []*assignmentRange
	for scanner.Scan() {
		ranges = makeRanges(scanner.Text())
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i].len() > ranges[j].len()
		})
		if ranges[0].hasAnyOverlap(ranges[1]) {
			answer += 1
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return answer, nil
}

func (this *assignmentRange) hasAnyOverlap(other *assignmentRange) bool {
	return (this.start <= other.start && this.end >= other.start) ||
		(this.start <= other.end && this.end >= other.end)
}

func (this *assignmentRange) fullyCovers(other *assignmentRange) bool {
	if other.start >= this.start && other.start <= this.end &&
		other.end >= this.start && other.end <= this.end {
		return true
	}
	return false
}

type assignmentRange struct {
	start int
	end   int
}

func makeRange(description string) *assignmentRange {
	splits := strings.Split(description, "-")
	s, _ := strconv.Atoi(splits[0])
	e, _ := strconv.Atoi(splits[1])
	r := &assignmentRange{start: s, end: e}
	return r
}

func makeRanges(csvLine string) []*assignmentRange {
	splits := strings.Split(csvLine, ",")
	result := make([]*assignmentRange, 0, len(splits))
	for _, v := range splits {
		result = append(result, makeRange(v))
	}
	return result
}

func (this *assignmentRange) len() int {
	return this.end - this.start + 1
}

func (this *assignmentRange) contains(n int) bool {
	return this.start <= n && this.end >= n
}
