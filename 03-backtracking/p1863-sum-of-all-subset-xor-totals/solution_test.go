package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    []int
	expected int
}

func TestSubsetXORSum(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{1, 3},
			expected: 6,
		},
		{
			input:    []int{5, 1, 6},
			expected: 28,
		},
		{
			input:    []int{3, 4, 5, 6, 7, 8},
			expected: 480,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, subsetXORSum(testcase.input))
	}
}
