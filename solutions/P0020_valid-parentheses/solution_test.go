package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    string
	expected bool
}

func TestIsValid(t *testing.T) {
	testcases := []testcase{
		{
			input:    "()",
			expected: true,
		},
		{
			input:    "()[]{}",
			expected: true,
		},
		{
			input:    "(]",
			expected: false,
		},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, isValid_Stack(testcase.input))
	}
}
