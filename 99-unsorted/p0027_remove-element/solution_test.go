package solution

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type input struct {
	nums []int
	val  int
}

type expected struct {
	nums []int
	val  int
}

type testcase struct {
	input    input
	expected expected
}

func TestRemoveElement(t *testing.T) {
	testcases := []testcase{
		{
			input: input{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			expected: expected{
				nums: []int{2, 2},
				val:  2,
			},
		},
		{
			input: input{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			expected: expected{
				nums: []int{0, 1, 4, 0, 3},
				val:  5,
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

		actualNums := make([]int, len(testcase.input.nums))

		copy(actualNums, testcase.input.nums)
		fTest(
			removeElement_TwoPointers1(actualNums, testcase.input.val),
			actualNums,
		)

		copy(actualNums, testcase.input.nums)
		fTest(
			removeElement_TwoPointers2(testcase.input.nums, testcase.input.val),
			testcase.input.nums,
		)
	}
}
