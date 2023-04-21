// Recursion
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func isPalindrome_Recursion(head *ListNode) bool {
	curr := head

	var fCheck func(head *ListNode) bool
	fCheck = func(head *ListNode) bool {
		if head == nil {
			return true
		}

		if !fCheck(head.Next) {
			return false
		}

		if head.Val == curr.Val {
			curr = curr.Next
			return true
		} else {
			return false
		}
	}

	return fCheck(head)
}
