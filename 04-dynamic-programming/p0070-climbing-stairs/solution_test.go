package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    int
	expected int
}

func TestClimbStairs(t *testing.T) {
	testcases := []testcase{
		{
			input:    2,
			expected: 2,
		},
		{
			input:    3,
			expected: 3,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			climbStairs(testcase.input),
			"Testcase: %+v", testcase.input)
	}
}
