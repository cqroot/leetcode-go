package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	a int
	b int
}

type testcase struct {
	input    input
	expected int
}

func TestAdd(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				a: 1,
				b: 1,
			},
			expected: 2,
		},
		{
			input: input{
				a: 2,
				b: 2,
			},
			expected: 4,
		},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, add(testcase.input.a, testcase.input.b))
	}
}
