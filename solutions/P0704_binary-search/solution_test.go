package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			expected: 4,
		},
		{
			input: input{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			expected: -1,
		},
		{
			input: input{
				nums:   []int{0, 1, 2},
				target: 2,
			},
			expected: 2,
		},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, search(testcase.input.nums, testcase.input.target))
	}
}
