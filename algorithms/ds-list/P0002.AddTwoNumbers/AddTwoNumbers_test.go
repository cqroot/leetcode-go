package p0002addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeListNode(input []int) *ListNode {
	node := new(ListNode)
	node.Val = input[0]
	if len(input) > 1 {
		node.Next = makeListNode(input[1:])
	}
	return node
}

func Test_AddTwoNumbers(t *testing.T) {
	assert.Equal(
		t, makeListNode([]int{7, 0, 8}),
		addTwoNumbers(makeListNode([]int{2, 4, 3}), makeListNode([]int{5, 6, 4})),
	)
	assert.Equal(
		t, makeListNode([]int{0}),
		addTwoNumbers(makeListNode([]int{0}), makeListNode([]int{0})),
	)
	assert.Equal(
		t, makeListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		addTwoNumbers(makeListNode([]int{9, 9, 9, 9, 9, 9, 9}), makeListNode([]int{9, 9, 9, 9})),
	)
	assert.Equal(
		t, makeListNode([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		addTwoNumbers(
			makeListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			makeListNode([]int{5, 6, 4}),
		),
	)
}
