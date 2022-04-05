package P0673_NumberOfLongestIncreasingSubsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumberOfLIS(t *testing.T) {
	assert.Equal(t, 2, findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	assert.Equal(t, 5, findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	assert.Equal(t, 3, findNumberOfLIS([]int{1, 2, 4, 3, 5, 4}))
	assert.Equal(t, 3, findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
}
