package dec5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText string = `    [D]   
[N] [C]   
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestDay3Part1(t *testing.T) {
	expected := "CMZ"
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}

func TestDay3Part2(t *testing.T) {
	expected := "MCD"
	result, err := Run2(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}
