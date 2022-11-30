package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum_BruteForce([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum_BruteForce([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum_BruteForce([]int{3, 3}, 6))
	assert.Equal(t, []int{0, 2}, twoSum_BruteForce([]int{3, 2, 3}, 6))

	assert.Equal(t, []int{0, 1}, twoSum_HashTable([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum_HashTable([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum_HashTable([]int{3, 3}, 6))
	assert.Equal(t, []int{0, 2}, twoSum_HashTable([]int{3, 2, 3}, 6))
}
