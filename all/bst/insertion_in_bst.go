package bst

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// InsertionInBST http://www.techiedelight.com/insertion-in-bst/
func InsertionInBST(node *types.BinTreeNode, key int32) *types.BinTreeNode {
	if node == nil {
		return &types.BinTreeNode{Data: key}
	}

	if key < node.Data {
		node.Left = InsertionInBST(node.Left, key)
	} else {
		node.Right = InsertionInBST(node.Right, key)
	}

	return node
}
