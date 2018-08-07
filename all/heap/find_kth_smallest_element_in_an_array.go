package heap

import (
	"github.com/zhiruchen/500Algorithms/types"
)

// FindKthSmallestElementInAnArray http://www.techiedelight.com/find-kth-smallest-element-array/
func FindKthSmallestElementInAnArray(array []int32, k int) int32 {
	h := types.NewMinHeap(array)
	for i := 1; i <= k-1; i++ {
		h.Pop()
	}

	return h.Pop()
}
