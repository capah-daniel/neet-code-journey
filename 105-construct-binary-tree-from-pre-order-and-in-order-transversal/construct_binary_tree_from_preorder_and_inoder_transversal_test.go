package constructbinarytreefrompreorderandinordertransversal

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{name: "test 1", preorder: []int{3, 9, 20, 15, 7}, inorder: []int{9, 3, 15, 20, 7}, want: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}},
		{name: "test 2", preorder: []int{-1}, inorder: []int{-1}, want: &TreeNode{Val: -1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := BuildTree(test.preorder, test.inorder)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("BuildTree(%v, %v) = %v, want %v", test.preorder, test.inorder, got, test.want)
			}
		})
	}
}
