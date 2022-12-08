package dec6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText string = `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`

func TestDay6Part1(t *testing.T) {
	expected := 11
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}

func TestDay6Part2(t *testing.T) {
	expected := 26
	result, err := Run2(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}
