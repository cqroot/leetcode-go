package datastructure

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(nums []int) *ListNode {
	var result, current *ListNode

	for _, num := range nums {
		node := &ListNode{
			Val: num,
		}
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

func PrintList(list *ListNode) {
	for list != nil {
		if list.Next != nil {
			fmt.Printf("%d ", list.Val)
		} else {
			fmt.Println(list.Val)
		}
		list = list.Next
	}
}
