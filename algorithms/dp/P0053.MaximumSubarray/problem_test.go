package P0053_MaximumSubarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 1, maxSubArray([]int{1}))
	assert.Equal(t, 23, maxSubArray([]int{5, 4, -1, 7, 8}))
}
