package linkedlist

import (
	"reflect"
	"testing"

	"github.com/zhiruchen/500Algorithms/types"
)

func TestReverseLinkedListPart1IterativeSolution(t *testing.T) {
	list := &types.LinkedList{
		Head: &types.LinkedListNode{
			Data: 10,
			Next: &types.LinkedListNode{
				Data: 9,
				Next: &types.LinkedListNode{
					Data: 8,
					Next: &types.LinkedListNode{
						Data: 7,
						Next: nil,
					},
				},
			},
		},
	}

	newHead := ReverseLinkedListPart1IterativeSolution(list.Head)

	var es []int32
	head := newHead
	for head != nil {
		es = append(es, head.Data)
		head = head.Next
	}

	expect := []int32{7, 8, 9, 10}
	if !reflect.DeepEqual(expect, es) {
		t.Errorf("expect %v. get: %v\n", expect, es)
	}
}
