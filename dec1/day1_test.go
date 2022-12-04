package dec1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestDay1Part1(t *testing.T) {
	expected := 24000
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, result, expected)
}

func TestDay1Part2(t *testing.T) {
	expected := 45000
	result, err := Run2(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, result, expected)
}
