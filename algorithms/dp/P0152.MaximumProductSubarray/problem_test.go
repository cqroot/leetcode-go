package P0152_MaximumProductSubarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	assert.Equal(t, -1, maxProduct([]int{-1}))
	assert.Equal(t, 6, maxProduct([]int{2, 3, -2, 4}))
	assert.Equal(t, 0, maxProduct([]int{-2, 0, -1}))
	assert.Equal(t, 2, maxProduct([]int{0, 2}))
	assert.Equal(t, 24, maxProduct([]int{-2, 3, -4}))
}
