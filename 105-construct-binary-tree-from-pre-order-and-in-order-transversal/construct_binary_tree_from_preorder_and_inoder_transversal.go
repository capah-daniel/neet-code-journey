package constructbinarytreefrompreorderandinordertransversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {

	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}

	var helper func(preStart, preEnd, inStart, inEnd int) *TreeNode

	helper = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}

		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}

		mid := inorderMap[rootVal]
		leftSize := mid - inStart

		root.Left = helper(preStart+1, preStart+leftSize, inStart, mid-1)
		root.Right = helper(preStart+leftSize+1, preEnd, mid+1, inEnd)

		return root
	}

	return helper(0, len(preorder)-1, 0, len(inorder)-1)
}
