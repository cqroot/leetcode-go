package P0413_arithmetic_slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
	assert.Equal(t, 0, numberOfArithmeticSlices([]int{1}))
	assert.Equal(t, 0, numberOfArithmeticSlices([]int{1, 2}))
	assert.Equal(t, 1, numberOfArithmeticSlices([]int{1, 2, 3}))
	assert.Equal(t, 3, numberOfArithmeticSlices([]int{1, 2, 3, 4}))
}
