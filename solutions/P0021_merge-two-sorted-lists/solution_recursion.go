// Recursion
// Time complexity  : O(len(list1) + len(list2))
// Space complexity : O(len(list1) + len(list2))

package solution

func mergeTwoLists_Recursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists_Recursion(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists_Recursion(list1, list2.Next)
		return list2
	}
}
