package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			add(testcase.input.a, testcase.input.b),
			"Testcase: %+v", testcase.input)
	}
}
