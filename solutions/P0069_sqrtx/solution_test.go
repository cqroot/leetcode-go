package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    int
	expected int
}

func TestMySqrt(t *testing.T) {
	testcases := []testcase{
		{
			input:    4,
			expected: 2,
		},
		{
			input:    8,
			expected: 2,
		},
		{
			input:    6,
			expected: 2,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, mySqrt(testcase.input))
	}
}
