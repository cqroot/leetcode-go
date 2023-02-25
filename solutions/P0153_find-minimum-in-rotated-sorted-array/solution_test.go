package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    []int
	expected int
}

func TestFindMin(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			input:    []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			input:    []int{11, 13, 15, 17},
			expected: 11,
		},
		{
			input:    []int{1},
			expected: 1,
		},
		{
			input:    []int{1, 2},
			expected: 1,
		},
		{
			input:    []int{2, 1},
			expected: 1,
		},
		{
			input:    []int{1, 2, 3},
			expected: 1,
		},
		{
			input:    []int{2, 3, 1},
			expected: 1,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, findMin(testcase.input), "input: %+v", testcase.input)
	}
}
