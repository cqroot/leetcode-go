package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type input struct {
	m int
	n int
}

type testcase struct {
	input    input
	expected int
}

func TestUniquePaths(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				m: 3,
				n: 7,
			},
			expected: 28,
		},
		{
			input: input{
				m: 3,
				n: 2,
			},
			expected: 3,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			uniquePaths(testcase.input.m, testcase.input.n),
			"Testcase: %+v", testcase.input)
	}
}
