package solution

// Recursion
// Time complexity  : O(max(len(l1), len(l2)))
// Space complexity : O(max(len(l1), len(l2)))
func addTwoNumbers_Recursion(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var t int

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		if l1.Val < 10 {
			return l1
		} else {
			l2 = &ListNode{
				Val: 0,
			}
		}
	}

	t = l1.Val + l2.Val
	if t > 9 {
		if l1.Next == nil {
			l1.Next = &ListNode{
				Val:  t / 10,
				Next: nil,
			}
		} else {
			l1.Next.Val = l1.Next.Val + t/10
		}
	}
	result = &ListNode{
		Val:  t % 10,
		Next: addTwoNumbers_Recursion(l1.Next, l2.Next),
	}

	return result
}
