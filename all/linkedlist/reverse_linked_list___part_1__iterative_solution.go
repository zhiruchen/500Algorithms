package linkedlist

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// ReverseLinkedListPart1IterativeSolution http://www.techiedelight.com/reverse-linked-list-part-1-iterative-solution/
func ReverseLinkedListPart1IterativeSolution(head *types.LinkedListNode) *types.LinkedListNode {
	var cur *types.LinkedListNode

	for head != nil {
		var tmp *types.LinkedListNode
		tmp, head.Next = head.Next, cur
		cur = head
		head = tmp
	}

	return cur
}
