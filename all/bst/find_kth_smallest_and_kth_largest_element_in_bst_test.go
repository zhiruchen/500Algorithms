package bst

import (
	"testing"

	"github.com/zhiruchen/500Algorithms/types"
)

func TestFindKthSmallestAndKthLargestElementInBST(t *testing.T) {
	cases := []struct {
		arr   []int32
		k     int
		small int32
		large int32
	}{
		{
			arr:   []int32{15, 10, 20, 8, 12, 16, 25},
			k:     2,
			small: 10,
			large: 20,
		},
		{
			arr:   []int32{15, 10, 20, 8, 12, 16, 25},
			k:     5,
			small: 16,
			large: 12,
		},
	}

	for _, c := range cases {
		bt := &types.BinarySearchTree{}
		for _, v := range c.arr {
			bt.Insert(v)
		}

		small, large := FindKthSmallestAndKthLargestElementInBST(bt.Root, c.k)
		if small != c.small {
			t.Errorf("small not equal, expect: %d, get: %d\n", c.small, small)
		}

		if large != c.large {
			t.Errorf("large not equal, expect: %d, get: %d\n", c.large, large)
		}
	}
}
