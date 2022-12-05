// Recursion
// Time complexity  : O(len(list))
// Space complexity : O(len(list))

package solution

func reverseList_Recursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := reverseList_Recursion(head.Next)

	head.Next.Next = head
	head.Next = nil

	return result
}
