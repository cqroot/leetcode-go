package p0005longestpalindromicsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("bb"))
}
