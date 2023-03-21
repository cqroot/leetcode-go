// Two Pointers
// Time complexity  : O(len(listA) + len(listB))
// Space complexity : O(1)

package solution

func getIntersectionNode_TwoPointers(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB

	for pA != pB {
		pA = pA.Next
		pB = pB.Next

		if pA == nil && pB != nil {
			pA = headB
		}
		if pB == nil && pA != nil {
			pB = headA
		}
	}

	return pA
}
