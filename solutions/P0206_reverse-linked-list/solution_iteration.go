// Iteration
// Time complexity  : O(len(list))
// Space complexity : O(1)

package solution

func reverseList_Iteration(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev, next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
