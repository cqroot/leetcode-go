package P0354_RussianDollEnvelopes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	assert.Equal(t, 3, maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	assert.Equal(t, 1, maxEnvelopes([][]int{{1, 1}, {1, 1}, {1, 1}}))
}
