package P0918_MaximumSumCircularSubarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubarraySumCircular(t *testing.T) {
	assert.Equal(t, 3, maxSubarraySumCircular([]int{1, -2, 3, -2}))
	assert.Equal(t, 10, maxSubarraySumCircular([]int{5, -3, 5}))
	assert.Equal(t, -2, maxSubarraySumCircular([]int{-3, -2, -3}))
}
