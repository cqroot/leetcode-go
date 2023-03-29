// Iteration
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func inorderTraversal_Iteration(root *TreeNode) []int {
	vals := make([]int, 0, 100)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, root.Val)
		root = root.Right
	}

	return vals
}
