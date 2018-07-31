package binarytree

import (
	"github.com/zhiruchen/500Algorithms/common"
	"github.com/zhiruchen/500Algorithms/types"
)

// TreeHeight Binary tree height
func TreeHeight(root *types.BinTreeNode) int32 {
	if root == nil {
		return 0
	}

	return common.Max(TreeHeight(root.Left), TreeHeight(root.Right)) + 1
}
