package solution

// Iterative
// Time complexity  : O(len(list1) + len(list2))
// Space complexity : O(1)
func mergeTwoLists_Iterative(list1 *ListNode, list2 *ListNode) *ListNode {
	var result, current, node *ListNode

	for list1 != nil && list2 != nil {
		node = new(ListNode)

		if list1.Val < list2.Val {
			node.Val = list1.Val
			list1 = list1.Next
		} else {
			node.Val = list2.Val
			list2 = list2.Next
		}

		if result == nil {
			result = node
			current = result
		} else {
			current.Next = node
			current = current.Next
		}
	}

	if list2 != nil {
		list1 = list2
	}
	for list1 != nil {
		node = new(ListNode)
		node.Val = list1.Val

		list1 = list1.Next

		if result == nil {
			result = node
			current = result
		} else {
			current.Next = node
			current = current.Next
		}
	}

	return result
}
