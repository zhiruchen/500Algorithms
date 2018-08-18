package bst

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zhiruchen/500Algorithms/types"
)

func TestFloorAndCeilInABinarySearchTree(t *testing.T) {
	ast := assert.New(t)

	cases := []struct {
		elements    []int32
		key         int32
		floor, ceil int32
	}{
		{
			elements: []int32{2, 4, 6, 8, 9, 10, 12},
			key:      1,
			floor:    math.MinInt32,
			ceil:     2,
		},
		{
			elements: []int32{2, 4, 6, 8, 9, 10, 12},
			key:      3,
			floor:    2,
			ceil:     4,
		},
	}

	for _, c := range cases {
		bst := &types.BinarySearchTree{}
		for _, v := range c.elements {
			bst.Insert(v)
		}

		floor, ceil := FloorAndCeilInABinarySearchTree(bst.Root, c.key)
		ast.Equal(c.floor, floor, "floor not equal")
		ast.Equal(c.ceil, ceil, "ceil not equal")
	}
}
