package p0002addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)

	tmpVal := 0
	tmpNode := result
	l1Val := 0
	l2Val := 0

	for true {
		if l1 == nil {
			l1Val = 0
		} else {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			l2Val = 0
		} else {
			l2Val = l2.Val
			l2 = l2.Next
		}

		tmpVal = l1Val + l2Val + tmpVal
		if tmpVal >= 10 {
			tmpNode.Val = tmpVal % 10
			tmpVal = (tmpVal - tmpNode.Val) / 10
		} else {
			tmpNode.Val = tmpVal
			tmpVal = 0
		}

		if l1 != nil || l2 != nil || tmpVal != 0 {
			nextNode := new(ListNode)
			tmpNode.Next = nextNode
			tmpNode = nextNode
		} else {
			break
		}
	}

	return result
}
