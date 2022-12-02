package solution

import (
	"testing"

	"github.com/cqroot/leetcode-go/datastructure"
	"github.com/stretchr/testify/assert"
)

type input struct {
	list1 []int
	list2 []int
}

type testcase struct {
	input    input
	expected []int
}

func TestMergeTwoLists(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				list1: []int{1, 2, 4},
				list2: []int{1, 3, 4},
			},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input: input{
				list1: []int{},
				list2: []int{},
			},
			expected: []int{},
		},
		{
			input: input{
				list1: []int{0},
				list2: []int{},
			},
			expected: []int{0},
		},
	}

	for _, testcase := range testcases {
		assert.Equal(
			t,
			datastructure.MakeList(testcase.expected),
			mergeTwoLists_Recursive(
				datastructure.MakeList(testcase.input.list1),
				datastructure.MakeList(testcase.input.list2),
			),
			mergeTwoLists_Iterative(
				datastructure.MakeList(testcase.input.list1),
				datastructure.MakeList(testcase.input.list2),
			),
		)
	}
}
