package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    int
	expected int
}

func TestGuessNumber(t *testing.T) {
	testcases := []testcase{
		{
			input:    10,
			expected: 6,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 1,
		},
	}

	for _, testcase := range testcases {
		pick = testcase.expected
		require.Equal(t, testcase.expected, guessNumber(testcase.input))
	}
}
