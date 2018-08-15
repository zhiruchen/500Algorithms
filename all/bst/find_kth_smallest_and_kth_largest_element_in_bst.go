package bst

import (
	"math"

	"github.com/zhiruchen/500Algorithms/types"
)

// FindKthSmallestAndKthLargestElementInBST http://www.techiedelight.com/find-kth-smallest-largest-element-bst/
func FindKthSmallestAndKthLargestElementInBST(node *types.BinTreeNode, k int) (small, large int32) {
	start := 0
	start1 := 0
	return findKthSmallestNode(node, &start, k), findKthLargestNode(node, &start1, k)
}

func findKthSmallestNode(node *types.BinTreeNode, n *int, k int) int32 {
	if node == nil {
		return math.MaxInt32
	}

	left := findKthSmallestNode(node.Left, n, k)
	if left != math.MaxInt32 {
		return left
	}

	*n = *n + 1
	if *n == k {
		return node.Data
	}

	return findKthSmallestNode(node.Right, n, k)
}

func findKthLargestNode(node *types.BinTreeNode, n *int, k int) int32 {
	if node == nil {
		return math.MaxInt32
	}

	right := findKthLargestNode(node.Right, n, k)
	if right != math.MaxInt32 {
		return right
	}

	*n = *n + 1
	if *n == k {
		return node.Data
	}

	return findKthLargestNode(node.Left, n, k)
}
