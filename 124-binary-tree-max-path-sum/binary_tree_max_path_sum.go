package binarytreemaxpathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var maxValue int

	maxValue = -1000
	helper(root, &maxValue)
	return maxValue
}

func helper(root *TreeNode, maxValue *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, maxValue)
	right := helper(root.Right, maxValue)

	left = max(left, 0)
	right = max(right, 0)

	*maxValue = max(*maxValue, root.Val+left+right)

	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
