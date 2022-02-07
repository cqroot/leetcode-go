package P0027_RemoveElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	assert.Equal(t, 2, removeElement(nums, 3))
	assert.Equal(t, []int{2, 2}, nums[:2])

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	assert.Equal(t, 5, removeElement(nums, 2))
	assert.Equal(t, []int{0, 1, 4, 0, 3}, nums[:5])

	nums = []int{}
	assert.Equal(t, 0, removeElement(nums, 0))

	nums = []int{0}
	assert.Equal(t, 0, removeElement(nums, 0))
}
