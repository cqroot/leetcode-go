package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    int
	expected []string
}

func TestGenerateParenthesis(t *testing.T) {
	testcases := []testcase{
		{
			input:    3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			input:    1,
			expected: []string{"()"},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			generateParenthesis(testcase.input),
			"Testcase: %+v", testcase.input)
	}
}
