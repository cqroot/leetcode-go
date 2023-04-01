// Iteration
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func preorderTraversal_Iteration(root *TreeNode) []int {
	vals := make([]int, 0, 100)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			vals = append(vals, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return vals
}
