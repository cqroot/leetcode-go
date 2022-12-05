package solution

import (
	"testing"

	"github.com/cqroot/leetcode-go/datastructure"
	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    *ListNode
	expected *ListNode
}

func TestReverseList(t *testing.T) {
	l1A := datastructure.MakeList([]int{1, 2, 3, 4, 5})
	l1B := l1A
	l1B.Next.Next.Next.Next.Next = l1B.Next.Next.Next
	l1B.Next.Next.Next.Next = l1B.Next.Next
	l1B.Next.Next.Next = l1B.Next
	l1B.Next.Next = l1B
	l1B.Next = nil

	l2A := datastructure.MakeList([]int{1, 2})
	l2B := l2A
	l2B.Next.Next = l2B
	l2B.Next = nil

	l3A := datastructure.MakeList([]int{})
	l3B := l3A

	testcases := []testcase{
		{
			input:    l1A,
			expected: l1B,
		},
		{
			input:    l2A,
			expected: l2B,
		},
		{
			input:    l3A,
			expected: l3B,
		},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, reverseList_Iteration(testcase.input))
		assert.Equal(t, testcase.expected, reverseList_Recursion(testcase.input))
	}
}
