package binarytree

import (
	"fmt"

	"github.com/zhiruchen/500Algorithms/types"
)

func rootToLeafSum(root *types.BinTreeNode) int32 {
	if root == nil {
		return 0
	}

	leftSum := rootToLeafSum(root.Left)
	rightSum := rootToLeafSum(root.Right)

	if leftSum >= rightSum {
		return leftSum + root.Data
	}

	return rightSum + root.Data
}

func printPath(root *types.BinTreeNode, sum int32) bool {
	if root == nil {
		return false
	}

	if sum == 0 {
		return true
	}

	left := printPath(root.Left, sum-root.Data)
	right := printPath(root.Right, sum-root.Data)

	if left || right {
		fmt.Println(root.Data, " ")
	}

	return left || right
}

// FindMaxSumPath find max sum path root to leaf
func FindMaxSumPath(root *types.BinTreeNode) {
	sum := rootToLeafSum(root)
	printPath(root, sum)
}
