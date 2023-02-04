package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    []int
	expected int
}

func TestSingleNumber(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{2, 2, 1},
			expected: 1,
		},
		{
			input:    []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			input:    []int{1},
			expected: 1,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, singleNumber(testcase.input))
	}
}
