package dec3

// https://adventofcode.com/2022/day/2

import (
	"bufio"
	"io"

	mapset "github.com/deckarep/golang-set/v2"
)

func Run(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)

	answer := 0
	for scanner.Scan() {
		answer += makeRucksack(items(scanner.Text())).
			misplacedItem().
			priority()
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return answer, nil
}

func Run2(input io.Reader) (any, error) {
	scanner := bufio.NewScanner(input)

	answer := 0
	elfGroup := [3]items{}
	elfIndex := 0
	for scanner.Scan() {
		elfGroup[elfIndex] = items(scanner.Text())
		elfIndex += 1
		if elfIndex == 3 {
			answer += badgePriority(elfGroup)
			elfIndex = 0
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return answer, nil
}

func badgePriority(this [3]items) int {
	s1 := mapset.NewSet([]byte(this[0])...)
	s2 := mapset.NewSet([]byte(this[1])...)
	s3 := mapset.NewSet([]byte(this[2])...)
	badge, _ := s1.Intersect(s2).Intersect(s3).Pop()
	return items(badge).priority()
}

type items string

func (this items) split() (items, items) {
	return this[0 : len(this)/2], this[len(this)/2:]
}

func (this items) priority() int {
	total := 0
	for _, char := range []byte(this) {
		total += bytePriority(char)
	}
	return total
}

func bytePriority(b byte) int {
	var result int = int(b) - 96
	if result > 0 {
		return result
	}
	return result + 58
}

type rucksack struct {
	compartment1 items
	compartment2 items
}

func makeRucksack(itms items) *rucksack {
	c1, c2 := itms.split()
	return &rucksack{c1, c2}
}

func (this *rucksack) misplacedItem() items {
	s1 := mapset.NewSet([]byte(this.compartment1)...)
	s2 := mapset.NewSet([]byte(this.compartment2)...)
	b, _ := s1.Intersect(s2).Pop()
	return items(b)
}
