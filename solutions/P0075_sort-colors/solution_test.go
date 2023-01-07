package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    []int
	expected []int
}

func TestSortColors(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			input:    []int{0},
			expected: []int{0},
		},
	}

	for _, testcase := range testcases {
		actualNums := make([]int, len(testcase.input))
		copy(actualNums, testcase.input)

		sortColors_TwoPointers(actualNums)
		assert.Equal(t, testcase.expected, actualNums)
	}
}
