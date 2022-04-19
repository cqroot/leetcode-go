package P0091_decode_ways

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	assert.Equal(t, 2, numDecodings("12"))
	assert.Equal(t, 3, numDecodings("226"))
	assert.Equal(t, 0, numDecodings("0"))
	assert.Equal(t, 2, numDecodings("11106"))
	assert.Equal(t, 1, numDecodings("27"))
	assert.Equal(t, 1836311903, numDecodings("111111111111111111111111111111111111111111111"))
}
