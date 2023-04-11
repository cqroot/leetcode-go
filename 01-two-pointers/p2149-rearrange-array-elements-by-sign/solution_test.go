package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    []int
	expected []int
}

func TestRearrangeArray(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{3, 1, -2, -5, 2, -4},
			expected: []int{3, -2, 1, -5, 2, -4},
		},
		{
			input:    []int{-1, 1},
			expected: []int{1, -1},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, rearrangeArray(testcase.input))
	}
}
