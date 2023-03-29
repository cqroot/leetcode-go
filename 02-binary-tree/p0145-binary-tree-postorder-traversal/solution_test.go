package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    *TreeNode
	expected []int
}

func TestAdd(t *testing.T) {
	testcases := []testcase{
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{3, 2, 1},
		},
		{
			input:    nil,
			expected: []int{},
		},
		{
			input:    &TreeNode{Val: 1},
			expected: []int{1},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, postorderTraversal_Recursion(testcase.input))
	}
}
