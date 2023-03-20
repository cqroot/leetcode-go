// Name
// Time complexity  : O(1)
// Space complexity : O(1)

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
