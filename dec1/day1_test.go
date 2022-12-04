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

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	expected := 24000
	result, err := Run(strings.NewReader(testText))
	assert.Equal(t, err, nil)
	assert.Equal(t, result, expected)
}
