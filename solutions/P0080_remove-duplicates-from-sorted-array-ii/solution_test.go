package solution

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type input = []int

type expected struct {
	nums []int
	val  int
}

type testcase struct {
	input    input
	expected expected
}

func TestRemoveDuplicates(t *testing.T) {
	testcases := []testcase{
		{
			input: []int{1, 1, 1, 2, 2, 3},
			expected: expected{
				nums: []int{1, 1, 2, 2, 3},
				val:  5,
			},
		},
		{
			input: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected: expected{
				nums: []int{0, 0, 1, 1, 2, 3, 3},
				val:  7,
			},
		},
		{
			input: []int{1},
			expected: expected{
				nums: []int{1},
				val:  1,
			},
		},
	}

	for _, testcase := range testcases {
		fTest := func(actualVal int, actualNums []int) {
			expectedNums := testcase.expected.nums

			sort.Slice(actualNums[:actualVal], func(i, j int) bool {
				return actualNums[i] < actualNums[j]
			})
			sort.Slice(expectedNums, func(i, j int) bool {
				return expectedNums[i] < expectedNums[j]
			})

			require.Equal(t, testcase.expected.val, actualVal)
			for i := 0; i < testcase.expected.val; i++ {
				require.Equal(t, expectedNums[i], actualNums[i])
			}
		}

		actualNums := make([]int, len(testcase.input))

		copy(actualNums, testcase.input)
		fTest(
			removeDuplicates_TwoPointers(actualNums),
			actualNums,
		)
	}
}
