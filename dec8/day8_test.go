package dec8

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText string = `30373
25512
65332
33549
35390
`

func TestDay8Part1(t *testing.T) {
	expected := 21
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}

func TestDay8Part2(t *testing.T) {
	expected := 8
	result, err := Run2(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}
