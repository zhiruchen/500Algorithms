package bst

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// SearchGivenKeyInBST http://www.techiedelight.com/search-given-key-in-bst/
func SearchGivenKeyInBST(node *types.BinTreeNode, key int32) *types.BinTreeNode {
	if node == nil {
		return nil
	}

	if node.Data == key {
		return node
	}

	if key < node.Data {
		return SearchGivenKeyInBST(node.Left, key)
	}

	return SearchGivenKeyInBST(node.Right, key)
}
