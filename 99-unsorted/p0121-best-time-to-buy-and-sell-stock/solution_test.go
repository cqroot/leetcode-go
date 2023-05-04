package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    []int
	expected int
}

func TestMaxProfit(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			maxProfit(testcase.input),
			"Testcase: %+v", testcase.input)
	}
}
