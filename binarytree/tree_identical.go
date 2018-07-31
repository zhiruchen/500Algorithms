package binarytree

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// IsIdentical if two binary tree is isdentical
func IsIdentical(x, y *types.BinTreeNode) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil || x != nil && y == nil {
		return false
	}

	if x.Data != y.Data {
		return false
	}

	return IsIdentical(x.Left, y.Left) && IsIdentical(x.Right, y.Right)
}
