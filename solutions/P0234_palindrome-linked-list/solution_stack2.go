// Stack
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func isPalindrome_Stack2(head *ListNode) bool {
	fast, slow := head, head
	stack := make([]int, 0)

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)

		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] != slow.Val {
			return false
		}
		slow = slow.Next
	}

	return true
}
