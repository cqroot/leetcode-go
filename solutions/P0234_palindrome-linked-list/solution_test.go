package solution

import (
	"testing"

	"github.com/cqroot/leetcode-go/datastructure"
	"github.com/stretchr/testify/assert"
)

type input *ListNode

type testcase struct {
	input    input
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	testcases := []testcase{
		{
			input:    datastructure.MakeList([]int{1, 2, 2, 1}),
			expected: true,
		},
		{
			input:    datastructure.MakeList([]int{1, 2}),
			expected: false,
		},
		{
			input:    datastructure.MakeList([]int{1}),
			expected: true,
		},
		{
			input:    datastructure.MakeList([]int{1, 0, 0}),
			expected: false,
		},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, isPalindrome_Stack1(testcase.input))
		assert.Equal(t, testcase.expected, isPalindrome_Stack2(testcase.input))
		assert.Equal(t, testcase.expected, isPalindrome_Recursion(testcase.input))
	}
}
