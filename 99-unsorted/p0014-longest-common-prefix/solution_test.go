package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    []string
	expected string
}

func TestLongestCommonPrefix(t *testing.T) {
	testcases := []testcase{
		{
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			longestCommonPrefix(testcase.input),
			"Testcase: %+v", testcase.input)
	}
}
