package P1027_longest_arithmetic_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestArithSeqLength(t *testing.T) {
	//assert.Equal(t, 4, longestArithSeqLength([]int{3, 6, 9, 12}))
	//assert.Equal(t, 3, longestArithSeqLength([]int{1, 3, 7, 11, 12, 14, 18}))
	assert.Equal(t, 3, longestArithSeqLength([]int{9, 4, 7, 2, 10}))
}
