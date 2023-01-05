package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input = []int

type testcase struct {
	input    input
	expected []int
}

func TestMoveZeroes(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			input:    []int{0},
			expected: []int{0},
		},
		{
			input:    []int{1, 0, 1},
			expected: []int{1, 1, 0},
		},
	}

	for _, testcase := range testcases {
		moveZeroes_TwoPointers(testcase.input)
		assert.Equal(t, testcase.expected, testcase.input)
	}
}
