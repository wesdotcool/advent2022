package dec2

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText string = `A Y
B X
C Z
`

func TestDay2Part1(t *testing.T) {
	expected := 15
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, result, expected)
}

func TestDay2Part2(t *testing.T) {
	expected := 12
	result, err := Run2(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, result, expected)
}
