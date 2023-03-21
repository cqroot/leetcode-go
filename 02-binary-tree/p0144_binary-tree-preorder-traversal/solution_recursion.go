// Recursion
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func preorderTraversal_Recursion(root *TreeNode) []int {
	vals := make([]int, 0, 100)

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)

	return vals
}
