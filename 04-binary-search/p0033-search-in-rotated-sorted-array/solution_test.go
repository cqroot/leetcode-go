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
	expected int
}

func TestSearch(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			expected: 4,
		},
		{
			input: input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			expected: -1,
		},
		{
			input: input{
				nums:   []int{1},
				target: 0,
			},
			expected: -1,
		},
		{
			input: input{
				nums:   []int{5, 1, 3},
				target: 5,
			},
			expected: 0,
		},
		{
			input: input{
				nums:   []int{4, 5, 6, 7, 8, 1, 2, 3},
				target: 8,
			},
			expected: 4,
		},
		{
			input: input{
				nums:   []int{3, 1},
				target: 1,
			},
			expected: 1,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, search(testcase.input.nums, testcase.input.target))
	}
}
