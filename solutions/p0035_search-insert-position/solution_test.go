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

func TestSearchInsert(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			expected: 2,
		},
		{
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			expected: 1,
		},
		{
			input: input{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			expected: 4,
		},
		{
			input: input{
				nums:   []int{1},
				target: 0,
			},
			expected: 0,
		},
		{
			input: input{
				nums:   []int{1},
				target: 1,
			},
			expected: 0,
		},
		{
			input: input{
				nums:   []int{1},
				target: 2,
			},
			expected: 1,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, searchInsert(testcase.input.nums, testcase.input.target))
	}
}
