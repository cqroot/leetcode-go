// Tow pointers
// Time complexity  : O(len(list))
// Space complexity : O(1)

package solution

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head.Next, head

	for fast != slow {
		if fast == nil || fast.Next == nil || slow == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}
