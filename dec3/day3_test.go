package dec3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestDay3Part1(t *testing.T) {
	expected := 157
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}

func TestDay3Part2(t *testing.T) {
	expected := 70
	result, err := Run2(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}
