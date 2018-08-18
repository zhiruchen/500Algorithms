package bst

import (
	"math"
	"github.com/zhiruchen/500Algorithms/types"
)

// FloorAndCeilInABinarySearchTree http://www.techiedelight.com/floor-ceil-bst-iterative-recursive/
func FloorAndCeilInABinarySearchTree(root *types.BinTreeNode, key int32) (int32, int32) {
	var min,max int32 = math.MinInt32, math.MaxInt32
	var floor, ceil = &min, &max
	findFloorAndCeilRecurInBST(root, floor, ceil, key)
	return *floor, *ceil
}

func findFloorAndCeilRecurInBST(root *types.BinTreeNode, floor, ceil *int32, key int32) {
	if root == nil {
		return
	}

	// 如果找到了这个key, 那么floor和ceil都是这个key
	if key == root.Data {
		*ceil = root.Data
		*floor = root.Data
		return
	}

	// 如果key小于root，则在左子树寻找floor
	if key < root.Data {
		*ceil = root.Data
		findFloorAndCeilRecurInBST(root.Left, floor, ceil, key)
		return
	}

	// 如果key大于root, 则在右子树寻找ceil
	*floor = root.Data
	findFloorAndCeilRecurInBST(root.Right, floor, ceil, key)
}
