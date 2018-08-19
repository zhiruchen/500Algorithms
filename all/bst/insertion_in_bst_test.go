package bst

import (
	"reflect"
	"testing"

	"github.com/zhiruchen/500Algorithms/types"
)

func TestInsertionInBST(t *testing.T) {
	cases := []struct {
		arr    []int32
		result *types.BinarySearchTree
	}{
		{
			arr: []int32{1, 3, 2},
			result: &types.BinarySearchTree{
				Root: &types.BinTreeNode{
					Data: 1,
					Right: &types.BinTreeNode{
						Data: 3,
						Left: &types.BinTreeNode{Data: 2},
					},
				},
			},
		},
		{
			arr: []int32{5, 6, 4},
			result: &types.BinarySearchTree{
				Root: &types.BinTreeNode{
					Data:  5,
					Left:  &types.BinTreeNode{Data: 4},
					Right: &types.BinTreeNode{Data: 6},
				},
			},
		},
	}

	for _, c := range cases {
		bst := &types.BinarySearchTree{}
		for _, v := range c.arr {
			bst.Root = InsertionInBST(bst.Root, v)
		}

		if !reflect.DeepEqual(bst, c.result) {
			t.Errorf("expect: %v, get %v\n", c.result, bst)
		}
	}
}
