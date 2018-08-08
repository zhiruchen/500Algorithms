package linkedlist

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// DetectCycleInALinkedListFloydsCycleDetectionAlgorithm http://www.techiedelight.com/detect-cycle-linked-list-floyds-cycle-detection-algorithm/
// 检测链表是否有环
func DetectCycleInALinkedListFloydsCycleDetectionAlgorithm(head *types.LinkedListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
