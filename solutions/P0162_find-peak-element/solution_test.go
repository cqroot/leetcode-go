package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    []int
	expected int
}

func TestFindPeakElement(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{1, 2, 3, 1},
			expected: 2,
		},
		{
			input:    []int{1, 2, 1, 3, 5, 6, 4},
			expected: 5,
		},
		{
			input:    []int{0},
			expected: 0,
		},
		{
			input:    []int{1, 2},
			expected: 1,
		},
		{
			input:    []int{1, 6, 5, 4, 3, 2, 1},
			expected: 1,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, findPeakElement(testcase.input))
	}
}
