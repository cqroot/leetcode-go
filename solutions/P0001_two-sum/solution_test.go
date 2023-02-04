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

func TestTwoSum(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			input: input{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
		{
			input: input{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
		{
			input: input{
				nums:   []int{3, 2, 3},
				target: 6,
			},
			expected: []int{0, 2},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, twoSum_BruteForce(testcase.input.nums, testcase.input.target))
		require.Equal(t, testcase.expected, twoSum_HashTable(testcase.input.nums, testcase.input.target))
	}
}
