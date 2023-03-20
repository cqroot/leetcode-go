// Recursion
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func preorderTraversal_Iteration(root *TreeNode) []int {
	vals := make([]int, 0, 100)
	stack := []*TreeNode{}
	node := root

	for node != nil || len(stack) != 0 {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return vals
}
