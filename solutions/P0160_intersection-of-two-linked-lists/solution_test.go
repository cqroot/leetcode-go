package solution

import (
	"testing"

	"github.com/cqroot/leetcode-go/datastructure"
	"github.com/stretchr/testify/assert"
)

type input struct {
	a *ListNode
	b *ListNode
}

type testcase struct {
	input    input
	expected *ListNode
}

func TestGetIntersectionNode(t *testing.T) {
	l1A := datastructure.MakeList([]int{4, 1, 8, 4, 5})
	l1B := datastructure.MakeList([]int{5, 6, 1})
	l1B.Next.Next.Next = l1A.Next.Next

	l2A := datastructure.MakeList([]int{1, 9, 1, 2, 4})
	l2B := datastructure.MakeList([]int{3})
	l2B.Next = l2A.Next.Next.Next

	l3A := datastructure.MakeList([]int{2, 6, 4})
	l3B := datastructure.MakeList([]int{1, 5})

	l4A := datastructure.MakeList([]int{3})
	l4B := datastructure.MakeList([]int{2})
	l4B.Next = l4A

	testcases := []testcase{
		{
			input: input{
				a: l1A,
				b: l1B,
			},
			expected: l1A.Next.Next,
		},
		{
			input: input{
				a: l2A,
				b: l2B,
			},
			expected: l2A.Next.Next.Next,
		},
		{
			input: input{
				a: l3A,
				b: l3B,
			},
			expected: nil,
		},
		{
			input: input{
				a: l4A,
				b: l4B,
			},
			expected: l4A,
		},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, getIntersectionNode_TwoPointers(testcase.input.a, testcase.input.b))
		assert.Equal(t, testcase.expected, getIntersectionNode_HashTable(testcase.input.a, testcase.input.b))
	}
}
