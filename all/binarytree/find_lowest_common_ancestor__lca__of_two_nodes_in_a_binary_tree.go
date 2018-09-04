package binarytree

import "github.com/zhiruchen/500Algorithms/types"

// FindLowestCommonAncestorLCAOfTwoNodesInABinaryTree http://www.techiedelight.com/find-lowest-common-ancestor-lca-two-nodes-binary-tree/
func FindLowestCommonAncestorLCAOfTwoNodesInABinaryTree(root, x, y *types.BinTreeNode) *types.BinTreeNode {
	if root == nil {
		return nil
	}

	if root.Data == x.Data || root.Data == y.Data {
		return root
	}

	left := FindLowestCommonAncestorLCAOfTwoNodesInABinaryTree(root.Left, x, y)
	right := FindLowestCommonAncestorLCAOfTwoNodesInABinaryTree(root.Right, x, y)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
