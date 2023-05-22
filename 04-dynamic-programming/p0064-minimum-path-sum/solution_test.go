package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    [][]int
	expected int
}

func TestMinPathSum(t *testing.T) {
	testcases := []testcase{
		{
			input:    [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expected: 7,
		},
		{
			input:    [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 12,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			minPathSum(testcase.input),
			"Testcase: %+v", testcase.input)
	}
}
