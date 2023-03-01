package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type input struct {
	nums   []int
	target int
}

type testcase struct {
	input    input
	expected []int
}

func TestSearchRange(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			expected: []int{3, 4},
		},
		{
			input: input{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			expected: []int{-1, -1},
		},
		{
			input: input{
				nums:   []int{},
				target: 0,
			},
			expected: []int{-1, -1},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, searchRange(testcase.input.nums, testcase.input.target))
	}
}
