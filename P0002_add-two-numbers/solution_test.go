package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cqroot/leetcode-go/datastructure"
)

type input struct {
	list1 []int
	list2 []int
}

type testcase struct {
	input    input
	expected []int
}

func TestAddTwoNumbers(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				list1: []int{2, 4, 3},
				list2: []int{5, 6, 4},
			},
			expected: []int{7, 0, 8},
		},
		{
			input: input{
				list1: []int{0},
				list2: []int{0},
			},
			expected: []int{0},
		},
		{
			input: input{
				list1: []int{9, 9, 9, 9, 9, 9, 9},
				list2: []int{9, 9, 9, 9},
			},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, testcase := range testcases {
		assert.Equal(
			t,
			datastructure.MakeList(testcase.expected),
			addTwoNumbers_Recursion(
				datastructure.MakeList(testcase.input.list1),
				datastructure.MakeList(testcase.input.list2),
			),
		)
		assert.Equal(
			t,
			datastructure.MakeList(testcase.expected),
			addTwoNumbers_Iteration(
				datastructure.MakeList(testcase.input.list1),
				datastructure.MakeList(testcase.input.list2),
			),
		)
	}
}
