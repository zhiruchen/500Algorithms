package bst

import (
	"math"

	"github.com/zhiruchen/500Algorithms/types"
)

// DetermineIfGivenBinaryTreeIsABSTOrNot http://www.techiedelight.com/determine-given-binary-tree-is-a-bst-or-not/
func DetermineIfGivenBinaryTreeIsABSTOrNot(node *types.BinTreeNode) bool {
	return isBST(node, math.MinInt32, math.MaxInt32)
}

func isBST(node *types.BinTreeNode, minVal, maxVal int32) bool {
	if node == nil {
		return true
	}

	if node.Data < minVal || node.Data > maxVal {
		return false
	}

	return isBST(node.Left, minVal, node.Data) && isBST(node.Right, node.Data, maxVal)
}
