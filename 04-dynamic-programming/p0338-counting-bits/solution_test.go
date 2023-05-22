package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    int
	expected []int
}

func TestCountBits(t *testing.T) {
	testcases := []testcase{
		{
			input:    2,
			expected: []int{0, 1, 1},
		},
		{
			input:    5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			countBits(testcase.input),
			"Testcase: %+v", testcase.input)
	}
}
