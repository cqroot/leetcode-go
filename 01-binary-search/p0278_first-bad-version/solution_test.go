package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type input struct {
	n   int
	bad int
}

type testcase struct {
	input    input
	expected int
}

func TestFirstBadVersion(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				n:   5,
				bad: 4,
			},
			expected: 4,
		},
		{
			input: input{
				n:   1,
				bad: 1,
			},
			expected: 1,
		},
		{
			input: input{
				n:   5,
				bad: 5,
			},
			expected: 5,
		},
	}

	for _, testcase := range testcases {
		badVersion = testcase.input.bad
		require.Equal(t, testcase.expected, firstBadVersion(testcase.input.n))
	}
}
