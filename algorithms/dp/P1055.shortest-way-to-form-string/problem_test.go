package P1055_shortest_way_to_form_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestWay(t *testing.T) {
	assert.Equal(t, 2, shortestWay("abc", "abcbc"))
	assert.Equal(t, -1, shortestWay("abc", "acdbc"))
	assert.Equal(t, 3, shortestWay("xyz", "xzyxz"))
}
