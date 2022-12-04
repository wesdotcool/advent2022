package dec4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func TestDay3Part1(t *testing.T) {
	expected := 2
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}

func TestDay3Part2(t *testing.T) {
	expected := 4
	result, err := Run2(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}
