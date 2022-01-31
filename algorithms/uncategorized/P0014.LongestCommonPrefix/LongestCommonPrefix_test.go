package p0014longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
	assert.Equal(t, longestCommonPrefix([]string{"dog", "racecar", "car"}), "")
	assert.Equal(t, longestCommonPrefix([]string{"dog"}), "dog")
}
