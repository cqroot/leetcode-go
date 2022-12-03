// Hash Table
// Time complexity  : O(len(listA) + len(listB))
// Space complexity : O(len(listA))

package solution

func getIntersectionNode_HashTable(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]int)

	for headA != nil {
		nodeMap[headA] = 1
		headA = headA.Next
	}

	for headB != nil {
		if nodeMap[headB] == 1 {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
