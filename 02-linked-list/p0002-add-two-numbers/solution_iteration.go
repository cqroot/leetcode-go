// Iteration
// Time complexity  : O(max(len(l1), len(l2)))
// Space complexity : O(1)

package solution

func newNode(result, current *ListNode, val int) (*ListNode, *ListNode) {
	if result == nil {
		result = &ListNode{
			Val: val,
		}
		current = result
	} else {
		node := &ListNode{
			Val: val,
		}
		current.Next = node
		current = current.Next
	}
	return result, current
}

func addTwoNumbers_Iteration(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var result, current *ListNode
	carry := 0

	for l1 != nil && l2 != nil {
		t := l1.Val + l2.Val + carry
		carry = t / 10
		t = t % 10

		result, current = newNode(result, current, t)

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil {
		l1 = l2
	}

	for l1 != nil {
		t := l1.Val + carry
		carry = t / 10
		t = t % 10

		result, current = newNode(result, current, t)

		l1 = l1.Next
	}

	if carry != 0 {
		result, _ = newNode(result, current, carry)
	}

	return result
}
