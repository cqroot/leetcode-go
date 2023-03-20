package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    *TreeNode
	expected []int
}

func TestInorderTraversal(t *testing.T) {
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
			expected: []int{1, 3, 2},
		},
		{
			input:    nil,
			expected: []int{},
		},
		{
			input: &TreeNode{
				Val: 1,
			},
			expected: []int{1},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, inorderTraversal_Recursion(testcase.input))
		require.Equal(t, testcase.expected, inorderTraversal_Iteration(testcase.input))
	}
}
