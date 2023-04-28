package solution

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    int
	expected [][]int
}

func TestGenerate(t *testing.T) {
	testcases := []testcase{
		{
			input:    5,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			input:    1,
			expected: [][]int{{1}},
		},
	}

	for _, testcase := range testcases {
		require.Equal(t, testcase.expected,
			generate(testcase.input),
			"Testcase: %+v", testcase.input)
	}
}
