package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    string
	expected int
}

func TestRomanToInt(t *testing.T) {
	testcases := []testcase{
		{
			input:    "III",
			expected: 3,
		},
		{
			input:    "LVIII",
			expected: 58,
		},
		{
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			romanToInt(testcase.input),
			"Testcase: %+v", testcase.input)
	}
}
