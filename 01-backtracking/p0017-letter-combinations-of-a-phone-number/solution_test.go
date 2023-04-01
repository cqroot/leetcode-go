package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    string
	expected []string
}

func TestLetterCombinations(t *testing.T) {
	testcases := []testcase{
		{
			input:    "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, letterCombinations(testcase.input))
	}
}
