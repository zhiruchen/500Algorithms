package bst

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// FindInorderSuccessorForGivenKeyInABST http://www.techiedelight.com/find-inorder-successor-given-key-bst/
func FindInorderSuccessorForGivenKeyInABST(node, succ *types.BinTreeNode, key int32) {
	if node == nil {
		return
	}

	if node.Data == key {
		if node.Right != nil {
			succ = findMinNode(node)
		}
		return
	}

	if key < node.Data {
		succ = node
		FindInorderSuccessorForGivenKeyInABST(node.Left, succ, key)
		return
	}

	FindInorderSuccessorForGivenKeyInABST(node.Right, succ, key)
}

func findMinNode(node *types.BinTreeNode) *types.BinTreeNode {
	for node.Left != nil {
		node = node.Left
	}

	return node
}
