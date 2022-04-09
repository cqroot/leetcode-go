package P0368_largest_divisible_subset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	//assert.Equal(t, []int{1, 2}, largestDivisibleSubset([]int{1, 2, 3}))
	//assert.Equal(t, []int{1, 2, 4, 8}, largestDivisibleSubset([]int{1, 2, 4, 8}))
	//assert.Equal(t, []int{1, 2, 4, 8, 16}, largestDivisibleSubset([]int{1, 2, 4, 8, 16}))
	assert.Equal(t, []int{4, 8, 16}, largestDivisibleSubset([]int{3, 4, 16, 8}))
}
