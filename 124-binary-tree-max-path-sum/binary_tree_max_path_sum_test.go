package binarytreemaxpathsum

import "testing"

func TestBinaryTreeMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{name: "test 1", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxPathSum(test.root)
			if got != test.want {
				t.Errorf("MaxPathSum(%v) = %d, want %d", test.root, got, test.want)
			}
		})
	}
}
