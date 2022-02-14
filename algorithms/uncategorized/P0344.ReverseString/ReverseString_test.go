package P0344_ReverseString

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	var s []byte

	s = []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	assert.Equal(t, []byte{'o', 'l', 'l', 'e', 'h'}, s)

	s = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	assert.Equal(t, []byte{'h', 'a', 'n', 'n', 'a', 'H'}, s)
}
