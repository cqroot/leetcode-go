// Stack
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func isPalindrome_Stack1(head *ListNode) bool {
	fast, slow := head, head
	lenStack := 0
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		lenStack += 1
	}
	if fast != nil {
		slow = slow.Next
	}

	stack := make([]int, lenStack)
	for i := 0; slow != nil; i++ {
		stack[i] = slow.Val
		slow = slow.Next
	}

	slow = head
	for i := lenStack - 1; i >= 0; i-- {
		if stack[i] != slow.Val {
			return false
		}
		slow = slow.Next
	}

	return true
}
