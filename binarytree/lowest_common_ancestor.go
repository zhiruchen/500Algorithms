package binarytree

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// LowestCommonAncestor lowest node in binary tree that has both x and y as descendants
// http://www.techiedelight.com/find-lowest-common-ancestor-lca-two-nodes-binary-tree/
func LowestCommonAncestor(root, x, y *types.BinTreeNode) *types.BinTreeNode {
	if root == nil {
		return nil
	}

	if root.Data == x.Data || root.Data == y.Data {
		return root
	}

	left := LowestCommonAncestor(root.Left, x, y)
	right := LowestCommonAncestor(root.Right, x, y)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
