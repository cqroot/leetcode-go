// Recursion
// Time complexity  : O(n)
// Space complexity : O(n)

package solution

func postorderTraversal_Recursion(root *TreeNode) []int {
	vals := make([]int, 0, 100)

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		vals = append(vals, node.Val)
	}

	traversal(root)
	return vals
}
