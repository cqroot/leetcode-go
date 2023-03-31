package solution

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type input struct {
	n int
	k int
}

type testcase struct {
	input    input
	expected [][]int
}

func TestCombine(t *testing.T) {
	testcases := []testcase{
		// {
		// 	input: input{
		// 		n: 4,
		// 		k: 2,
		// 	},
		// 	expected: [][]int{
		// 		{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
		// 	},
		// },
		// {
		// 	input: input{
		// 		n: 1,
		// 		k: 1,
		// 	},
		// 	expected: [][]int{{1}},
		// },
		{
			input: input{
				n: 5,
				k: 4,
			},
			expected: [][]int{
				{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5},
			},
		},
	}

	for _, testcase := range testcases {
		sort.Slice(testcase.expected, func(i, j int) bool {
			for idx := 0; idx < testcase.input.k; idx++ {
				if testcase.expected[i][idx] == testcase.expected[j][idx] {
					continue
				}
				return testcase.expected[i][idx] > testcase.expected[j][idx]
			}
			return true
		})
		actual := combine(testcase.input.n, testcase.input.k)
		sort.Slice(actual, func(i, j int) bool {
			for idx := 0; idx < testcase.input.k; idx++ {
				if testcase.expected[i][idx] == testcase.expected[j][idx] {
					continue
				}
				return testcase.expected[i][idx] > testcase.expected[j][idx]
			}
			return true
		})

		require.Equal(t, testcase.expected, actual)
	}
}
