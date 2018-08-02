package binarytree

import (
	list "github.com/emirpasic/gods/lists/arraylist"

	"github.com/zhiruchen/500Algorithms/types"
)

// IsComplete every level except last is completely filled
// http://www.techiedelight.com/check-given-binary-tree-complete-binary-tree-not/
func IsComplete(root *types.BinTreeNode) bool {
	if root == nil {
		return true
	}

	queue := list.New()
	queue.Add(root)

	var isNonFull bool
	for !queue.Empty() {
		v, _ := queue.Get(0)
		front := v.(*types.BinTreeNode)
		queue.Remove(0)

		// 前面的遍历过程中遇到了不完全的节点，而且当前节点不是叶子节点，则二叉树不完全
		if isNonFull && (front.Left != nil || front.Right != nil) {
			return false
		}

		if front.Left == nil && front.Right != nil {
			return false
		}

		if front.Left != nil {
			queue.Add(front.Left)
		} else {
			isNonFull = true
		}

		if front.Right != nil {
			queue.Add(front.Right)
		} else {
			isNonFull = true
		}
	}

	return true
}
