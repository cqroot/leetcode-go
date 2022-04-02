package p0021mergetwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeListNode(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	node := new(ListNode)

	node.Val = input[0]
	if len(input) > 1 {
		node.Next = makeListNode(input[1:])
	}
	return node
}

func TestMergeTwoLists(t *testing.T) {
	assert.Equal(
		t, makeListNode([]int{1, 1, 2, 3, 4, 4}),
		mergeTwoLists(makeListNode([]int{1, 2, 4}), makeListNode([]int{1, 3, 4})),
	)
	assert.Equal(
		t, makeListNode([]int{}),
		mergeTwoLists(makeListNode([]int{}), makeListNode([]int{})),
	)
	assert.Equal(
		t, makeListNode([]int{0}),
		mergeTwoLists(makeListNode([]int{}), makeListNode([]int{0})),
	)
}
