package dec9

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`
const testTextLong string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

// Unit testing methods
func TestDay8UpdateTailPosition1(t *testing.T) {
	// Situation T.H -> .TH
	tail := &point{0, 0}
	head := &point{2, 0}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{1, 0})
}

func TestDay8UpdateTailPosition2(t *testing.T) {
	// Situation H.T -> HT.
	tail := &point{2, 0}
	head := &point{0, 0}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{1, 0})
}

func TestDay8UpdateTailPosition3(t *testing.T) {
	// Situation:
	// T    .
	// . -> T
	// H    H
	tail := &point{0, 2}
	head := &point{0, 0}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{0, 1})
}

func TestDay8UpdateTailPosition4(t *testing.T) {
	// Situation:
	// H    H
	// . -> T
	// T    .
	tail := &point{0, 0}
	head := &point{0, 2}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{0, 1})
}

func TestDay8UpdateTailPosition5(t *testing.T) {
	// Situation:
	// .H    .H
	// .. -> .T
	// T.    ..
	tail := &point{0, 0}
	head := &point{1, 2}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{1, 1})
}

func TestDay8UpdateTailPosition6(t *testing.T) {
	// Situation:
	// .T    ..
	// .. -> T.
	// H.    H.
	tail := &point{1, 2}
	head := &point{0, 0}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{0, 1})
}

func TestDay8UpdateTailPosition7(t *testing.T) {
	// Situation:
	// ...    ...
	// ..T -> ...
	// H..    HT.
	tail := &point{2, 1}
	head := &point{0, 0}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{1, 0})
}

func TestDay8UpdateTailPosition8(t *testing.T) {
	// Situation:
	// ...    ...
	// ..H -> .TH
	// T..    ...
	tail := &point{0, 0}
	head := &point{2, 1}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{1, 1})
}

func TestDay8UpdateTailPosition9(t *testing.T) {
	// Situation:
	// ...    ...
	// .H. -> .H.
	// ...    ...
	tail := &point{1, 1}
	head := &point{1, 1}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{1, 1})
}

func TestDay8UpdateTailPosition10(t *testing.T) {
	// Situation:
	// ..T    ..T
	// .H. -> .H.
	// ...    ...
	tail := &point{2, 2}
	head := &point{1, 1}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{2, 2})
}

func TestDay8UpdateTailPosition11(t *testing.T) {
	// Situation:
	// ...    ...
	// .HT -> .HT
	// ...    ...
	tail := &point{2, 1}
	head := &point{1, 1}
	tail.updateTailPosition(head)
	assert.Equal(t, tail, &point{2, 1})
}

// Overall tests
func TestDay8Part1(t *testing.T) {
	expected := 13
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}

func TestDay8Part2(t *testing.T) {
	expected := 36
	result, err := Run2(strings.NewReader(testTextLong))
	assert.Equal(t, err, nil)
	assert.Equal(t, expected, result)
}
