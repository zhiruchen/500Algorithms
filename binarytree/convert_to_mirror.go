package binarytree

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// ConvertToMirror convert binary tree to it's mirror
func ConvertToMirror(root *types.BinTreeNode) {
	if root == nil {
		return
	}

	ConvertToMirror(root.Left)
	ConvertToMirror(root.Right)

	root.Left, root.Right = root.Right, root.Left
}
