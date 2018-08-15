package types

import (
	"reflect"
	"testing"
)

func TestBST(t *testing.T) {
	arr := []int32{
		1,
		2,
		6,
		3,
		0,
		8,
		9,
	}

	bst := &BinarySearchTree{}
	for _, n := range arr {
		bst.Insert(n)
	}

	data := bst.InOrder()
	expect := []int32{0, 1, 2, 3, 6, 8, 9}
	if !reflect.DeepEqual(data, expect) {
		t.Errorf("expect: %v, get: %v\n", expect, data)
	}
}
