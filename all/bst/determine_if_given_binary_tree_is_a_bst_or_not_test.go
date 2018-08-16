package bst

import (
	"testing"

	. "github.com/zhiruchen/500Algorithms/types"
)

func TestDetermineIfGivenBinaryTreeIsABSTOrNot(t *testing.T) {
	cases := []struct {
		tree  *BinTreeNode
		isBST bool
	}{
		{
			tree: &BinTreeNode{
				Data: 20,
				Left: &BinTreeNode{
					Data: 10,
				},
				Right: &BinTreeNode{
					Data: 30,
					Left: &BinTreeNode{
						Data: 5,
					},
					Right: &BinTreeNode{
						Data: 40,
					},
				},
			},
			isBST: false,
		},
		{
			tree: &BinTreeNode{
				Data: 20,
				Left: &BinTreeNode{
					Data: 10,
				},
				Right: &BinTreeNode{
					Data: 30,
					Left: &BinTreeNode{
						Data: 25,
					},
					Right: &BinTreeNode{
						Data: 40,
					},
				},
			},
			isBST: true,
		},
	}

	for _, c := range cases {
		result := DetermineIfGivenBinaryTreeIsABSTOrNot(c.tree)
		if result != c.isBST {
			t.Errorf("expect: %t, get: %t\n", c.isBST, result)
		}
	}
}
