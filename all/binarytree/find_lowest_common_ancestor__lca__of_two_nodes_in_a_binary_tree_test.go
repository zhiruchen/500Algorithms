package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhiruchen/500Algorithms/types"
)

type node = types.BinTreeNode

func TestFindLowestCommonAncestorLCAOfTwoNodesInABinaryTree(t *testing.T) {
	root := &node{
		Data: 1,
		Left: &node{
			Data:  2,
			Left:  &node{Data: 4},
			Right: &node{Data: 5},
		},
		Right: &node{
			Data:  3,
			Left:  &node{Data: 6},
			Right: &node{Data: 7},
		},
	}

	x, y := &node{Data: 2}, &node{Data: 3}
	ancestor := FindLowestCommonAncestorLCAOfTwoNodesInABinaryTree(root, x, y)
	assert.Equal(t, int32(1), ancestor.Data)

	x, y = &node{Data: 4}, &node{Data: 5}
	ancestor = FindLowestCommonAncestorLCAOfTwoNodesInABinaryTree(root, x, y)
	assert.Equal(t, int32(2), ancestor.Data)

	x, y = &node{Data: 2}, &node{Data: 4}
	ancestor = FindLowestCommonAncestorLCAOfTwoNodesInABinaryTree(root, x, y)
	assert.Equal(t, int32(2), ancestor.Data)

	x, y = &node{Data: 4}, &node{Data: 6}
	ancestor = FindLowestCommonAncestorLCAOfTwoNodesInABinaryTree(root, x, y)
	assert.Equal(t, int32(1), ancestor.Data)
}
