package solution

import (
	"testing"

	"github.com/cqroot/leetcode-go/datastructure"
	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    *ListNode
	expected bool
}

func TestHasCycle(t *testing.T) {
	l1 := datastructure.MakeList([]int{3, 2, 0, -4})
	l1.Next.Next.Next.Next = l1.Next

	l2 := datastructure.MakeList([]int{1, 2})
	l2.Next.Next = l2

	testcases := []testcase{
		{
			input:    l1,
			expected: true,
		},
		{
			input:    l2,
			expected: true,
		},
		{
			input:    datastructure.MakeList([]int{1}),
			expected: false,
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected, hasCycle(testcase.input))
	}
}
